# 49 · TTL Cache

**Difficulty:** Hard · **Topics:** Generics, sync.RWMutex, Maps, Background Goroutines, Time · **Asked at:** Amazon, Uber, Grab, Razorpay, any backend team with a caching layer

## Problem

Implement a generic, goroutine-safe in-memory cache whose entries expire after a **time-to-live (TTL)**.

Requirements:

1. `New[K, V](defaultTTL, clock)` creates an empty cache. A `nil` clock means `RealClock`. All time is read through `clock.Now()`.
2. `Set(k, v)` stores a value with the cache's **default TTL**; `SetWithTTL(k, v, ttl)` uses a per-entry TTL. A TTL `<= 0` (default or per-entry) means **never expires**. Setting an existing key replaces the value **and resets its expiry**.
3. `Get(k)` returns `(value, true)` only if the key exists and is not expired. An entry written at time `t` with TTL `d` is **expired once `now >= t + d`**.
4. **Lazy expiry:** `Get` removes an expired entry it finds (so memory is reclaimed on access).
5. `Delete(k)` removes a key; deleting a missing key is a no-op.
6. `Len()` returns the number of **stored** entries — this includes expired entries that have not been evicted yet (it is a cheap O(1) count, not a scan).
7. **Janitor (follow-up, tested):** `StartJanitor(interval)` starts one background goroutine that evicts all expired entries every `interval`. A second call while running has no effect. `Stop()` stops it and **waits for the goroutine to exit**; it is safe to call twice or without a janitor.
8. All methods are safe for concurrent use and race-free; use `sync.RWMutex` so concurrent `Get`s don't serialise.

## Examples

```go
clock := &fakeClock{now: start}
c := New[string, int](10*time.Second, clock)

c.Set("a", 1)
c.Get("a")                  // 1, true
clock.Advance(9 * time.Second)
c.Get("a")                  // 1, true
clock.Advance(1 * time.Second)
c.Get("a")                  // 0, false  (expired exactly at t+10s, and evicted)
```

```go
c := New[string, string](time.Minute, clock)
c.SetWithTTL("session", "abc", 5*time.Second)
c.SetWithTTL("config", "v1", 0)       // never expires
clock.Advance(time.Hour)
c.Len()                                // 2 — nothing evicted yet (lazy)
c.Get("session")                       // "", false — evicted now
c.Len()                                // 1
```

```go
c := New[int, User](5*time.Minute, nil) // real clock
c.StartJanitor(30 * time.Second)
defer c.Stop()
```

## Constraints

- Up to `10^5` entries; any `comparable` key type and any value type.
- Tests use a fake clock for expiry. The janitor's **tick interval** uses real time (`time.Ticker`), but expiry decisions must still use `clock.Now()`.
- Standard library only.

## Function Signature

```go
type Clock interface {
	Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time { return time.Now() }

type Cache[K comparable, V any] struct {
	// TODO: implement
}

func New[K comparable, V any](defaultTTL time.Duration, clock Clock) *Cache[K, V]

func (c *Cache[K, V]) Set(k K, v V)

func (c *Cache[K, V]) SetWithTTL(k K, v V, ttl time.Duration)

func (c *Cache[K, V]) Get(k K) (V, bool)

func (c *Cache[K, V]) Delete(k K)

func (c *Cache[K, V]) Len() int

func (c *Cache[K, V]) StartJanitor(interval time.Duration)

func (c *Cache[K, V]) Stop()
```

## Hints

<details><summary>Hint 1</summary>

Store a small generic struct per key: `type entry[V any] struct { value V; expiresAt time.Time }`, in a `map[K]entry[V]`. Use the zero `time.Time` to mean "never expires" (`expiresAt.IsZero()`). Expired means `!now.Before(expiresAt)`.

</details>

<details><summary>Hint 2</summary>

`Get` can read under `RLock`, but deleting an expired entry needs the write lock. You cannot upgrade an `RLock` in Go — release it, take `Lock`, and **re-check** the entry (another goroutine may have just `Set` a fresh value under the same key; don't delete that one).

</details>

<details><summary>Hint 3</summary>

For the janitor, keep `stop chan struct{}` and `done chan struct{}`. The goroutine `defer close(done)` and loops on `select { case <-ticker.C: evict; case <-stop: return }`. `Stop()` closes `stop` then waits on `<-done`. Guard start/stop with a separate mutex (or `sync.Once` for stop) so a double `Stop()` doesn't panic with "close of closed channel". Remember `defer ticker.Stop()`.

</details>

## Approach Guide

**Naive:** a `map` plus one `time.AfterFunc` timer per entry that deletes it. Simple, but one timer per key is heavy at scale, every `Set` must cancel the old timer, and it can't be tested without real sleeping.

**Lazy expiry (core):**

```text
SetWithTTL(k, v, ttl):
  exp := zero time if ttl <= 0 else clock.Now() + ttl
  lock; items[k] = {v, exp}; unlock

Get(k):
  now := clock.Now()
  rlock; e, ok := items[k]; runlock
  if !ok: return zero, false
  if not expired(e, now): return e.value, true
  lock; if items[k] still expired: delete; unlock
  return zero, false
```

**Plus active expiry (janitor):** lazy expiry alone never frees keys that are written once and never read again — a memory leak for write-heavy workloads. A periodic sweep fixes that. Redis combines both: lazy expiry on access + a probabilistic sampler of keys with TTLs.

**Pitfalls:**

- Returning a pointer to internal state or holding the lock while calling user code.
- Reading the map without a lock "because it's just a read" — concurrent map read/write is a fatal runtime error, not just a race.
- Using `time.Now()` directly instead of the injected clock.
- A janitor goroutine that is never stopped (leak), or a `Stop` that returns before the goroutine has exited.
- Long sweeps holding the write lock — for huge caches, sweep in batches or shard the map.

## Target Complexity

- **Time:** `Set`, `Get`, `Delete`, `Len` — O(1) average. Janitor sweep — O(n) per tick.
- **Space:** O(n) entries.
- **Guarantees:** never returns an expired value; expiry boundary is exact (`now >= expiresAt`); janitor runs at most once and fully stops on `Stop()`; race-free under `go test -race`.

## Go Concepts You'll Practise

- Generics: type parameters on types and functions, the `comparable` constraint, `var zero V`.
- `sync.RWMutex` (`RLock` vs `Lock`) and why lock upgrades need a re-check.
- Background goroutines with `time.Ticker`, stop/done channels and clean shutdown.
- Interfaces for testable time (`Clock`) and `time.Time` arithmetic (`Add`, `Before`, `IsZero`).

## Interview Follow-ups

1. Add a max size with **LRU eviction** (combine with challenge 27): which data structure gives O(1) for both TTL checks and recency?
2. Replace the O(n) janitor sweep with a **min-heap** or timing wheel keyed by expiry time. What changes when a key's TTL is reset?
3. How would you reduce lock contention under heavy load (sharding by key hash, `sync.Map` trade-offs)?
4. Cache stampede: 1,000 requests miss the same expired key simultaneously. How do you ensure only one recomputes it (`singleflight`, early refresh with jitter)?

## Run the Tests

```bash
cd 49-ttl-cache
go test -v
go test -race -v
```
