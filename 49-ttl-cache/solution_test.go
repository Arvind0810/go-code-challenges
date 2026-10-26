package ttlcache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

type fakeClock struct {
	mu  sync.Mutex
	now time.Time
}

func newFakeClock() *fakeClock {
	return &fakeClock{now: time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}
}

func (c *fakeClock) Now() time.Time {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.now
}

func (c *fakeClock) Advance(d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.now = c.now.Add(d)
}

func assertGet[K comparable, V comparable](t *testing.T, c *Cache[K, V], k K, wantV V, wantOK bool) {
	t.Helper()
	v, ok := c.Get(k)
	if ok != wantOK || v != wantV {
		t.Fatalf("Get(%v) = (%v, %v), want (%v, %v)", k, v, ok, wantV, wantOK)
	}
}

func assertLen[K comparable, V any](t *testing.T, c *Cache[K, V], want int) {
	t.Helper()
	if got := c.Len(); got != want {
		t.Fatalf("Len() = %d, want %d", got, want)
	}
}

func TestSetGetDelete(t *testing.T) {
	c := New[string, int](time.Minute, newFakeClock())
	assertGet(t, c, "missing", 0, false)
	assertLen(t, c, 0)

	c.Set("a", 1)
	c.Set("b", 2)
	assertGet(t, c, "a", 1, true)
	assertGet(t, c, "b", 2, true)
	assertLen(t, c, 2)

	c.Set("a", 10)
	assertGet(t, c, "a", 10, true)
	assertLen(t, c, 2)

	c.Delete("a")
	assertGet(t, c, "a", 0, false)
	assertLen(t, c, 1)

	c.Delete("does-not-exist")
	assertLen(t, c, 1)
}

func TestDefaultTTLExpiry(t *testing.T) {
	clock := newFakeClock()
	c := New[string, string](10*time.Second, clock)
	c.Set("k", "v")

	clock.Advance(10*time.Second - time.Nanosecond)
	assertGet(t, c, "k", "v", true)

	clock.Advance(time.Nanosecond)
	assertGet(t, c, "k", "", false)
}

func TestLazyExpiryEvictsOnGet(t *testing.T) {
	clock := newFakeClock()
	c := New[int, int](time.Second, clock)
	c.Set(1, 100)
	c.Set(2, 200)
	clock.Advance(2 * time.Second)

	assertLen(t, c, 2)
	assertGet(t, c, 1, 0, false)
	assertLen(t, c, 1)
	assertGet(t, c, 2, 0, false)
	assertLen(t, c, 0)
}

func TestSetResetsExpiry(t *testing.T) {
	clock := newFakeClock()
	c := New[string, int](5*time.Second, clock)
	c.Set("k", 1)
	clock.Advance(4 * time.Second)
	c.Set("k", 2)
	clock.Advance(4 * time.Second)
	assertGet(t, c, "k", 2, true)
	clock.Advance(time.Second)
	assertGet(t, c, "k", 0, false)
}

func TestSetWithTTL(t *testing.T) {
	tests := []struct {
		name    string
		ttl     time.Duration
		advance time.Duration
		wantOK  bool
	}{
		{name: "shorter than default, still fresh", ttl: time.Second, advance: 999 * time.Millisecond, wantOK: true},
		{name: "shorter than default, expired", ttl: time.Second, advance: time.Second, wantOK: false},
		{name: "longer than default, still fresh", ttl: time.Hour, advance: 30 * time.Minute, wantOK: true},
		{name: "zero ttl never expires", ttl: 0, advance: 24 * 365 * time.Hour, wantOK: true},
		{name: "negative ttl never expires", ttl: -time.Second, advance: 24 * time.Hour, wantOK: true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			clock := newFakeClock()
			c := New[string, string](time.Minute, clock)
			c.SetWithTTL("k", "v", tc.ttl)
			clock.Advance(tc.advance)
			want := ""
			if tc.wantOK {
				want = "v"
			}
			assertGet(t, c, "k", want, tc.wantOK)
		})
	}
}

