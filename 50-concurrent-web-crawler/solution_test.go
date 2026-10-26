package crawler

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

var errNotFound = errors.New("not found")

type fakePage struct {
	body string
	urls []string
}

type fakeFetcher struct {
	pages map[string]fakePage
	gen   func(url string) (fakePage, bool)
	delay time.Duration

	mu     sync.Mutex
	counts map[string]int

	current atomic.Int64
	peak    atomic.Int64
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	n := f.current.Add(1)
	defer f.current.Add(-1)
	for {
		p := f.peak.Load()
		if n <= p || f.peak.CompareAndSwap(p, n) {
			break
		}
	}

	f.mu.Lock()
	if f.counts == nil {
		f.counts = map[string]int{}
	}
	f.counts[url]++
	f.mu.Unlock()

	if f.delay > 0 {
		time.Sleep(f.delay)
	}

	page, ok := f.pages[url]
	if !ok && f.gen != nil {
		page, ok = f.gen(url)
	}
	if !ok {
		return "", nil, fmt.Errorf("fetch %s: %w", url, errNotFound)
	}
	return page.body, slices.Clone(page.urls), nil
}

func (f *fakeFetcher) fetchCounts() map[string]int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return maps.Clone(f.counts)
}

func crawl(t *testing.T, ctx context.Context, start string, depth, conc int, f Fetcher) map[string]Result {
	t.Helper()
	done := make(chan map[string]Result, 1)
	go func() { done <- Crawl(ctx, start, depth, conc, f) }()
	select {
	case res := <-done:
		return res
	case <-time.After(5 * time.Second):
		t.Fatal("Crawl did not return within 5s (deadlock or not respecting cancellation?)")
		return nil
	}
}

func assertFetchedOnce(t *testing.T, f *fakeFetcher) {
	t.Helper()
	for url, n := range f.fetchCounts() {
		if n != 1 {
			t.Errorf("%s fetched %d times, want exactly once", url, n)
		}
	}
}

func checkCrawl(t *testing.T, got map[string]Result, f *fakeFetcher, wantDepths map[string]int) {
	t.Helper()
	assertFetchedOnce(t, f)

	gotURLs := slices.Sorted(maps.Keys(got))
	wantURLs := slices.Sorted(maps.Keys(wantDepths))
	if !slices.Equal(gotURLs, wantURLs) {
		t.Fatalf("crawled URLs = %v, want %v", gotURLs, wantURLs)
	}
	if fetched := slices.Sorted(maps.Keys(f.fetchCounts())); !slices.Equal(fetched, wantURLs) {
		t.Fatalf("fetched URLs = %v, want %v", fetched, wantURLs)
	}

	for url, depth := range wantDepths {
		r := got[url]
		if r.Depth != depth {
			t.Errorf("result[%s].Depth = %d, want %d (shortest link distance from start)", url, r.Depth, depth)
		}
		page, exists := f.pages[url]
		if !exists {
			if !errors.Is(r.Err, errNotFound) {
				t.Errorf("result[%s].Err = %v, want the fetcher's not-found error", url, r.Err)
			}
			continue
		}
		if r.Err != nil {
			t.Errorf("result[%s].Err = %v, want nil", url, r.Err)
		}
		if r.Body != page.body || !slices.Equal(r.URLs, page.urls) {
			t.Errorf("result[%s] = {Body:%q URLs:%v}, want {Body:%q URLs:%v}", url, r.Body, r.URLs, page.body, page.urls)
		}
	}
}

func goTourPages() map[string]fakePage {
	return map[string]fakePage{
		"https://golang.org/": {
			"The Go Programming Language",
			[]string{"https://golang.org/pkg/", "https://golang.org/cmd/"},
		},
		"https://golang.org/pkg/": {
			"Packages",
			[]string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
		},
		"https://golang.org/pkg/fmt/": {
			"Package fmt",
			[]string{"https://golang.org/", "https://golang.org/pkg/"},
		},
		"https://golang.org/pkg/os/": {
			"Package os",
			[]string{"https://golang.org/", "https://golang.org/pkg/"},
		},
	}
}

func TestCrawlGoTour(t *testing.T) {
	const root = "https://golang.org/"
	depth1 := map[string]int{root: 0, "https://golang.org/pkg/": 1, "https://golang.org/cmd/": 1}
	all := maps.Clone(depth1)
	all["https://golang.org/pkg/fmt/"] = 2
	all["https://golang.org/pkg/os/"] = 2

	tests := []struct {
		name     string
		maxDepth int
		conc     int
		want     map[string]int
	}{
		{name: "negative depth fetches nothing", maxDepth: -1, conc: 4, want: map[string]int{}},
		{name: "depth 0 fetches only start", maxDepth: 0, conc: 4, want: map[string]int{root: 0}},
		{name: "depth 1", maxDepth: 1, conc: 4, want: depth1},
		{name: "depth 2", maxDepth: 2, conc: 4, want: all},
		{name: "deep limit", maxDepth: 10, conc: 4, want: all},
		{name: "sequential", maxDepth: 10, conc: 1, want: all},
		{name: "zero concurrency treated as one", maxDepth: 10, conc: 0, want: all},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := &fakeFetcher{pages: goTourPages(), delay: time.Millisecond}
			got := crawl(t, context.Background(), root, tc.maxDepth, tc.conc, f)
			checkCrawl(t, got, f, tc.want)
		})
	}
}

