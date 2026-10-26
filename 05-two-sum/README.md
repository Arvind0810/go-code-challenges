# 05 · Two Sum

**Difficulty:** Easy · **Topics:** Arrays, Hash Map · **Asked at:** Google, Amazon, Meta, Apple, Adobe

## Problem

Given a slice of integers `nums` and an integer `target`, return the **indices** of the two numbers that add up to `target`.

- You may not use the same element twice (two equal *values* at different indices are fine).
- Return the indices in ascending order: `[i, j]` with `i < j`.
- If no such pair exists, return `nil`.
- When a pair exists, assume exactly one valid answer.

## Examples

```text
Input:  nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
Explanation: nums[0] + nums[1] = 2 + 7 = 9.
```

```text
Input:  nums = [3, 2, 4], target = 6
Output: [1, 2]
Explanation: 2 + 4 = 6. Using index 0 twice (3 + 3) is not allowed.
```

```text
Input:  nums = [3, 3], target = 6
Output: [0, 1]
```

## Constraints

- `0 <= len(nums) <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- At most one valid answer exists.

## Function Signature

```go
func TwoSum(nums []int, target int) []int
```

## Hints

<details><summary>Hint 1</summary>

For each number `x`, the number you need is fixed: `target - x`. How quickly can you check whether you've already seen it?

</details>

<details><summary>Hint 2</summary>

Use a map from value to index. A single pass is enough if you check before you insert.

</details>

<details><summary>Hint 3</summary>

For each `i, x`: if `j, ok := seen[target-x]; ok`, return `[]int{j, i}`. Otherwise store `seen[x] = i`. Return `nil` after the loop.

</details>

## Approach Guide

**Brute force:** try every pair `(i, j)` with `i < j` and check `nums[i] + nums[j] == target`. O(n²) time, O(1) space. Always say this out loud first in an interview.

**Sort + two pointers:** sort a copy of `(value, index)` pairs and walk inward from both ends. O(n log n) — useful if memory is tight, but you must track original indices.

**Optimal (one-pass hash map):**

1. Create `seen := map[int]int{}` (value → index).
2. For each index `i` with value `x`, compute `need := target - x`.
3. If `need` is in `seen`, you have the answer: the earlier index and `i`.
4. Otherwise record `x` → `i` and continue.

Checking *before* inserting is what prevents using the same element twice.

## Target Complexity

- **Time:** O(n)
- **Space:** O(n)

## Go Concepts You'll Practise

- The comma-ok idiom: `v, ok := m[key]` — distinguishing "missing" from "zero value".
- `make(map[int]int, len(nums))` capacity hints.
- `for i, x := range nums` over slices.
- Returning `nil` vs an empty slice, and how callers can treat them the same with `len`.

## Interview Follow-ups

1. The input is already sorted — can you solve it in O(1) extra space? (Two Sum II)
2. Return **all** unique pairs of values that sum to `target`.
3. Extend to Three Sum: find all unique triplets that sum to zero.
4. What if `nums` is a stream too large for memory, but you can make multiple passes?

## Run the Tests

```bash
cd 05-two-sum
go test -v
```
