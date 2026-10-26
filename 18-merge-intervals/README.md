# 18 · Merge Intervals

**Difficulty:** Medium · **Topics:** Arrays, Sorting, Intervals · **Asked at:** Google, Meta, Amazon, LinkedIn

## Problem

Given a slice of intervals where `intervals[i] = [start_i, end_i]`, merge all overlapping intervals and return a slice of the non-overlapping intervals that cover exactly the same ranges.

Intervals that **touch** (one ends where the next starts, e.g. `[1, 4]` and `[4, 5]`) count as overlapping. The result must be sorted by start.

## Examples

```
Input:  intervals = [[1, 3], [2, 6], [8, 10], [15, 18]]
Output: [[1, 6], [8, 10], [15, 18]]
Explanation: [1, 3] and [2, 6] overlap, so they merge into [1, 6].
```

```
Input:  intervals = [[1, 4], [4, 5]]
Output: [[1, 5]]
Explanation: Touching intervals are merged.
```

```
Input:  intervals = [[1, 4], [0, 4]]
Output: [[0, 4]]
Explanation: The input is not necessarily sorted.
```

## Constraints

- `0 <= len(intervals) <= 10^4`
- `len(intervals[i]) == 2`
- `0 <= start_i <= end_i <= 10^4`
- An empty input returns an empty (or nil) slice.

## Function Signature

```go
func Merge(intervals [][]int) [][]int
```

## Hints

<details><summary>Hint 1</summary>

If the intervals were sorted by start, which intervals could possibly overlap with a given one?

</details>

<details><summary>Hint 2</summary>

After sorting by start, an interval either overlaps the **last** merged interval or starts a new one. You never need to look further back.

</details>

<details><summary>Hint 3</summary>

Sort, then push the first interval into `merged`. For each next `[s, e]`: if `s <= last[1]`, set `last[1] = max(last[1], e)`; otherwise append `[s, e]`.

</details>

## Approach Guide

**Brute force — O(n²) or worse:** repeatedly scan for any pair that overlaps, merge them, and start over until no pair overlaps.

**Optimal — sort then sweep, O(n log n):**
1. Sort intervals by start (`slices.SortFunc`).
2. Initialise `merged` with a copy of the first interval.
3. For each subsequent interval, compare its start with the end of the last merged interval.
4. Overlap → extend the end with `max`. No overlap → append a new interval.

Watch out for the "contained" case: `[1, 10]` followed by `[2, 3]` must keep end `10`, not shrink to `3`.

## Target Complexity

- **Time:** O(n log n) — dominated by sorting
- **Space:** O(n) for the output (plus sort's internal stack)

## Go Concepts You'll Practise

- `slices.SortFunc` with a comparator like `func(a, b []int) int { return cmp.Compare(a[0], b[0]) }`
- Mutating the last element via `merged[len(merged)-1][1] = ...`
- Aliasing: appending an input sub-slice into the output shares memory — copy it if you will mutate it
- Working with `[][]int` and its nil/empty edge cases

## Interview Follow-ups

- **Insert Interval:** add one new interval into an already sorted, non-overlapping list in O(n).
- Find the minimum number of intervals to remove so the rest don't overlap.
- How many meeting rooms are needed for a set of meetings?
- What if intervals arrive as a stream and you must keep them merged at all times?

## Run the Tests

```bash
cd 18-merge-intervals
go test -v
```
