# 07 · Maximum Subarray

**Difficulty:** Easy · **Topics:** Arrays, Dynamic Programming, Greedy · **Asked at:** Amazon, Microsoft, LinkedIn, Apple, Goldman Sachs

## Problem

Given a non-empty slice of integers `nums`, find the contiguous subarray (containing at least one element) with the largest sum, and return that sum.

## Examples

```text
Input:  nums = [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output: 6
Explanation: the subarray [4, -1, 2, 1] has the largest sum, 6.
```

```text
Input:  nums = [1]
Output: 1
```

```text
Input:  nums = [-3, -1, -2]
Output: -1
Explanation: every subarray sum is negative; the best is the single element [-1].
```

## Constraints

- `1 <= len(nums) <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

## Function Signature

```go
func MaxSubArray(nums []int) int
```

## Hints

<details><summary>Hint 1</summary>

Think about the best subarray that **ends at** index `i`. How does it relate to the best subarray ending at `i-1`?

</details>

<details><summary>Hint 2</summary>

At each index you have exactly two choices: extend the previous subarray, or start a new one at `nums[i]`. Keep whichever is larger, and track the best seen overall.

</details>

<details><summary>Hint 3</summary>

Initialise `cur` and `best` to `nums[0]` (not 0 — think all-negative input). For each later `x`: `cur = max(x, cur + x)`, then `best = max(best, cur)`.

</details>

## Approach Guide

**Brute force:** for every start `i`, extend to every end `j` with a running sum and track the maximum. O(n²) time, O(1) space. (A triple loop recomputing each sum is O(n³) — avoid it.)

**Divide and conquer:** the best subarray lies entirely in the left half, entirely in the right half, or crosses the middle. O(n log n). Good to mention, rarely what the interviewer wants coded.

**Optimal (Kadane's algorithm):**

1. Let `cur` be the best sum of a subarray ending at the current index, and `best` the best sum seen so far. Start both at `nums[0]`.
2. For each subsequent element `x`, decide: is `cur + x` better than starting fresh with `x`? Update `cur` accordingly.
3. Update `best` if `cur` beats it.
4. Return `best`.

This is dynamic programming with the DP table collapsed to a single variable.

## Target Complexity

- **Time:** O(n)
- **Space:** O(1)

## Go Concepts You'll Practise

- Built-in generic `max` / `min` (Go 1.21+) instead of hand-written helpers.
- Re-slicing: `for _, x := range nums[1:]` without copying.
- Choosing safe initial values (`nums[0]` vs `0` vs `math.MinInt`).
- Integer overflow awareness — Go `int` is 64-bit on most platforms, but wraps silently.

## Interview Follow-ups

1. Return the start and end indices of the best subarray, not just the sum.
2. **Maximum Product Subarray:** what extra state do you need when multiplying (negatives flip signs)?
3. **Maximum Sum Circular Subarray:** the array wraps around. How does Kadane adapt?
4. What should the function do if `nums` is empty, and how would you express that in Go's API (error, `ok` bool, or panic)?

## Run the Tests

```bash
cd 07-maximum-subarray
go test -v
```
