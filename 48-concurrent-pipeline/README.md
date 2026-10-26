# 48 · Concurrent Pipeline (Fan-out / Fan-in)

**Difficulty:** Hard · **Topics:** Channels, Pipelines, Fan-out/Fan-in, context, Goroutine Leaks · **Asked at:** Google, Uber, Cloudflare, Datadog, Go-centric infrastructure teams

## Problem

Build the classic Go pipeline from composable stages. Every stage is a function that owns its goroutine(s) and its **output** channel.

1. `Generate(ctx, nums...)` — emits `nums` **in order** on a new channel, then closes it.
2. `Square(ctx, in)` — reads values from `in` and emits their squares **in order**; closes its output when `in` is closed.
3. `FanOut(ctx, in, n, fn)` — starts `n` copies of the stage `fn`, **all reading from the same `in`**, and returns their `n` output channels (`n < 1` is treated as `1`). Each value from `in` is processed by exactly one copy.
4. `Merge(ctx, chans...)` — fan-in: forwards every value from every input channel onto a single output channel, which is closed once **all** inputs are closed. With zero inputs it returns an already-closed channel.

Every stage must also:

- **Close its output channel** exactly once, when it is done (the sender closes).
- **Stop promptly when `ctx` is done**, even if nobody is reading its output any more — every blocking send and receive must be in a `select` with `<-ctx.Done()`. After cancellation, all output channels are eventually closed.
- **Leak no goroutines.** After the pipeline finishes, or after `ctx` is cancelled (even if the consumer stops reading), `runtime.NumGoroutine()` must return to its previous value.
- Be race-free.

The tests use timeouts, so a stage that never closes its channel fails instead of hanging the test run.

## Examples

```go
ctx := context.Background()
for v := range Square(ctx, Generate(ctx, 1, 2, 3)) {
    fmt.Println(v) // 1, 4, 9 — in order
}
```

```go
ctx := context.Background()
in := Generate(ctx, 1, 2, 3, 4, 5, 6)
workers := FanOut(ctx, in, 3, Square)  // 3 Square stages share `in`
for v := range Merge(ctx, workers...) {
    fmt.Println(v) // 1 4 9 16 25 36 in *some* order
}
```

```go
ctx, cancel := context.WithCancel(context.Background())
out := Merge(ctx, FanOut(ctx, Generate(ctx, hugeSlice...), 4, Square)...)
first := <-out
cancel() // consumer walks away — every goroutine in the pipeline must exit
```

## Constraints

- Up to `10^5` input values; values fit in `int` when squared.
- `fn` passed to `FanOut` follows the same contract as `Square`.
- Unbuffered channels are fine (and make leaks easier to spot); buffered channels are allowed but do not fix leaks.
- Standard library only.

## Function Signature

```go
func Generate(ctx context.Context, nums ...int) <-chan int

func Square(ctx context.Context, in <-chan int) <-chan int

func FanOut(ctx context.Context, in <-chan int, n int, fn func(context.Context, <-chan int) <-chan int) []<-chan int

func Merge(ctx context.Context, chans ...<-chan int) <-chan int
```

The stub also contains an unexported `closedChan()` helper so that unimplemented stages fail fast instead of blocking; you may keep or delete it.

## Hints

<details><summary>Hint 1</summary>

Every stage has the same shape: `out := make(chan int)`, start a goroutine that does `defer close(out)` and the work, then `return out` immediately. Returning `<-chan int` (receive-only) stops callers from closing or sending on it.

</details>

<details><summary>Hint 2</summary>

A plain `out <- v` blocks forever if the consumer has gone away. Write every send as:

```go
select {
case out <- v:
case <-ctx.Done():
    return
}
```

and do the same for receives from `in` (`case v, ok := <-in:`), so a stage also exits if its upstream never closes.

</details>

<details><summary>Hint 3</summary>

`Merge` needs one forwarding goroutine per input plus a `sync.WaitGroup`; a **separate** goroutine does `wg.Wait(); close(out)`. Closing `out` from the forwarders would close it too early (or twice → panic). `FanOut` itself starts no goroutines — it just calls `fn(ctx, in)` `n` times; Go channels already distribute each value to exactly one receiver.

</details>

## Approach Guide

**Naive:** a single goroutine that loops over `nums`, squares, and appends to a slice. Correct, but there is no streaming, no parallelism, and no composition — the interviewer wants to see channels as the "pipes" between independent stages.

**Pipeline without cancellation:** each stage uses `for v := range in { out <- f(v) }` and `defer close(out)`. Works when the consumer drains everything, but if the consumer stops early every upstream goroutine blocks on its send forever — a **goroutine leak**.

**Optimal — cancellable stages:**

```text
stage(ctx, in):
  out := make(chan int)
  go:
    defer close(out)
    loop:
      select:
        v, ok := <-in:   if !ok return
                         select { out <- f(v) | ctx.Done: return }
        ctx.Done:        return
  return out

merge(ctx, chans):
  out := make(chan int); wg
  for each c: wg.Add(1); go forward(c) (same select pattern), wg.Done on exit
  go: wg.Wait(); close(out)
  return out
```

**Common pitfalls:**

- **Closing a channel from the receiver side**, or from several forwarders — "close of closed channel" panic. Only the single owner closes.
- **`wg.Add(1)` inside the goroutine** — `wg.Wait()` may run first and close `out` early.
- **Leaking on early exit** — any send/receive not guarded by `select` with `ctx.Done()`.
- **Loop variable capture** — fixed since Go 1.22, but know why `go func() { use(c) }()` was a bug before.
- **Assuming order after fan-out** — fan-in output order is non-deterministic; the tests sort before comparing.
- **Ranging over a `nil` channel** — blocks forever; return a closed channel when there's nothing to send.

## Target Complexity

- **Time:** O(N) channel operations for N values; with fan-out the CPU-bound stage runs on `n` goroutines in parallel.
- **Space:** O(1) per stage (unbuffered channels) plus O(n + len(chans)) goroutines.
- **Guarantees:** all outputs closed exactly once; order preserved within `Generate` and `Square`; every value delivered exactly once through fan-out/fan-in; prompt shutdown and zero leaked goroutines on cancellation; race-free.

## Go Concepts You'll Practise

- Directional channel types (`<-chan int`, `chan<- int`).
- Ownership rule: the sender closes; `defer close(out)`.
- `select` with `ctx.Done()`, the comma-ok receive `v, ok := <-ch`.
- `sync.WaitGroup` for fan-in, and the "closer goroutine" pattern.
- Detecting leaks with `runtime.NumGoroutine()` (and tools like `go.uber.org/goleak` in real projects).
- Variadic parameters and spreading a slice: `Merge(ctx, outs...)`.

## Interview Follow-ups

1. How would you preserve input order after fan-out (sequence numbers + reorder buffer, or a slice of result channels)?
2. Propagate errors through the pipeline: a stage fails on a bad value and the whole pipeline must stop (`errgroup.WithContext`).
3. Make the stages generic: `Map[T, U any](ctx, in <-chan T, f func(T) U) <-chan U`.
4. When does adding buffered channels help throughput, and when does it just hide back-pressure problems?

## Run the Tests

```bash
cd 48-concurrent-pipeline
go test -v
go test -race -v
```