func TestNonPositiveDefaultTTLNeverExpires(t *testing.T) {
	clock := newFakeClock()
	c := New[string, int](0, clock)
	c.Set("k", 7)
	clock.Advance(1000 * time.Hour)
	assertGet(t, c, "k", 7, true)
}

func TestNilClockUsesRealClock(t *testing.T) {
	c := New[string, int](time.Hour, nil)
	c.Set("k", 1)
	assertGet(t, c, "k", 1, true)
}

func TestStructValues(t *testing.T) {
	type user struct {
		Name string
		Age  int
	}
	c := New[int, user](time.Minute, newFakeClock())
	c.Set(1, user{"Ada", 36})
	assertGet(t, c, 1, user{"Ada", 36}, true)
	assertGet(t, c, 2, user{}, false)
}

func waitFor(t *testing.T, cond func() bool, msg func() string) {
	t.Helper()
	deadline := time.Now().Add(2 * time.Second)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatal(msg())
		}
		time.Sleep(5 * time.Millisecond)
	}
}

func TestJanitorEvictsExpiredEntries(t *testing.T) {
	baseline := runtime.NumGoroutine()
	clock := newFakeClock()
	c := New[string, int](time.Second, clock)
	c.Set("a", 1)
	c.Set("b", 2)
	c.SetWithTTL("keep", 3, time.Hour)
	clock.Advance(2 * time.Second)

	c.StartJanitor(5 * time.Millisecond)
	c.StartJanitor(5 * time.Millisecond)
	waitFor(t, func() bool { return c.Len() == 1 }, func() string {
		return fmt.Sprintf("janitor did not evict expired entries within 2s: Len() = %d, want 1", c.Len())
	})
	assertGet(t, c, "keep", 3, true)

	c.Stop()
	c.Stop()
	waitFor(t, func() bool { return runtime.NumGoroutine() <= baseline }, func() string {
		return fmt.Sprintf("janitor goroutine still running after Stop: %d goroutines, want <= %d", runtime.NumGoroutine(), baseline)
	})
}

func TestStopWithoutJanitor(t *testing.T) {
	done := make(chan struct{})
	go func() {
		defer close(done)
		New[string, int](time.Second, newFakeClock()).Stop()
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("Stop() without StartJanitor blocked")
	}
}

func TestConcurrentAccess(t *testing.T) {
	clock := newFakeClock()
	c := New[int, int](50*time.Millisecond, clock)
	c.StartJanitor(time.Millisecond)
	defer c.Stop()

	const (
		goroutines = 16
		ops        = 2000
		keys       = 64
	)
	done := make(chan struct{})
	go func() {
		defer close(done)
		var wg sync.WaitGroup
		for g := range goroutines {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := range ops {
					k := (g*31 + i) % keys
					switch i % 6 {
					case 0, 1:
						c.Set(k, k*10)
					case 2:
						if v, ok := c.Get(k); ok && v != k*10 {
							t.Errorf("Get(%d) = %d, want %d", k, v, k*10)
						}
					case 3:
						c.Delete(k)
					case 4:
						c.SetWithTTL(k, k*10, time.Duration(i%3)*time.Millisecond)
					case 5:
						if n := c.Len(); n > keys {
							t.Errorf("Len() = %d, want at most %d distinct keys", n, keys)
						}
					}
					if i%100 == 0 {
						clock.Advance(10 * time.Millisecond)
					}
				}
			}()
		}
		wg.Wait()
	}()

	select {
	case <-done:
	case <-time.After(10 * time.Second):
		t.Fatal("concurrent operations did not finish within 10s (deadlock?)")
	}
	if n := c.Len(); n < 0 || n > keys {
		t.Fatalf("Len() = %d after concurrent use, want between 0 and %d", n, keys)
	}
}
