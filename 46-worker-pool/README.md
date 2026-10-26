# 46 · Worker Pool

**Difficulty:** Hard · **Topics:** Goroutines, Channels, sync.WaitGroup, context · **Asked at:** Uber, Cloudflare, Google, Stripe, most Go-heavy backend teams

## Problem

Implement a bounded worker pool:

```go
RunPool(ctx, workers, jobs, fn) []Result
```

Requirements:

1. **One result per job, in input order.** `len(results) == len(jobs)` and `results[i]` belongs to `jobs[i]`, even though jobs finish in arbitrary order.
2. **`JobID` and `Err` are set by the pool.** `results[i].JobID == jobs[i].ID`; `results[i].Err` is the error returned by `fn` for that job. One failing job must not stop the others.
3. **Bounded concurrency.** At no point are more than `workers` calls to `fn` running at once. `workers < 1` is treated as `1`. Do not start more goroutines than needed (`min(workers, len(jobs))`).
4. **Cancellation.** `fn` receives `ctx`. Once `ctx` is done, no new job is started; every job that was not started gets `Err = ctx.Err()`. Jobs already running are allowed to finish (they can observe `ctx` themselves).
5. **No leaks.** `RunPool` returns only after every `fn` call it started has returned and every goroutine it started has exited.
6. Must be free of data races (`go test -race`).

## Examples

```go
jobs := []Job{{ID: 1, Input: 2}, {ID: 2, Input: 3}, {ID: 3, Input: 4}}
square := func(_ context.Context, j Job) (Result, error) {
    return Result{Output: j.Input * j.Input}, nil
}
res := RunPool(context.Background(), 2, jobs, square)
// res == []Result{{JobID: 1, Output: 4}, {JobID: 2, Output: 9}, {JobID: 3, Output: 16}}
```

```go
fn := func(_ context.Context, j Job) (Result, error) {
    if j.Input < 0 {
        return Result{}, errors.New("negative input")
    }
    return Result{Output: j.Input}, nil
}
res := RunPool(ctx, 4, []Job{{ID: 7, Input: -1}, {ID: 8, Input: 5}}, fn)
// res[0] == Result{JobID: 7, Err: errors.New("negative input")}
// res[1] == Result{JobID: 8, Output: 5}
```

```go
ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
defer cancel()
res := RunPool(ctx, 2, thousandSlowJobs, slowFn)
// returns shortly after 50ms; most results have errors.Is(r.Err, context.DeadlineExceeded)
```

## Constraints

- `0 <= len(jobs) <= 10^4`
- `workers` may be any `int` (including `0` or negative).
- `fn` is safe to call concurrently.
- Standard library only.

## Function Signature

```go
type Job struct {
	ID    int
	Input int
}

type Result struct {
	JobID  int
	Output int
	Err    error
}

func RunPool(ctx context.Context, workers int, jobs []Job, fn func(context.Context, Job) (Result, error)) []Result
```

## Hints

<details><summary>Hint 1</summary>

Ordering is easy if you pre-allocate `results := make([]Result, len(jobs))` and have each worker write to `results[i]` for the **index** it processed. Distinct indices are distinct memory locations — no mutex needed, and `wg.Wait()` provides the happens-before edge for the final read.

</details>

<details><summary>Hint 2</summary>

Start `min(workers, len(jobs))` goroutines that `range` over an unbuffered `chan int` of job indices. The producer (the `RunPool` goroutine) sends indices and then `close`s the channel — the sender closes, never the receivers.

</details>

<details><summary>Hint 3</summary>

Make the producer's send a `select` with `case idx <- i:` and `case <-ctx.Done():`. On cancellation, fill the remaining results with `ctx.Err()` and stop sending. Workers should also check `ctx.Err()` right after receiving an index, because `select` picks randomly when both cases are ready.

</details>

## Approach Guide

**Naive:** start one goroutine per job and wait on a `WaitGroup`. Ordering is fine with an index-addressed slice, but concurrency is unbounded — 10,000 jobs means 10,000 simultaneous calls (think DB connections or HTTP requests).

**Semaphore variant:** still one goroutine per job, but each acquires a slot from a buffered `chan struct{}` of size `workers` before calling `fn`. Bounded concurrency, but you still create `len(jobs)` goroutines.

**Fixed pool (optimal):**

```text
results := make([]Result, len(jobs))
idx := make(chan int)
for w in 0..min(workers, len(jobs))-1:
    wg.Add(1)                       // before `go`, not inside the goroutine
    go worker: for i := range idx { if ctx done: record ctx.Err(); else call fn, record }
for i in jobs:
    select { send i | ctx.Done: mark remaining as ctx.Err(); stop }
close(idx)
wg.Wait()
return results
```

**Common pitfalls:**

- **`wg.Add` inside the goroutine** — `Wait` can run before `Add` and return early. Always `Add` before the `go` statement.
- **Closing from the wrong side** — only the sender closes `idx`. Closing it from a worker panics the producer with "send on closed channel".
- **Forgetting to close** — workers block forever in `range idx`: a goroutine leak and a hung `wg.Wait()`.
- **Collecting results via a shared `append`** — racy without a mutex, and loses ordering anyway.
- **Ignoring `ctx.Done()` in the producer** — a blocked send keeps the pool alive after cancellation.
- **Overwriting `JobID`/`Err` from `fn`'s return** — the pool must set them.

## Target Complexity

- **Time:** O(len(jobs)) scheduling overhead plus the work itself; wall time ≈ total work / `workers`.
- **Space:** O(len(jobs)) for results, O(workers) goroutines.
- **Guarantees:** input order preserved; peak concurrency ≤ `workers`; no work started after cancellation; zero goroutines left behind when `RunPool` returns; race-free.

## Go Concepts You'll Practise

- Goroutines, unbuffered channels, `range` over a channel, `close`.
- `sync.WaitGroup` and its happens-before guarantee.
- `context.Context`: `Done()`, `Err()`, `context.Canceled` vs `context.DeadlineExceeded`.
- `select` with multiple ready cases and labelled `break`.
- `sync/atomic` (`atomic.Int64`, `CompareAndSwap`) — used by the tests to measure peak concurrency.
- Running tests with the race detector.

## Interview Follow-ups

1. Make it generic: `RunPool[J, R any](ctx, workers, jobs []J, fn func(context.Context, J) (R, error)) []Result[R]`.
2. Add "fail fast": cancel all remaining work on the first error (compare with `golang.org/x/sync/errgroup` and `SetLimit`).
3. What happens if `fn` panics? How would you recover and turn it into a per-job error?
4. Turn it into a long-lived pool with `Submit(job)` and graceful `Shutdown(ctx)` — how do you handle back-pressure when the queue is full?

## Run the Tests

```bash
cd 46-worker-pool
go test -v
go test -race -v
```
