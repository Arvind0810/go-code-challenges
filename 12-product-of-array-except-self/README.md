# 12 · Product of Array Except Self

**Difficulty:** Medium · **Topics:** Arrays, Prefix/Suffix Products · **Asked at:** Amazon, Meta, Apple, Microsoft

## Problem

Given an integer slice `nums`, return a new slice `answer` such that `answer[i]` equals the product of every element of `nums` **except** `nums[i]`.

You must solve it in **O(n)** time and **without using division**. The input must not be modified.

## Examples

```
Input:  nums = [1, 2, 3, 4]
Output: [24, 12, 8, 6]
Explanation: answer[0] = 2*3*4 = 24, answer[1] = 1*3*4 = 12, and so on.
```

```
Input:  nums = [-1, 1, 0, -3, 3]
Output: [0, 0, 9, 0, 0]
Explanation: Every product except the one skipping the 0 contains a 0.
```

```
Input:  nums = [5, 7]
Output: [7, 5]
```

## Constraints

- `0 <= len(nums) <= 10^5`
- `-30 <= nums[i] <= 30`
- Every prefix and suffix product fits in an `int`.
- An empty input returns an empty (or nil) slice.
- A single-element input returns `[1]` (the empty product).

## Function Signature

```go
func ProductExceptSelf(nums []int) []int
```

## Hints

<details><summary>Hint 1</summary>

`answer[i]` = (product of everything to the left of `i`) × (product of everything to the right of `i`).

</details>

<details><summary>Hint 2</summary>

Compute all left products in one forward pass and all right products in one backward pass.

</details>

<details><summary>Hint 3</summary>

Store left products directly in the output slice. Then walk backwards with a single running `right` variable, multiplying `answer[i] *= right` and then `right *= nums[i]`.

</details>

## Approach Guide

**Brute force — O(n²):** for each index, multiply every other element. Simple but quadratic.

**Division trick (not allowed):** multiply everything then divide by `nums[i]` — breaks on zeros and violates the rule.

**Optimal — prefix × suffix, O(n):**
1. Allocate `answer` of length `n`. Set `answer[0] = 1`, and for `i ≥ 1`, `answer[i] = answer[i-1] * nums[i-1]` (prefix products).
2. Keep `right := 1`. Iterate `i` from `n-1` down to `0`: `answer[i] *= right`, then `right *= nums[i]`.
3. Return `answer`. The output slice does not count as extra space, so auxiliary space is O(1).

## Target Complexity

- **Time:** O(n)
- **Space:** O(1) extra (excluding the output slice)

## Go Concepts You'll Practise

- `make([]int, n)` to pre-size an output slice
- Reverse iteration with `for i := len(nums) - 1; i >= 0; i--`
- Never mutating a caller's slice — slices share their backing array
- The empty product convention (`1`) as an initial value

## Interview Follow-ups

- How would you handle the case where division **is** allowed but the input may contain zeros?
- What if the products could overflow `int64`? (Think `math/big` or modular arithmetic.)
- Can you do it with exactly one extra variable beyond the output?
- How does this prefix/suffix pattern apply to "trapping rain water"?

## Run the Tests

```bash
cd 12-product-of-array-except-self
go test -v
```
