# 47 · Rate Limiter (Token Bucket)

**Difficulty:** Hard · **Topics:** Concurrency, sync.Mutex, System Design, Time · **Asked at:** Stripe, Cloudflare, Uber, Amazon, API-platform teams

## Problem

Implement a **token-bucket** rate limiter.

- The bucket holds at most `capacity` tokens and **starts full**.
- Tokens refill **continuously** at `refillPerSecond` tokens per second (fractional tokens accumulate — after 250 ms at 2 tokens/s you have 0.5 more tokens). The bucket never exceeds `capacity`.
- `Allow()` consumes one token if available and returns `true`; otherwise returns `false`.
- `AllowN(n)` is **all or nothing**: if at least `n` tokens are available, consume `n` and return `true`; otherwise consume nothing and return `false`. `n <= 0` returns `true` and consumes nothing.
- Time comes from an injected `Clock`. A `nil` clock means `RealClock`. If the clock goes **backwards**, no tokens are added (and none are removed).
- All methods must be safe for concurrent use: `N` goroutines racing on a bucket with no refill must get **exactly** `capacity` successes in total.

## Examples

```go
clock := &fakeClock{now: start}           // implements Clock; tests advance it manually
b := NewTokenBucket(2, 1, clock)          // 2 tokens, +1 token per second

b.Allow() // true  (1 left)
b.Allow() // true  (0 left)
b.Allow() // false
clock.Advance(time.Second)
b.Allow() // true  (refilled 1)
```

```go
b := NewTokenBucket(10, 5, clock)
b.AllowN(4) // true  (6 left)
b.AllowN(7) // false (still 6 left — nothing consumed)
b.AllowN(6) // true  (0 left)
clock.Advance(time.Hour)
b.AllowN(11) // false — more than capacity can never succeed
```

```go
limiter := NewTokenBucket(100, 50, nil) // real clock: bursts of 100, 50 req/s sustained
http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
    if !limiter.Allow() {
        http.Error(w, "slow down", http.StatusTooManyRequests)
        return
    }
    // ...
})
```

## Constraints

- `0 <= capacity <= 10^9`
- `0 <= refillPerSecond <= 10^9` (a rate of `0` never refills)
- Tests use a fake clock — never call `time.Now()` directly inside the limiter, always `clock.Now()`.
- Standard library only.

## Function Signature

```go
type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

type TokenBucket struct {
	// TODO: implement
}

func NewTokenBucket(capacity int, refillPerSecond float64, clock Clock) *TokenBucket

func (b *TokenBucket) Allow() bool

func (b *TokenBucket) AllowN(n int) bool
```

## Hints

<details><summary>Hint 1</summary>

You don't need a background goroutine adding tokens on a ticker. Store `tokens float64` and `lastRefill time.Time`; on every call, compute how much time has passed and add `elapsed.Seconds() * rate` tokens ("lazy refill").

</details>

<details><summary>Hint 2</summary>

Refill, check, and consume must happen as **one atomic step** — hold a `sync.Mutex` for the whole operation. Two separate locked sections ("check" then "consume") let two goroutines both see one token and both take it.

</details>

<details><summary>Hint 3</summary>

Clamp with `min(capacity, tokens + added)`. Only update `lastRefill` (and add tokens) when `now.Sub(lastRefill) > 0`, which handles a clock that jumps backwards. `Allow()` is just `AllowN(1)`.

</details>

## Approach Guide

**Naive — fixed window counter:** count requests per wall-clock second and reset when the second changes. Simple, but allows `2 × limit` requests across a window boundary (end of one second + start of the next).

**Naive — background refiller:** a goroutine with a `time.Ticker` adds tokens every tick. Works, but costs a goroutine per bucket (bad for per-user limiters), needs a `Stop()`, and is hard to test without sleeping.

**Optimal — lazy token bucket:**

```text
AllowN(n):
  if n <= 0: return true
  lock
  now := clock.Now()
  if now after last:
      tokens = min(capacity, tokens + (now - last).Seconds() * rate)
      last = now
  if tokens < n: unlock; return false
  tokens -= n
  unlock; return true
```

O(1) per call, no goroutines, and deterministic under a fake clock.

**Pitfalls:**

- Reading `time.Now()` directly (untestable, flaky tests that sleep).
- Integer token counts with integer division — at low rates you lose fractional refills forever.
- Updating `last` even when the clock went backwards → future refills are miscalculated.
- Locking with `RLock` for the "read" part — every call mutates state, so it needs the write lock.
- Copying a struct that contains a `sync.Mutex` (always use `*TokenBucket`; `go vet` catches this).

## Target Complexity

- **Time:** O(1) per `Allow`/`AllowN`.
- **Space:** O(1) per bucket.
- **Guarantees:** never more than `capacity + rate × elapsed` successes over any period; exact accounting under concurrency (no lost or double-spent tokens); race-free under `go test -race`.

## Go Concepts You'll Practise

- `sync.Mutex` with `defer Unlock()`.
- Interfaces for dependency injection (`Clock`) and test fakes.
- `time.Duration`, `time.Time.Sub`, `Duration.Seconds()`.
- `sync/atomic` counters and `sync.WaitGroup` in concurrent tests.
- Zero-value-friendly constructors (`nil` clock → default).

## Interview Follow-ups

1. **Sliding window log:** store request timestamps in a ring buffer / deque and drop those older than the window. What are the memory costs vs. a token bucket? What about the sliding-window **counter** approximation?
2. **Per-user limiter:** `map[userID]*TokenBucket` behind a mutex (or `sync.Map`/sharded maps). How do you evict idle users so memory doesn't grow forever?
3. **Distributed limiter (Redis):** multiple API servers share limits. Store `tokens` and `last` in a Redis hash and do refill+consume in a single Lua script (atomic). How do you handle clock skew between servers and Redis latency? When is approximate local limiting with periodic sync good enough?
4. Add `Wait(ctx, n)` that blocks until tokens are available or `ctx` is done — how do you compute how long to sleep (compare `golang.org/x/time/rate`)?

## Run the Tests

```bash
cd 47-rate-limiter
go test -v
go test -race -v
```
