# 14 · Three Sum

**Difficulty:** Medium · **Topics:** Arrays, Sorting, Two Pointers · **Asked at:** Meta, Amazon, Google, Apple

## Problem

Given an integer slice `nums`, return all **unique** triplets `[nums[i], nums[j], nums[k]]` such that `i`, `j`, `k` are distinct indices and `nums[i] + nums[j] + nums[k] == 0`.

The result must not contain duplicate triplets. The order of the triplets, and the order of numbers inside each triplet, does not matter. If there are none, return an empty (or nil) slice.

## Examples

```
Input:  nums = [-1, 0, 1, 2, -1, -4]
Output: [[-1, -1, 2], [-1, 0, 1]]
Explanation: (-1) + (-1) + 2 = 0 and (-1) + 0 + 1 = 0.
The triplet [-1, 0, 1] can be formed in two ways but appears only once.
```

```
Input:  nums = [0, 1, 1]
Output: []
Explanation: No three numbers sum to zero.
```

```
Input:  nums = [0, 0, 0, 0]
Output: [[0, 0, 0]]
```

## Constraints

- `0 <= len(nums) <= 3000`
- `-10^5 <= nums[i] <= 10^5`
- You may sort or otherwise reorder `nums`.

## Function Signature

```go
func ThreeSum(nums []int) [][]int
```

## Hints

<details><summary>Hint 1</summary>

Fix one number. The problem becomes "find two numbers that sum to its negation" — a Two Sum.

</details>

<details><summary>Hint 2</summary>

Sort first. On a sorted slice, Two Sum can be solved with two pointers, and duplicates sit next to each other so they are easy to skip.

</details>

<details><summary>Hint 3</summary>

For each `i` (skipping when `nums[i] == nums[i-1]`), set `lo = i+1`, `hi = n-1`. If the sum is too small, `lo++`; too big, `hi--`; equal — record it, then advance `lo` past all equal values (and `hi` likewise).

</details>

## Approach Guide

**Brute force — O(n³):** three nested loops, sort each found triplet and put it in a set to dedupe. Too slow and memory-heavy.

**Hash-set Two Sum — O(n²):** for each `i`, run a hash-based Two Sum on the rest. Works, but deduplication is fiddly.

**Optimal — sort + two pointers, O(n²):**
1. Sort `nums`.
2. For each index `i`: if `nums[i] > 0`, stop early (all remaining are positive). Skip `i` if it equals the previous value.
3. Run two pointers `lo`, `hi` on the subarray to the right of `i`, targeting `-nums[i]`.
4. On a match, append the triplet, then move both pointers while skipping duplicates.

## Target Complexity

- **Time:** O(n²)
- **Space:** O(1) extra beyond the output (ignoring sort's internal stack)

## Go Concepts You'll Practise

- `slices.Sort` from the standard library (Go 1.21+)
- Building a `[][]int` with `append(result, []int{a, b, c})`
- Duplicate-skipping loops with careful bounds checks
- Awareness that sorting mutates the caller's slice — consider `slices.Clone` if that matters

## Interview Follow-ups

- Generalise to **k-Sum** (e.g. Four Sum). What is the complexity?
- What if you must not modify the input and can't afford a copy?
- Return the triplet closest to a target instead of exactly zero ("3Sum Closest").
- How would you count triplets with sum less than a target?

## Run the Tests

```bash
cd 14-three-sum
go test -v
```
