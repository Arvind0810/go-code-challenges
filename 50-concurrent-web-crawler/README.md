# 50 · Concurrent Web Crawler (Capstone)

**Difficulty:** Hard · **Topics:** Goroutines, Mutex, Semaphores, BFS, context, System Design · **Asked at:** Google, Meta, Amazon, Dropbox, Cloudflare — the interview edition of the Go Tour's final exercise

## Problem

Crawl a web graph concurrently, starting from `startURL`, using a `Fetcher` (a fake one in tests — no real network).

```go
Crawl(ctx, startURL, maxDepth, maxConcurrency, fetcher) map[string]Result
```

Requirements:

1. **Depth.** `startURL` has depth `0`; a link found on a page at depth `d` has depth `d+1`. Fetch every URL whose **shortest** link distance from the start is `<= maxDepth`, and record that shortest distance in `Result.Depth`. Do not fetch anything deeper. `maxDepth < 0` fetches nothing (empty map).
2. **No duplicates.** Never call `Fetch` for the same URL twice, even when many pages link to it concurrently or the graph has cycles.
3. **Bounded concurrency.** At most `maxConcurrency` `Fetch` calls run at the same time; values below `1` are treated as `1`.
4. **Results.** Return a `Result` for every URL that was fetched, including failed ones (`Err` holds the fetcher's error; `Body`/`URLs` hold what it returned). URLs that were discovered but not fetched are not in the map.
5. **Cancellation.** Once `ctx` is done, start no new fetches. `Fetch` itself takes no context, so wait for in-flight fetches to complete, then return what you have. An already-cancelled `ctx` means zero fetches.
6. **No leaks, no races.** Every goroutine you start has exited when `Crawl` returns; passes `go test -race`.

## Examples

```go
// Go Tour graph:
// golang.org/        -> pkg/, cmd/
// golang.org/pkg/    -> golang.org/, cmd/, pkg/fmt/, pkg/os/
// golang.org/cmd/    -> (not found)
res := Crawl(ctx, "https://golang.org/", 1, 4, fetcher)
// res has 3 keys:
//   "https://golang.org/"     {Depth: 0, Body: "The Go Programming Language", ...}
//   "https://golang.org/pkg/" {Depth: 1, Body: "Packages", ...}
//   "https://golang.org/cmd/" {Depth: 1, Err: not found}
```

```go
// start -> a, x ;  a -> b ;  b -> x, c ;  x -> y
res := Crawl(ctx, "start", 2, 4, fetcher)
// start:0  a:1  x:1  b:2  y:2   — x is depth 1 (via start), not 3 (via a->b),
// so its child y IS crawled; c is depth 3 and is NOT fetched.
```

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
res := Crawl(ctx, "https://example.com/", 50, 16, httpFetcher)
// returns soon after 2s with everything fetched so far
```

## Constraints

- Up to ~10^4 reachable pages in tests; the cancellation test uses an **infinite** graph.
- `Fetch` is safe for concurrent use but may be slow (tests add delays).
- Standard library only.

## Function Signature

```go
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type Result struct {
	Body  string
	URLs  []string
	Depth int
	Err   error
}

func Crawl(ctx context.Context, startURL string, maxDepth int, maxConcurrency int, f Fetcher) map[string]Result
```

## Hints

<details><summary>Hint 1</summary>

Start with a **sequential BFS**: a queue of `(url, depth)` and a `seen` set marked when a URL is **enqueued** (not when fetched). BFS naturally discovers each URL at its shortest depth. Make that pass the tests with `maxConcurrency = 1`, then parallelise.

</details>

<details><summary>Hint 2</summary>

Process the BFS **level by level**: fetch every URL in the current frontier concurrently (a goroutine per URL, limited by a buffered-channel semaphore `make(chan struct{}, maxConcurrency)`), `wg.Wait()`, then build the next frontier from the results, skipping anything in `seen`. Writes to the shared results map need a `sync.Mutex`.

</details>

<details><summary>Hint 3</summary>

Acquire the semaphore with `select { case sem <- struct{}{}: case <-ctx.Done(): }` so a cancelled crawl stops dispatching; check `ctx.Err()` again afterwards (and release the slot if you got one). Release with `defer func() { <-sem }()` inside the goroutine. Still `wg.Wait()` before returning so in-flight fetches don't outlive `Crawl`.

</details>

## Approach Guide

**Go Tour version (buggy for interviews):** recursive `Crawl(url, depth)` that spawns `go Crawl(...)` per link. Typical bugs: fetching URLs twice (no shared `seen` set or check-then-set race), unbounded goroutines, nobody knows when the crawl is finished, and DFS order records a URL at a non-shortest depth — so its children get cut off by the depth limit.

**Level-synchronous BFS (recommended):**

```text
seen := {start}; frontier := [start]
for depth in 0..maxDepth while frontier non-empty and ctx not done:
    for url in frontier:
        acquire semaphore (or stop on ctx.Done)
        wg.Add(1); go fetch url, store Result{..., Depth: depth} under mutex, release
    wg.Wait()
    next := links of this level's results not in seen (mark seen as you add)
    frontier = next
return results
```

Simple, deterministic, and obviously correct. Downside: a single slow page at a level stalls the next level (a "barrier").

**Fully asynchronous (follow-up quality):** a coordinator goroutine owns `seen` and the results; workers send `(url, depth, links)` back over a channel; the coordinator dispatches new URLs immediately and tracks an in-flight counter to know when to stop. Removes the barrier, but to keep shortest depths correct you need a priority queue by depth, or allow re-crawling when a URL is later found at a shallower depth (which the "fetch once" rule forbids).

**Common pitfalls:**

- **Check-then-act race:** `if !seen[u] { seen[u] = true }` must happen under one lock (or in one goroutine).
- `wg.Add` inside the spawned goroutine → `Wait` returns early.
- Holding the mutex while calling `Fetch` → serialises the whole crawl.
- Sending on a results channel nobody reads after cancellation → leaked goroutines.
- Dropping failed fetches — an error page still needs a `Result` with `Err` set.

## Target Complexity

- **Time:** O(V + E) map/set operations for V pages and E links; wall time ≈ (sum of fetch latencies per level) / `maxConcurrency`.
- **Space:** O(V) for `seen` and results, O(width of a level) for the frontier.
- **Guarantees:** each URL fetched at most once; depths are shortest distances; never more than `maxConcurrency` concurrent fetches; prompt stop on cancellation with no leaked goroutines; race-free.

## Go Concepts You'll Practise

- Buffered channels as **counting semaphores**.
- `sync.WaitGroup` + `sync.Mutex` for shared maps.
- `context.Context` checks around blocking operations.
- Interfaces for dependency injection (`Fetcher`) and fakes in tests.
- BFS with a frontier slice; `maps`/`slices` helpers in tests.

## Design Discussion

How would this grow from an interview function into a production crawler?

- **Politeness & robots.txt.** Fetch and cache `/robots.txt` per host and honour `Disallow` rules and `Crawl-delay`. Rate-limit **per host** (a token bucket per domain — see challenge 47) rather than only globally, set a descriptive `User-Agent`, respect `429`/`503` with exponential back-off, and cap concurrent connections per host. Per-host queues (a "frontier of frontiers") keep one slow site from blocking others.
- **URL normalisation.** Lower-case scheme/host, strip fragments, resolve relative links, sort or strip tracking query params, and handle redirects/canonical tags — otherwise "the same page" appears under many URLs and dedupe fails.
- **Distributed crawling.** Partition the URL space by **hash of the host** (consistent hashing) so each worker node owns a set of domains — this keeps per-host politeness local to one node. Nodes forward discovered URLs they don't own to the owner via a queue (Kafka/SQS). The frontier lives in durable storage so crashes don't lose progress; fetches are idempotent so at-least-once delivery is fine.
- **Dedupe at scale.** A `map[string]bool` of billions of URLs doesn't fit in memory. Use a **Bloom filter** (fixed memory, no false negatives, tunable false-positive rate — a false positive just means skipping a page you've never seen, which is usually acceptable), backed by a persistent key-value store for exact checks where needed. Shard the filter alongside the URL partitioning. For **content** dedupe (mirrors, near-duplicate pages), hash the normalised body (SHA-256) or use SimHash/MinHash for near-duplicates.
- **Prioritisation & freshness.** Replace FIFO with a priority queue (PageRank-ish scores, domain importance, change frequency) and schedule re-crawls with conditional requests (`ETag`, `If-Modified-Since`).
- **Robustness.** Timeouts on every fetch (`http.Client{Timeout}` plus a `ctx`), max body size, content-type filtering, trap detection (infinite calendars, session IDs in URLs) via per-host page budgets and depth limits.

## Interview Follow-ups

1. Remove the level barrier: implement the coordinator/worker version and explain how you detect termination.
2. Add per-host rate limiting and a per-host concurrency cap on top of the global `maxConcurrency`.
3. Change `Fetcher` to `Fetch(ctx, url)` so in-flight requests are cancelled too. What changes in `Crawl`?
4. Stream results (`<-chan Result` or `iter.Seq2[string, Result]`) instead of returning a map, for crawls too large for memory.

## Run the Tests

```bash
cd 50-concurrent-web-crawler
go test -v
go test -race -v
```
