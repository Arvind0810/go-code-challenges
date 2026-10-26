# 19 · Rotate Matrix

**Difficulty:** Medium · **Topics:** Arrays, Matrix, In-Place Manipulation · **Asked at:** Amazon, Microsoft, Apple, Cisco

## Problem

You are given an `n × n` matrix of integers. Rotate it **90 degrees clockwise, in place**.

You must modify the input matrix directly — do not allocate another 2D matrix for the result.

## Examples

```
Input:
  [[1, 2, 3],
   [4, 5, 6],
   [7, 8, 9]]
Output:
  [[7, 4, 1],
   [8, 5, 2],
   [9, 6, 3]]
Explanation: The first column read bottom-to-top becomes the first row.
```

```
Input:
  [[ 5,  1,  9, 11],
   [ 2,  4,  8, 10],
   [13,  3,  6,  7],
   [15, 14, 12, 16]]
Output:
  [[15, 13,  2,  5],
   [14,  3,  4,  1],
   [12,  6,  8,  9],
   [16,  7, 10, 11]]
```

```
Input:  [[1]]
Output: [[1]]
```

## Constraints

- `0 <= n <= 100`
- `matrix` is square: every row has length `n`.
- `-1000 <= matrix[i][j] <= 1000`

## Function Signature

```go
func Rotate(matrix [][]int)
```

## Hints

<details><summary>Hint 1</summary>

Track where a single cell moves: element `(i, j)` ends up at `(j, n-1-i)`.

</details>

<details><summary>Hint 2</summary>

A clockwise rotation equals two simpler in-place operations: a transpose followed by something else.

</details>

<details><summary>Hint 3</summary>

Transpose (swap `m[i][j]` with `m[j][i]` for `j > i`), then reverse each row. Alternatively, rotate four cells at a time, layer by layer, from the outside in.

</details>

## Approach Guide

**Brute force — O(n²) time, O(n²) space:** build a new matrix with `out[j][n-1-i] = m[i][j]`, then copy it back. Simple, but violates the in-place requirement.

**Optimal A — transpose + reverse, O(1) extra space:**
1. Transpose: for each `i`, for `j` from `i+1` to `n-1`, swap `m[i][j]` and `m[j][i]`.
2. Reverse every row in place (two pointers, or `slices.Reverse`).

**Optimal B — four-way cycle swap:**
1. For each layer `l` from `0` to `n/2 - 1`, for each offset in the layer,
2. save the top cell, then move left → top, bottom → left, right → bottom, saved → right.

Both touch every cell a constant number of times.

## Target Complexity

- **Time:** O(n²)
- **Space:** O(1) extra

## Go Concepts You'll Practise

- Tuple assignment for swaps: `m[i][j], m[j][i] = m[j][i], m[i][j]`
- `slices.Reverse` on each row
- Slices are reference-like: mutating `matrix[i][j]` is visible to the caller without returning anything
- Iterating 2D slices with nested `for` loops and careful index bounds

## Interview Follow-ups

- Rotate **counter-clockwise**. Which two operations would you combine?
- Rotate by 180°.
- How would you handle a non-square `m × n` matrix? (In place is no longer possible — why?)
- Rotate an image represented as a flat 1D slice of length `n*n`.

## Run the Tests

```bash
cd 19-rotate-matrix
go test -v
```
