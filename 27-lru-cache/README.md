# 27 · LRU Cache

**Difficulty:** Medium · **Topics:** Design, Hash Map, Doubly Linked List · **Asked at:** Amazon, Google, Meta, Microsoft, Uber, Netflix

## Problem

Design a data structure that behaves like a **Least Recently Used (LRU) cache** with a fixed capacity.

- `NewLRUCache(capacity)` — creates an empty cache holding at most `capacity` entries.
- `Get(key)` — returns the value for `key` if present (and marks it as most recently used), otherwise `-1`.
- `Put(key, value)` — inserts or updates `key` and marks it as most recently used. If this makes the cache exceed `capacity`, evict the **least recently used** key.

Both `Get` and `Put` must run in **O(1)** average time.

## Examples

```
cache := NewLRUCache(2)
cache.Put(1, 1)     // cache: {1=1}
cache.Put(2, 2)     // cache: {1=1, 2=2}
cache.Get(1)        // returns 1    order (MRU -> LRU): 1, 2
cache.Put(3, 3)     // evicts key 2 order: 3, 1
cache.Get(2)        // returns -1
cache.Put(4, 4)     // evicts key 1 order: 4, 3
cache.Get(1)        // returns -1
cache.Get(3)        // returns 3
cache.Get(4)        // returns 4
```

```
cache := NewLRUCache(2)
cache.Put(1, 1)
cache.Put(2, 2)
cache.Put(1, 100)   // update: no eviction, key 1 becomes MRU
cache.Get(1)        // returns 100
cache.Get(2)        // returns 2
```

```
The structure you're aiming for:

  map[key] ──────────────┐
                         v
  head <-> [4:4] <-> [3:3] <-> tail
           MRU               LRU (evict from here)
```

## Constraints

- `1 <= capacity <= 3000`
- `0 <= key <= 10^4`
- `0 <= value <= 10^5`
- At most `2 * 10^5` calls to `Get` and `Put`.

## Function Signature

```go
type LRUCache struct{}

func NewLRUCache(capacity int) *LRUCache
func (c *LRUCache) Get(key int) int
func (c *LRUCache) Put(key, value int)
```

You will need to add fields to `LRUCache` (and probably a node type).

## Hints

<details><summary>Hint 1</summary>

A map gives O(1) lookup but no ordering. A list gives ordering but O(n) lookup. Combine them: the map stores **pointers to list nodes**.

</details>

<details><summary>Hint 2</summary>

Use a **doubly** linked list so that, given a node pointer, you can unlink it in O(1) without walking the list. Most recent at the front, least recent at the back.

</details>

<details><summary>Hint 3</summary>

Sentinel `head` and `tail` nodes remove almost every nil check. Each node must store its **key** too — when you evict the tail node you need the key to `delete` it from the map.

</details>

## Approach Guide

**Brute force:** Store entries in a slice ordered by recency; `Get`/`Put` search the slice and move the entry to the front. O(n) per operation.

**Optimal (hash map + doubly linked list):**
1. `node{key, val, prev, next}`; cache holds `capacity`, `map[int]*node`, and sentinels.
2. Helpers: `remove(n)` unlinks a node; `addToFront(n)` inserts after `head`.
3. `Get`: look up; if missing return -1; else `remove` + `addToFront`, return value.
4. `Put`: if present, update value, `remove` + `addToFront`. Otherwise create a node, store in map, `addToFront`; if `len(map) > capacity`, remove `tail.prev` and delete its key from the map.

**About `container/list`:** the stdlib has a ready-made doubly linked list (`list.PushFront`, `list.MoveToFront`, `list.Remove`, `e.Value.(entry)`). It's a perfectly valid production choice — but in interviews you'll usually be asked to hand-roll the list, so do that first, then try a second version with `container/list` and compare.

## Target Complexity

- Time: **O(1)** average for `Get` and `Put`.
- Space: **O(capacity)**

## Go Concepts You'll Practise

- Struct types with self-referential pointer fields (`prev, next *node`).
- Maps of pointers (`map[int]*node`) and the comma-ok idiom (`n, ok := m[key]`).
- `delete(m, key)` and `len(m)`.
- Initialising internal state in a constructor (a nil map panics on write!).
- Unexported helper methods to keep the public API small.
- `container/list` and type assertions on `any` (for the alternative version).

## Interview Follow-ups

1. Make the cache safe for concurrent use. `sync.Mutex` vs. `sync.RWMutex` — why doesn't `RWMutex` help much for `Get` here?
2. Add a TTL so entries expire after a duration.
3. Implement an **LFU** cache (Least Frequently Used) in O(1).
4. Make it generic: `LRUCache[K comparable, V any]`.

## Run the Tests

```bash
cd 27-lru-cache
go test -v
```