func TestCrawlRecordsShortestDepth(t *testing.T) {
	pages := map[string]fakePage{
		"start": {"s", []string{"a", "x"}},
		"a":     {"a", []string{"b"}},
		"b":     {"b", []string{"x", "c"}},
		"x":     {"x", []string{"y"}},
		"y":     {"y", []string{"z"}},
		"c":     {"c", nil},
	}
	f := &fakeFetcher{pages: pages, delay: 2 * time.Millisecond}
	got := crawl(t, context.Background(), "start", 2, 4, f)
	checkCrawl(t, got, f, map[string]int{"start": 0, "a": 1, "x": 1, "b": 2, "y": 2})
}

func TestCrawlDenseCycleFetchesEachURLOnce(t *testing.T) {
	const n = 30
	pages := map[string]fakePage{}
	want := map[string]int{}
	for i := range n {
		url := fmt.Sprintf("p%d", i)
		var links []string
		for j := range n {
			links = append(links, fmt.Sprintf("p%d", j))
		}
		pages[url] = fakePage{body: url, urls: links}
		want[url] = 1
	}
	want["p0"] = 0
	f := &fakeFetcher{pages: pages, delay: time.Millisecond}
	got := crawl(t, context.Background(), "p0", 5, 8, f)
	checkCrawl(t, got, f, want)
}

func TestCrawlRespectsConcurrencyLimit(t *testing.T) {
	tests := []struct {
		name     string
		children int
		conc     int
		wantPeak int64
	}{
		{name: "limit 4", children: 40, conc: 4, wantPeak: 4},
		{name: "limit 1", children: 8, conc: 1, wantPeak: 1},
		{name: "limit above width", children: 3, conc: 10, wantPeak: 3},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			pages := map[string]fakePage{}
			want := map[string]int{"root": 0}
			var links []string
			for i := range tc.children {
				url := fmt.Sprintf("child%d", i)
				links = append(links, url)
				pages[url] = fakePage{body: url}
				want[url] = 1
			}
			pages["root"] = fakePage{body: "root", urls: links}
			f := &fakeFetcher{pages: pages, delay: 20 * time.Millisecond}
			got := crawl(t, context.Background(), "root", 3, tc.conc, f)
			checkCrawl(t, got, f, want)
			if peak := f.peak.Load(); peak != tc.wantPeak {
				t.Errorf("peak concurrent fetches = %d, want %d", peak, tc.wantPeak)
			}
		})
	}
}

func infiniteTree(url string) (fakePage, bool) {
	k, err := strconv.Atoi(strings.TrimPrefix(url, "node/"))
	if err != nil {
		return fakePage{}, false
	}
	return fakePage{body: url, urls: []string{fmt.Sprintf("node/%d", 2*k), fmt.Sprintf("node/%d", 2*k+1)}}, true
}

func TestCrawlCancellation(t *testing.T) {
	baseline := runtime.NumGoroutine()
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	f := &fakeFetcher{gen: infiniteTree, delay: 10 * time.Millisecond}
	start := time.Now()
	got := crawl(t, ctx, "node/1", 40, 4, f)
	elapsed := time.Since(start)

	if elapsed > time.Second {
		t.Errorf("Crawl returned %v after start, want shortly after the 100ms deadline", elapsed)
	}
	counts := f.fetchCounts()
	if len(counts) == 0 {
		t.Fatal("no URLs fetched before the deadline")
	}
	if len(counts) > 200 {
		t.Errorf("%d URLs fetched; Crawl must stop starting fetches once ctx is done", len(counts))
	}
	assertFetchedOnce(t, f)
	for url := range counts {
		if _, ok := got[url]; !ok {
			t.Errorf("%s was fetched but is missing from the results", url)
		}
	}
	if f.current.Load() != 0 {
		t.Errorf("%d fetches still in flight after Crawl returned; wait for them before returning", f.current.Load())
	}

	deadline := time.Now().Add(2 * time.Second)
	for runtime.NumGoroutine() > baseline {
		if time.Now().After(deadline) {
			t.Fatalf("goroutine leak: %d goroutines running after Crawl returned, want <= %d", runtime.NumGoroutine(), baseline)
		}
		time.Sleep(10 * time.Millisecond)
	}
}

func TestCrawlAlreadyCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	f := &fakeFetcher{gen: infiniteTree}
	got := crawl(t, ctx, "node/1", 10, 4, f)
	if n := len(f.fetchCounts()); n != 0 {
		t.Errorf("%d URLs fetched with an already-cancelled context, want 0", n)
	}
	if len(got) != 0 {
		t.Errorf("Crawl returned %d results with an already-cancelled context, want 0", len(got))
	}
}
