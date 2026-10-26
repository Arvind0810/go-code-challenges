# 39 · House Robber

**Difficulty:** Medium · **Topics:** Dynamic Programming, 1D DP, Array · **Asked at:** Amazon, Google, Microsoft, Cisco, Adobe

## Problem

You are a robber planning to rob houses along a street. `nums[i]` is the amount of money in house `i`. Adjacent houses have connected alarms: if you rob two **adjacent** houses on the same night, the police are alerted.

Return the maximum amount of money you can rob without robbing two adjacent houses. An empty street yields `0`.

## Examples

```text
Houses:   [1]  [2]  [3]  [1]
Rob:       ✓         ✓
Input:  nums = [1, 2, 3, 1]
Output: 4
Explanation: rob house 0 (1) and house 2 (3) → 1 + 3 = 4.
```

```text
Houses:   [2]  [7]  [9]  [3]  [1]
Rob:       ✓         ✓         ✓
Input:  nums = [2, 7, 9, 3, 1]
Output: 12
Explanation: 2 + 9 + 1 = 12.
```

```text
Input:  nums = [2, 1, 1, 2]
Output: 4
Explanation: rob house 0 and house 3 — skipping two houses in a row is sometimes best,
so "rob every other house" is not a valid strategy.
```

## Constraints

- `0 <= len(nums) <= 100`
- `0 <= nums[i] <= 400`

## Function Signature

```go
func Rob(nums []int) int
```

## Hints

<details><summary>Hint 1</summary>

Standing at house `i`, you have exactly two choices: rob it or skip it. What does each choice imply about house `i-1`?

</details>

<details><summary>Hint 2</summary>

Let `best(i)` be the most you can rob from houses `0..i`. If you rob house `i`, you get `nums[i] + best(i-2)`; if you skip it, you get `best(i-1)`. Take the max.

</details>

<details><summary>Hint 3</summary>

`best(i)` only depends on the previous two values, so you don't need an array — two variables (`prev`, `cur`) rolled forward through the loop are enough. Go's tuple assignment `prev, cur = cur, max(cur, prev+x)` does it in one line.

</details>

## Approach Guide

**State:** `dp[i]` = maximum money obtainable from the first `i` houses (houses `0..i-1`).

**Transition:** `dp[i] = max(dp[i-1], dp[i-2] + nums[i-1])` — skip house `i-1`, or rob it and add the best up to two houses back.

**Base cases:** `dp[0] = 0` (no houses), `dp[1] = nums[0]`.

**Answer:** `dp[len(nums)]`.

**1. Plain recursion (brute force):** `rob(i) = max(rob(i+1), nums[i] + rob(i+2))`, with `rob(i) = 0` for `i >= n`. Correct but O(2ⁿ) because the same `i` is recomputed many times.

**2. Memoisation (top-down):** cache `rob(i)` in a `[]int` initialised to `-1` (or a `map[int]int`). Each index is computed once → O(n) time, O(n) space (cache + recursion stack).

**3. Tabulation (bottom-up):** allocate `dp := make([]int, n+1)`, fill from `i = 2..n` using the transition. O(n) time, O(n) space, no recursion.

**4. Space-optimised:** notice `dp[i]` uses only `dp[i-1]` and `dp[i-2]`. Keep two integers `prev` (= `dp[i-2]`) and `cur` (= `dp[i-1]`) and roll them forward. O(n) time, O(1) space.

## Target Complexity

- **Time:** O(n)
- **Space:** O(1) (O(n) for the memo/tabulation versions)

## Go Concepts You'll Practise

- Tuple (parallel) assignment `a, b = b, a+x` — the right-hand side is evaluated fully before assigning, so no temp variable is needed.
- The built-in `max` (Go 1.21+) instead of a hand-written helper.
- Recursive closures with a memo slice captured from the enclosing scope.
- `for _, x := range nums` safely handling `nil`/empty slices without special cases.

## Interview Follow-ups

1. **House Robber II:** the houses are arranged in a **circle**, so the first and last houses are adjacent. (Hint: run the linear solution twice — on `nums[1:]` and `nums[:n-1]` — and take the max; handle `n == 1` separately.)
2. Return **which** houses to rob, not just the total. (Keep the `dp` array and backtrack.)
3. **House Robber III:** the houses form a binary tree; you cannot rob a parent and its child. (Tree DP returning a `(rob, skip)` pair per node.)
4. **Delete and Earn:** reduce it to House Robber by bucketing values.

## Run the Tests

```bash
cd 39-house-robber
go test -v
```
