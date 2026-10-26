# 10 · Merge Sorted Arrays

**Difficulty:** Easy · **Topics:** Arrays, Two Pointers, In-place · **Asked at:** Meta, Microsoft, Amazon, Apple, Bloomberg

## Problem

You are given two integer slices `nums1` and `nums2`, each sorted in **non-decreasing** order, and two integers `m` and `n`:

- the first `m` elements of `nums1` are the real values to merge,
- `nums1` has length `m + n`; its last `n` elements are zeros acting as free space,
- `nums2` has exactly `n` elements.

Merge `nums2` into `nums1` **in place** so that `nums1` becomes a single sorted slice of length `m + n`. The function returns nothing — modify `nums1` directly.

## Examples

```text
Input:  nums1 = [1, 2, 3, 0, 0, 0], m = 3, nums2 = [2, 5, 6], n = 3
Output: nums1 = [1, 2, 2, 3, 5, 6]
```

```text
Input:  nums1 = [1], m = 1, nums2 = [], n = 0
Output: nums1 = [1]
Explanation: nothing to merge.
```

```text
Input:  nums1 = [0], m = 0, nums2 = [1], n = 1
Output: nums1 = [1]
Explanation: nums1 has no real values; the single 0 is just free space.
```

## Constraints

- `len(nums1) == m + n`
- `len(nums2) == n`
- `0 <= m, n <= 200`
- `1 <= m + n <= 200`
- `-10^9 <= nums1[i], nums2[j] <= 10^9`

## Function Signature

```go
func Merge(nums1 []int, m int, nums2 []int, n int)
```

## Hints

<details><summary>Hint 1</summary>

Merging from the front would overwrite values of `nums1` you haven't used yet. Where is the free space?

</details>

<details><summary>Hint 2</summary>

Fill `nums1` from the back. The largest remaining element of either slice belongs at the last free position.

</details>

<details><summary>Hint 3</summary>

Use three indices: `i = m-1`, `j = n-1`, `k = m+n-1`. While `j >= 0`: if `i >= 0 && nums1[i] > nums2[j]`, write `nums1[i]` to `nums1[k]` and `i--`; otherwise write `nums2[j]` and `j--`. Always `k--`. When `j` runs out, the rest of `nums1` is already in place.

</details>

## Approach Guide

**Brute force:** copy `nums2` into the tail of `nums1` and sort the whole slice. O((m+n) log(m+n)) time. Works, but ignores that both inputs are already sorted.

**Extra buffer:** copy the first `m` elements of `nums1` into a new slice, then do a standard front-to-back two-pointer merge into `nums1`. O(m+n) time but O(m) extra space.

**Optimal (back-to-front, in place):**

1. Point `i` at the last real element of `nums1`, `j` at the last element of `nums2`, `k` at the last slot of `nums1`.
2. Compare `nums1[i]` and `nums2[j]`; write the larger to `nums1[k]` and move that pointer left.
3. Decrement `k` each step.
4. Stop when `nums2` is exhausted. Leftover `nums1` elements are already in the right place; leftover `nums2` elements must still be copied.

Because writes happen at `k >= i`, you never overwrite an unread `nums1` value.

## Target Complexity

- **Time:** O(m + n)
- **Space:** O(1)

## Go Concepts You'll Practise

- Slices are references to an underlying array — mutating `nums1` inside the function is visible to the caller.
- Why `nums1 = append(...)` inside the function would **not** be visible to the caller (the slice header is passed by value).
- Index arithmetic and loop conditions with multiple counters.
- The built-in `copy` for bulk-moving the leftover tail.

## Interview Follow-ups

1. Merge `k` sorted slices efficiently. (Min-heap with `container/heap`, O(N log k).)
2. Why does merging from the front fail without extra space? Walk through an example.
3. Merge two sorted linked lists — what changes?
4. If `nums1` had no spare capacity, how would you design the Go API (return a new slice, or `append` and return)?

## Run the Tests

```bash
cd 10-merge-sorted-arrays
go test -v
```
