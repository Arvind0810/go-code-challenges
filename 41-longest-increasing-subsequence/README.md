# 41 · Longest Increasing Subsequence

**Difficulty:** Medium · **Topics:** Dynamic Programming, Binary Search, Patience Sorting · **Asked at:** Google, Amazon, Microsoft, Meta, fintech onsite rounds

## Problem

Given an integer slice `nums`, return the length of the longest **strictly increasing subsequence**.

A subsequence is derived from the slice by deleting zero or more elements without changing the order of the remaining elements. For example, `[3, 6, 2, 7]` is a subsequence of `[0, 3, 1, 6, 2, 2, 7]`.

Return `0` for an empty (or `nil`) slice.

## Examples

```text
Input:  nums = [10, 9, 2, 5, 3, 7, 101, 18]
Output: 4
Explanation: one longest increasing subsequence is [2, 3, 7, 101].
```

```text
Input:  nums = [0, 1, 0, 3, 2, 3]
Output: 4
Explanation: [0, 1, 2, 3].
```

```text
Input:  nums = [7, 7, 7, 7, 7]
Output: 1
Explanation: equal values do not count as "increasing".
```

## Constraints

- `0 <= len(nums) <= 2 * 10^4`
- `-10^9 <= nums[i] <= 10^9`
- Strictly increasing: `a < b`, not `a <= b`.

## Function Signature

```go
func LengthOfLIS(nums []int) int
```

## Hints

<details><summary>Hint 1</summary>

Define `dp[i]` = length of the longest increasing subsequence that **ends at index `i`**. How can you compute `dp[i]` from `dp[0..i-1]`?

</details>

<details><summary>Hint 2</summary>

`dp[i] = 1 + max(dp[j])` over all `j < i` with `nums[j] < nums[i]`. The answer is `max(dp)`, not `dp[n-1]`. That is O(n²) — can you do better?

</details>

<details><summary>Hint 3</summary>

Keep a slice `tails` where `tails[k]` is the **smallest possible tail value** of any increasing subsequence of length `k+1` seen so far. `tails` is always sorted, so for each number you can binary search (`sort.SearchInts`) for the first tail `>= x` and either replace it or append.

</details>

## Approach Guide

**Brute force:** enumerate all 2ⁿ subsequences and check each. Exponential — useless beyond ~20 elements, but a good starting point to say out loud.

**O(n²) DP:**

```text
dp := slice of 1s, length n
for i in 0..n-1:
    for j in 0..i-1:
        if nums[j] < nums[i]: dp[i] = max(dp[i], dp[j] + 1)
return max(dp)
```

**O(n log n) — patience sorting:** imagine dealing cards into piles where each card goes on the leftmost pile whose top is `>=` the card (or starts a new pile on the right). The number of piles equals the LIS length.

```text
tails := empty
for x in nums:
    i := lower_bound(tails, x)      // first index with tails[i] >= x
    if i == len(tails): append x
    else: tails[i] = x
return len(tails)
```

Note: `tails` is **not** itself a valid LIS — it only has the right length. Using `lower_bound` (`>=`) gives strictly increasing; `upper_bound` (`>`) would give non-decreasing.

## Target Complexity

- **Time:** O(n log n) (O(n²) is an acceptable first answer)
- **Space:** O(n)

## Go Concepts You'll Practise

- `sort.SearchInts` / `slices.BinarySearch` for lower-bound search.
- Growing a slice with `append` vs overwriting in place.
- Nested loops over indices with `range` and the built-in `max`.
- Handling `nil` vs empty slices identically.

## Interview Follow-ups

1. Return the actual subsequence, not just its length. (Track predecessor indices.)
2. Count the **number** of longest increasing subsequences (LeetCode 673).
3. How would you change the algorithm for longest **non-decreasing** subsequence?
4. Russian Doll Envelopes (LeetCode 354): sort by width asc / height desc, then run LIS on heights — why the descending tie-break?

## Run the Tests

```bash
cd 41-longest-increasing-subsequence
go test -v
```
