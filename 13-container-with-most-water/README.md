# 13 · Container With Most Water

**Difficulty:** Medium · **Topics:** Arrays, Two Pointers, Greedy · **Asked at:** Amazon, Google, Meta, Adobe

## Problem

You are given a slice `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of line `i` are `(i, 0)` and `(i, height[i])`.

Pick two lines that, together with the x-axis, form a container holding the most water. Return the maximum amount of water the container can store.

The area between lines `i` and `j` is `min(height[i], height[j]) * (j - i)`. You may not slant the container.

## Examples

```
Input:  height = [1, 8, 6, 2, 5, 4, 8, 3, 7]
Output: 49
Explanation: Lines at index 1 (height 8) and index 8 (height 7): min(8, 7) * (8 - 1) = 49.
```

```
Input:  height = [1, 1]
Output: 1
Explanation: min(1, 1) * 1 = 1.
```

```
Input:  height = [4, 3, 2, 1, 4]
Output: 16
Explanation: The two outer lines of height 4 are 4 apart: 4 * 4 = 16.
```

## Constraints

- `0 <= len(height) <= 10^5`
- `0 <= height[i] <= 10^4`
- Fewer than two lines cannot hold water — return `0`.

## Function Signature

```go
func MaxArea(height []int) int
```

## Hints

<details><summary>Hint 1</summary>

Start with the widest possible container. Any narrower container must make up for lost width with a taller short side.

</details>

<details><summary>Hint 2</summary>

Use two pointers at both ends and move one of them inward each step.

</details>

<details><summary>Hint 3</summary>

Always move the pointer at the **shorter** line. Moving the taller one can never increase the area, because the shorter line still caps the height while the width shrinks.

</details>

## Approach Guide

**Brute force — O(n²):** evaluate every pair `(i, j)` and keep the maximum area.

**Optimal — two pointers, O(n):**
1. `left := 0`, `right := len(height) - 1`, `best := 0`.
2. While `left < right`: compute the area, update `best`.
3. If `height[left] < height[right]`, increment `left`; otherwise decrement `right`.
4. Return `best`.

Why it is safe: the shorter line at the current width is already paired with the farthest partner it could ever have. No container using it can beat the current one, so we can discard it.

## Target Complexity

- **Time:** O(n)
- **Space:** O(1)

## Go Concepts You'll Practise

- Classic `for left < right` two-pointer loop
- Built-in `min` / `max` for clean area computation
- Guarding against short inputs without explicit length checks (the loop condition does it)

## Interview Follow-ups

- Prove why moving the taller pointer can never produce a better answer.
- How is this different from "Trapping Rain Water"? Solve that one too.
- What if you also had to return the pair of indices?
- What changes if the lines had different widths?

## Run the Tests

```bash
cd 13-container-with-most-water
go test -v
```
