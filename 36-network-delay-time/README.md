# 36 · Network Delay Time

**Difficulty:** Medium · **Topics:** Graph, Shortest Path, Dijkstra, Heap (Priority Queue) · **Asked at:** Google, Amazon, Microsoft, Uber, Akamai

## Problem

You are given a network of `n` nodes labelled `1` to `n`, and a list `times` of **directed** edges where `times[i] = [u, v, w]` means a signal takes `w` time to travel from `u` to `v`.

A signal is sent from node `k`. Return the minimum time it takes for **all** `n` nodes to receive it. If some node can never receive the signal, return `-1`.

In other words: compute the shortest-path distance from `k` to every node, and return the largest of those distances.

## Examples

```text
        1
   (2) ───▶ (1)
    │
  1 │
    ▼    1
   (3) ───▶ (4)

Input:  times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
Output: 2
Explanation: dist(1)=1, dist(3)=1, dist(4)=2 → the slowest node hears the signal at time 2.
```

```text
Input:  times = [[1,2,1]], n = 2, k = 2
Output: -1
Explanation: edges are one-way; node 1 is unreachable from node 2.
```

```text
          10
   (1) ────────▶ (2)
    │             ▲
  1 │             │ 2
    ▼             │
   (3) ───────────┘

Input:  times = [[1,2,10],[1,3,1],[3,2,2]], n = 3, k = 1
Output: 3
Explanation: going 1 → 3 → 2 costs 3, which beats the direct edge of 10.
```

## Constraints

- `1 <= k <= n <= 100`
- `0 <= len(times) <= 6000`
- `times[i] = [u, v, w]` with `1 <= u, v <= n`, `u != v`, `0 <= w <= 100`
- Parallel edges between the same pair of nodes may exist.

## Function Signature

```go
func NetworkDelayTime(times [][]int, n int, k int) int
```

## Hints

<details><summary>Hint 1</summary>

All edge weights are non-negative and you need shortest paths from a single source. Which classic algorithm fits exactly that description?

</details>

<details><summary>Hint 2</summary>

Dijkstra: keep a `dist` array initialised to "infinity" (`math.MaxInt`) except `dist[k] = 0`, and a min-heap of `(distance, node)` pairs. Always expand the node with the smallest known distance next.

</details>

<details><summary>Hint 3</summary>

Use **lazy deletion**: push a new `(dist, node)` every time you improve a distance, and when you pop an entry whose distance is larger than `dist[node]`, skip it. The answer is `max(dist[1..n])`, or `-1` if any entry is still infinity.

</details>

## Approach Guide

**Brute force — Bellman-Ford:** relax every edge `n-1` times. Simple and handles negative weights, but O(n·E).

**Alternative — BFS-like relaxation (SPFA):** a queue of nodes whose distance improved; works but worst-case is no better than Bellman-Ford.

**Optimal — Dijkstra with a binary heap:**

1. Build an adjacency list `adj[u] = append(adj[u], edge{v, w})` (size `n+1` because nodes are 1-indexed).
2. `dist[i] = math.MaxInt` for all `i`; `dist[k] = 0`; push `(k, 0)`.
3. While the heap is non-empty:
   - Pop the smallest `(node, d)`. If `d > dist[node]`, continue (stale entry).
   - For each `(next, w)`: if `d + w < dist[next]`, update and push `(next, d+w)`.
4. Scan `dist[1..n]`: if any is still `math.MaxInt`, return `-1`; otherwise return the maximum.

For `n <= 100` an O(n²) array-based Dijkstra (no heap) is also acceptable — mention it as a trade-off for dense graphs.

## Target Complexity

- **Time:** O(E log E) with a binary heap (equivalently O(E log V)).
- **Space:** O(V + E) for the adjacency list, distances and heap.

## Go Concepts You'll Practise

- `container/heap`: implement `Len`, `Less`, `Swap` on a slice type, and `Push(x any)` / `Pop() any` on a **pointer** receiver.
- Call `heap.Push(h, x)` / `heap.Pop(h)` — **not** `h.Push` / `h.Pop` directly, which bypasses the heap ordering.
- Type assertions on `any` (`heap.Pop(h).(item)`).
- `math.MaxInt` as infinity, and why you must check for it before adding a weight (overflow).
- The built-in `max` function (Go 1.21+).

## Interview Follow-ups

1. Why does Dijkstra fail with negative edge weights? What would you use instead?
2. **Cheapest Flights Within K Stops:** limit the path to at most `k` edges. How does the algorithm change?
3. Return the actual path to the slowest node, not just the time. (Track a `prev` array.)
4. If the graph were huge and you only needed the delay to one target, what optimisations apply? (Early exit, bidirectional Dijkstra, A*.)

## Run the Tests

```bash
cd 36-network-delay-time
go test -v
```
