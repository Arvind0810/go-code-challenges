# 35 · Rotting Oranges

**Difficulty:** Medium · **Topics:** Graph, Grid, BFS, Multi-source BFS · **Asked at:** Amazon, Microsoft, Google, Flipkart, Walmart

## Problem

You are given an `m x n` grid where each cell is:

- `0` — empty,
- `1` — a fresh orange,
- `2` — a rotten orange.

Every minute, each fresh orange that is **4-directionally adjacent** to a rotten orange becomes rotten.

Return the minimum number of minutes until no fresh orange remains. If that is impossible, return `-1`. If there are no fresh oranges to begin with, return `0`.

## Examples

```text
Minute 0      Minute 1      Minute 2      Minute 3      Minute 4
2 1 1         2 2 1         2 2 2         2 2 2         2 2 2
1 1 0   ──▶   2 1 0   ──▶   2 2 0   ──▶   2 2 0   ──▶   2 2 0
0 1 1         0 1 1         0 1 1         0 2 1         0 2 2

Input:  [[2,1,1],[1,1,0],[0,1,1]]
Output: 4
```

```text
Input:
  2 1 1
  0 1 1
  1 0 1
Output: -1
Explanation: the bottom-left orange is walled off by empty cells and never rots.
```

```text
Input:  [[2,1,1,1,2]]
Output: 2
Explanation: rot spreads from both ends at the same time and meets in the middle.
```

## Constraints

- `1 <= m, n <= 10`
- `grid[i][j]` is `0`, `1`, or `2`.

## Function Signature

```go
func OrangesRotting(grid [][]int) int
```

## Hints

<details><summary>Hint 1</summary>

Running a separate BFS from each rotten orange gives wrong (or slow) answers when there are several. What if all rotten oranges start spreading **at the same time**?

</details>

<details><summary>Hint 2</summary>

Put **every** initially rotten orange into the queue before the BFS starts (multi-source BFS). Also count the fresh oranges up front.

</details>

<details><summary>Hint 3</summary>

Process the queue level by level — one level is one minute. Each time you rot a fresh neighbour, decrement the fresh count. At the end, if `fresh > 0` return `-1`; otherwise return the number of levels that actually rotted something.

</details>

## Approach Guide

**Brute force — simulate the whole grid:** each minute, scan every cell and collect fresh oranges next to a rotten one, then rot them all. Stop when nothing changes. Each pass is O(m·n) and there can be O(m·n) passes → O((m·n)²).

**Optimal — multi-source BFS:**

1. Scan once: enqueue all `2` cells; count `fresh`.
2. `minutes := 0`. While the queue is non-empty **and** `fresh > 0`:
   - Process exactly the cells currently in the queue (one level).
   - For each, check 4 neighbours; if a neighbour is `1`, set it to `2`, `fresh--`, enqueue it.
   - `minutes++` after the level.
3. Return `-1` if `fresh > 0`, else `minutes`.

The `fresh > 0` loop guard prevents counting an extra minute after the last orange rots.

## Target Complexity

- **Time:** O(m·n)
- **Space:** O(m·n) for the queue in the worst case.

## Go Concepts You'll Practise

- `[2]int` arrays as compact, comparable coordinate values (no allocation per point like a struct pointer would need).
- Level-by-level BFS in Go: either swap in a `next` slice per level, or record `size := len(q)` before the inner loop.
- Direction tables as fixed-size arrays: `[4][2]int{{1,0},{-1,0},{0,1},{0,-1}}`.
- Tagless or tagged `switch` for classifying cell states.
- Copying 2D slices (`append([]int(nil), row...)`) when you must not mutate the caller's grid.

## Interview Follow-ups

1. Can you solve it without mutating the input grid?
2. What if each orange had a different rotting delay? (Now it is a weighted shortest-path problem — Dijkstra.)
3. **Walls and Gates / 01 Matrix:** compute the distance from every cell to the nearest source — same multi-source BFS pattern.
4. How would you parallelise the simulation for a very large grid?

## Run the Tests

```bash
cd 35-rotting-oranges
go test -v
```
