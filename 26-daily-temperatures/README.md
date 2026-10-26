# 26 · Daily Temperatures

**Difficulty:** Medium · **Topics:** Stack, Monotonic Stack, Arrays · **Asked at:** Amazon, Meta, Google, Uber, Salesforce

## Problem

Given a slice `temperatures` of daily temperatures, return a slice `answer` of the same length where `answer[i]` is the number of days you have to wait after day `i` to get a **strictly warmer** temperature. If there is no future warmer day, `answer[i] = 0`.

Do not modify the input slice.

## Examples

```
Input:  [73, 74, 75, 71, 69, 72, 76, 73]
Output: [ 1,  1,  4,  2,  1,  1,  0,  0]
Explanation: Day 2 (75) waits until day 6 (76): 6 - 2 = 4 days.
```

```
Input:  [30, 40, 50, 60]
Output: [ 1,  1,  1,  0]
```

```
Input:  [70, 70, 71]
Output: [ 2,  1,  0]
Explanation: An equal temperature is not warmer.
```

## Constraints

- `0 <= len(temperatures) <= 10^5`
- `30 <= temperatures[i] <= 100`
- The result must have the same length as the input (an empty input gives an empty result).

## Function Signature

```go
func DailyTemperatures(temperatures []int) []int
```

## Hints

<details><summary>Hint 1</summary>

The brute force is two nested loops. Notice the wasted work: once you know day `j` is warmer than day `i`, every day between them that was colder than day `i` is never anyone's answer again... or is it?

</details>

<details><summary>Hint 2</summary>

Keep a stack of **indices** of days still waiting for a warmer day. The temperatures at those indices are always in non-increasing order from bottom to top — a *monotonic stack*.

</details>

<details><summary>Hint 3</summary>

For each new day `i`: while the stack is non-empty and `temps[i] > temps[top]`, pop `top` and set `answer[top] = i - top`. Then push `i`.

</details>

## Approach Guide

**Brute force:** For each `i`, scan forward until you find a warmer day. O(n²) time, O(1) extra space.

**Optimal (monotonic decreasing stack):**
1. `answer := make([]int, n)` (zeros by default — the "never warmer" case is free).
2. `stack := []int{}` of indices.
3. For each `i, t`:
   - while stack not empty and `t > temps[stack top]`: pop `j`, `answer[j] = i - j`.
   - push `i`.
4. Return `answer`.

Each index is pushed and popped at most once → linear time.

## Target Complexity

- Time: **O(n)**
- Space: **O(n)** for the stack

## Go Concepts You'll Practise

- `make([]int, n)` vs. `var s []int` — length vs. capacity, and why zero values give you the default answer.
- Using a slice as a stack: `stack[len(stack)-1]`, `stack = stack[:len(stack)-1]`.
- Pre-sizing capacity (`make([]int, 0, n)`) to avoid repeated reallocations.
- `range` over a slice yielding index and value.

## Interview Follow-ups

1. Can you solve it in O(1) extra space by iterating from the right and jumping using already-computed answers?
2. Next Greater Element I / II (the circular-array variant).
3. Largest Rectangle in Histogram — the classic hard monotonic-stack problem.
4. Online Stock Span — the same idea as a streaming API.

## Run the Tests

```bash
cd 26-daily-temperatures
go test -v
```
