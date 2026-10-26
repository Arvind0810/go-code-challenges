# 33 · Number of Islands

**Difficulty:** Medium · **Topics:** Graph, Grid, DFS, BFS, Union-Find · **Asked at:** Amazon, Google, Microsoft, Bloomberg, Uber

## Problem

Given an `m x n` grid of bytes where `'1'` is land and `'0'` is water, return the number of **islands**.

An island is a maximal group of `'1'` cells connected **horizontally or vertically** (not diagonally). Everything outside the grid is water.

You may modify the grid in place (the tests pass a fresh copy each time).

## Examples

```text
Input:
  1 1 1 1 0
  1 1 0 1 0
  1 1 0 0 0
  0 0 0 0 0
Output: 1
```

```text
Input:
  1 1 0 0 0
  1 1 0 0 0
  0 0 1 0 0
  0 0 0 1 1
Output: 3
Explanation: top-left 2x2 block, the lone centre cell, and the bottom-right pair.
```

```text
Input:
  1 0 1
  0 1 0
  1 0 1
Output: 5
Explanation: diagonal neighbours do not connect.
```

## Constraints

- `0 <= m, n <= 300`
- `grid[i][j]` is `'0'` or `'1'`.

## Function Signature

```go
func NumIslands(grid [][]byte) int
```

## Hints

<details><summary>Hint 1</summary>

Treat each land cell as a graph node with edges to its four neighbours. You are counting **connected components**.

</details>

<details><summary>Hint 2</summary>

Scan the grid. Whenever you find an unvisited `'1'`, you have found a new island — increment the count, then visit every cell of that island so you never count it again.

</details>

<details><summary>Hint 3</summary>

"Visit" can be done by flood fill: overwrite the cell with `'0'` and recurse (DFS) or enqueue (BFS) into its four neighbours, skipping out-of-bounds or water cells. Sinking the island removes the need for a separate `visited` matrix.

</details>

## Approach Guide

**Brute force — visited matrix + DFS:** keep a `[][]bool` of the same size, and for every unvisited land cell run a DFS marking reachable land. Correct and O(m·n), but uses extra memory.

**Optimal — in-place flood fill:**

1. `count := 0`.
2. For each cell `(r, c)`: if it is `'1'`, `count++` and flood fill from it.
3. Flood fill (DFS): if out of bounds or not `'1'`, return; set to `'0'`; recurse into up/down/left/right.
4. Return `count`.

**BFS variant:** same scan, but flood fill with a queue of `[2]int` coordinates. Mark a cell as water **when you enqueue it**, not when you dequeue it, to avoid enqueueing it twice. BFS avoids deep recursion on huge all-land grids.

**Union-Find variant:** union each land cell with its right/down land neighbours; the answer is the number of distinct roots among land cells. Useful when land is added over time (see follow-ups).

## Target Complexity

- **Time:** O(m·n) — each cell is visited a constant number of times.
- **Space:** O(m·n) worst case for the recursion stack / queue (all land); O(1) extra otherwise.

## Go Concepts You'll Practise

- `[][]byte` grids and byte literals (`'1'` is a `byte`/`rune` constant, not a string).
- Recursive closures (`var dfs func(r, c int)`) capturing the grid without passing it around.
- A direction table: `dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}`.
- Slices of `[2]int` as a lightweight queue for BFS.
- Bounds checking order — check `r`/`c` before indexing to avoid a panic.

## Interview Follow-ups

1. Implement it with BFS. When would you prefer BFS over DFS in Go?
2. **Number of Islands II:** cells turn into land one at a time; return the island count after each addition. (Union-Find.)
3. Return the **area of the largest island** instead of the count.
4. What if the grid is too large to fit in memory and arrives row by row?

## Run the Tests

```bash
cd 33-number-of-islands
go test -v
```
