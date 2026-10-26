# 17 · Top K Frequent Elements

**Difficulty:** Medium · **Topics:** Hash Map, Heap, Bucket Sort · **Asked at:** Amazon, Meta, Yelp, Oracle

## Problem

Given an integer slice `nums` and an integer `k`, return the `k` most frequent elements.

The answer is guaranteed to be unique — there are no ties at the boundary between the k-th and (k+1)-th most frequent elements. The returned elements may be in any order.

## Examples

```
Input:  nums = [1, 1, 1, 2, 2, 3], k = 2
Output: [1, 2]
Explanation: 1 appears 3 times, 2 appears twice, 3 once.
```

```
Input:  nums = [1], k = 1
Output: [1]
```

```
Input:  nums = [4, -1, -1, 4, 4, 7], k = 1
Output: [4]
Explanation: 4 appears 3 times, more than any other value.
```

## Constraints

- `1 <= len(nums) <= 10^5`
- `-10^4 <= nums[i] <= 10^4`
- `1 <= k <=` number of distinct values in `nums`
- The answer is unique.
- Aim for better than O(n log n).

## Function Signature

```go
func TopKFrequent(nums []int, k int) []int
```

## Hints

<details><summary>Hint 1</summary>

First count how many times each value appears. Then the problem is "pick the k largest counts".

</details>

<details><summary>Hint 2</summary>

Sorting the distinct values by count is O(m log m). Can you avoid a full sort? Consider a size-k min-heap, or note that a count can never exceed `n`.

</details>

<details><summary>Hint 3</summary>

Bucket sort: make `buckets := make([][]int, n+1)` where `buckets[c]` holds values that appear exactly `c` times. Walk buckets from `n` down to `1`, collecting values until you have `k`.

</details>

## Approach Guide

**Brute force — O(n log n):** count with a map, copy the distinct values into a slice, sort by descending count, take the first `k`.

**Min-heap — O(n log k):**
1. Count frequencies.
2. Push each `(value, count)` onto a min-heap keyed by count; if the heap grows past `k`, pop the smallest.
3. The heap now holds the answer.

**Bucket sort — O(n):**
1. Count frequencies.
2. Place each value into `buckets[count]`.
3. Iterate from the highest bucket down, appending values until `len(result) == k`.

## Target Complexity

- **Time:** O(n) with bucket sort (O(n log k) with a heap)
- **Space:** O(n)

## Go Concepts You'll Practise

- `map[int]int` counting with `counts[v]++` (zero value makes this safe)
- Implementing `heap.Interface` (`Len`, `Less`, `Swap`, `Push`, `Pop`) from `container/heap`
- Slices of slices as buckets: `make([][]int, n+1)`
- `slices.SortFunc` with a custom comparator for the brute-force version

## Interview Follow-ups

- What if there can be ties — how would you define and return a deterministic answer?
- Solve it with Quickselect for average O(n). What is its worst case?
- How would you maintain the top k over a never-ending stream?
- Top K Frequent **Words**, sorted by frequency then lexicographically.

## Run the Tests

```bash
cd 17-top-k-frequent-elements
go test -v
```
