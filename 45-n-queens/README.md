# 45 · N-Queens

**Difficulty:** Hard · **Topics:** Backtracking, Bit Manipulation, Recursion · **Asked at:** Google, Amazon, Microsoft, Apple, trading firms

## Problem

Place `n` queens on an `n x n` chessboard so that no two queens attack each other (no shared row, column or diagonal).

Implement two functions:

1. `SolveNQueens(n)` — return **all** distinct solutions. Each solution is a board of `n` strings, each of length `n`, using `'Q'` for a queen and `'.'` for an empty square. Boards may be returned in any order (the tests normalise order).
2. `TotalNQueens(n)` — return only the **number** of distinct solutions.

For `n < 1`, `SolveNQueens` returns an empty result and `TotalNQueens` returns `0`.

## Examples

```text
Input:  n = 4
Output: [
  [".Q..",
   "...Q",
   "Q...",
   "..Q."],

  ["..Q.",
   "Q...",
   "...Q",
   ".Q.."]
]
```

```text
Input:  n = 1
Output: [["Q"]]
```

```text
Input:  TotalNQueens(8)
Output: 92
```

## Constraints

- `0 <= n <= 11` (`n` up to 9 for `SolveNQueens` in practice; `TotalNQueens` is tested up to 11).
- Solutions that are rotations/reflections of each other count as distinct.

## Function Signature

```go
func SolveNQueens(n int) [][]string

func TotalNQueens(n int) int
```

## Hints

<details><summary>Hint 1</summary>

Every row must contain exactly one queen. So place queens **row by row** and, for each row, only decide the column.

</details>

<details><summary>Hint 2</summary>

Two squares `(r1,c1)` and `(r2,c2)` share a "\" diagonal when `r - c` is equal, and a "/" diagonal when `r + c` is equal. Keep three boolean slices: `cols[c]`, `diag1[r+c]`, `diag2[r-c+n]` for O(1) conflict checks.

</details>

<details><summary>Hint 3</summary>

For `TotalNQueens` you can replace the boolean slices with three `int` bitmasks. Available columns in the current row are `^(cols | d1 | d2) & full`; take the lowest set bit with `x & -x`, and when moving to the next row shift the diagonal masks left/right by one.

</details>

## Approach Guide

**Brute force:** try every way to choose `n` squares out of `n²` (or every permutation of columns, `n!`) and check validity. `C(64, 8)` ≈ 4.4 billion for `n = 8` — far too slow; permutations (40,320) are fine for 8 but the check is still wasteful.

**Backtracking (row by row):**

```text
place(row):
  if row == n: record board; return
  for col in 0..n-1:
    if col used or diag1[row+col] or diag2[row-col+n]: continue
    mark; queenCol[row] = col
    place(row + 1)
    unmark
```

Build the board strings only when you reach a full solution (e.g. `strings.Repeat(".", c) + "Q" + ...`).

**Bitmask counting:** the same search, but state is three integers; this is several times faster and the standard optimal answer for `TotalNQueens`. You can also halve the work by symmetry (only place the first queen in the left half and double, handling the middle column for odd `n`).

## Target Complexity

- **Time:** O(n!) upper bound — pruning makes it far smaller in practice.
- **Space:** O(n) for the recursion and conflict sets (plus the output for `SolveNQueens`).

## Go Concepts You'll Practise

- Recursive closures that capture and mutate shared state.
- Multiple assignment for mark/unmark: `a[i], b[j] = true, true`.
- Building strings with `strings.Repeat` / `strings.Builder` / `[]byte`.
- Bitwise ops: `&`, `|`, `&^` (AND NOT), `x & -x`, shifts.

## Interview Follow-ups

1. Return only solutions that are unique up to rotation and reflection.
2. How would you parallelise the search across goroutines (split on the first row's column)?
3. Sudoku Solver uses the same "constraint sets + backtracking" idea — how would you map rows/cols/boxes to bitmasks?
4. For very large `n`, finding **one** solution is easy (constructive formula or min-conflicts heuristic). How does min-conflicts work?

## Run the Tests

```bash
cd 45-n-queens
go test -v
```
