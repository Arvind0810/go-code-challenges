# 20 · Longest Consecutive Sequence

**Difficulty:** Medium · **Topics:** Arrays, Hash Set, Union-Find · **Asked at:** Google, Amazon, Meta, Spotify

## Problem

Given an unsorted integer slice `nums`, return the length of the longest run of **consecutive integers** that can be formed from its elements (e.g. `4, 5, 6, 7`).

The elements do not need to be adjacent in the input. Duplicates count only once. You must solve it in **O(n)** time.

## Examples

```
Input:  nums = [100, 4, 200, 1, 3, 2]
Output: 4
Explanation: The longest consecutive run is 1, 2, 3, 4.
```

```
Input:  nums = [0, 3, 7, 2, 5, 8, 4, 6, 0, 1]
Output: 9
Explanation: 0 through 8. The duplicate 0 does not extend the run.
```

```
Input:  nums = [1, 0, 1, 2]
Output: 3
```

## Constraints

- `0 <= len(nums) <= 10^5`
- `-10^9 <= nums[i] <= 10^9`
- Required time complexity: O(n).

## Function Signature

```go
func LongestConsecutive(nums []int) int
```

## Hints

<details><summary>Hint 1</summary>

Sorting makes this easy but costs O(n log n). What data structure gives O(1) "does x exist?" checks?

</details>

<details><summary>Hint 2</summary>

Put every number in a set. Only start counting from a number that is the **beginning** of a run.

</details>

<details><summary>Hint 3</summary>

`x` starts a run iff `x-1` is not in the set. From each start, count `x+1, x+2, ...` while they exist. Each number is visited by at most one counting loop, so the total is O(n).

</details>

## Approach Guide

**Brute force — O(n³) / O(n²):** for each number, repeatedly scan the slice to see if `x+1`, `x+2`, ... exist.

**Sort — O(n log n):**
1. Sort, then walk once, extending the current run when `nums[i] == nums[i-1] + 1`, ignoring equal neighbours, and resetting otherwise.

**Optimal — hash set, O(n):**
1. Insert all values into `set := make(map[int]struct{}, len(nums))`.
2. For each value `x` in the set, skip it if `x-1` is in the set.
3. Otherwise count upward from `x` while `x+length` is in the set.
4. Track the maximum length.

Iterate over the **set**, not the original slice, so heavy duplication can't cause repeated work.

## Target Complexity

- **Time:** O(n) average
- **Space:** O(n)

## Go Concepts You'll Practise

- Sets in Go: `map[int]struct{}` (zero-byte values) vs `map[int]bool`
- Membership test with the comma-ok idiom: `_, ok := set[x]`
- Pre-sizing maps with `make(map[K]V, hint)` to reduce rehashing
- Why the inner `for` loop doesn't make the algorithm O(n²) — amortised analysis

## Interview Follow-ups

- Return the sequence itself, not just its length.
- Solve it with Union-Find. What are the trade-offs?
- What if the numbers arrive as a stream and you must report the current longest run after each one?
- What if the data is too large for memory and must be processed in chunks?

## Run the Tests

```bash
cd 20-longest-consecutive-sequence
go test -v
```
