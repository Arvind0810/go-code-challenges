# 24 · Search in Rotated Sorted Array

**Difficulty:** Medium · **Topics:** Binary Search, Arrays · **Asked at:** Amazon, Google, Meta, Microsoft, LinkedIn

## Problem

An ascending array of **distinct** integers has been rotated at some unknown pivot — e.g. `[0,1,2,4,5,6,7]` might become `[4,5,6,7,0,1,2]`.

Given the rotated array `nums` and an integer `target`, return the **index** of `target`, or `-1` if it is not in the array.

You must write an algorithm with **O(log n)** runtime.

## Examples

```
Input:  nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

```
Input:  nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

```
Input:  nums = [1], target = 0
Output: -1
```

```
Explanation: why a rotated array is still searchable — it's two sorted runs.
Picture of [4,5,6,7,0,1,2]:

        7
      6 |
    5 | |
  4 | | |           2
  | | | |       1   |
  | | | |   0   |   |
  ---------+-----------
  sorted   | sorted
  (left)   | (right)
```

## Constraints

- `0 <= len(nums) <= 5000`
- `-10^4 <= nums[i], target <= 10^4`
- All values of `nums` are **unique**.
- `nums` is an ascending array that may have been rotated.

## Function Signature

```go
func Search(nums []int, target int) int
```

## Hints

<details><summary>Hint 1</summary>

Pick any `mid`. At least one of the halves `[lo..mid]` or `[mid..hi]` is **fully sorted**. How can you tell which one by comparing just two numbers?

</details>

<details><summary>Hint 2</summary>

If `nums[lo] <= nums[mid]`, the left half is sorted. Then `target` lies in it exactly when `nums[lo] <= target < nums[mid]`. Otherwise the right half is sorted — apply the mirror check.

</details>

<details><summary>Hint 3</summary>

Alternative two-pass approach: binary search for the index of the minimum element (the rotation point) first, then run a normal binary search on the correct side.

</details>

## Approach Guide

**Brute force:** Linear scan. O(n) — correct but fails the requirement.

**Optimal (one-pass binary search):**
1. `lo, hi = 0, len-1`. While `lo <= hi`:
2. `mid = lo + (hi-lo)/2`. If `nums[mid] == target` return `mid`.
3. If the left half is sorted (`nums[lo] <= nums[mid]`):
   - if target is within `[nums[lo], nums[mid])`, search left (`hi = mid-1`), else search right.
4. Else the right half is sorted:
   - if target is within `(nums[mid], nums[hi]]`, search right (`lo = mid+1`), else search left.
5. Return `-1`.

## Target Complexity

- Time: **O(log n)**
- Space: **O(1)**

## Go Concepts You'll Practise

- Overflow-safe midpoint: `lo + (hi-lo)/2` (or `int(uint(lo+hi) >> 1)`, as `sort.Search` does internally).
- Off-by-one discipline with inclusive `lo <= hi` loops.
- `sort.Search` / `slices.BinarySearch` from the stdlib — and why they don't directly apply to a rotated array.
- Handling an empty slice without special-casing (`hi = -1` skips the loop).

## Interview Follow-ups

1. What if the array **may contain duplicates**? (Search in Rotated Sorted Array II — worst case becomes O(n). Why?)
2. Find the minimum element of a rotated array in O(log n).
3. How many times has the array been rotated?
4. Implement it with Go's `sort.Search` by cleverly defining the predicate.

## Run the Tests

```bash
cd 24-search-rotated-sorted-array
go test -v
```
