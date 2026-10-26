# 25 · Koko Eating Bananas

**Difficulty:** Medium · **Topics:** Binary Search on the Answer · **Asked at:** Google, Amazon, Meta, Airbnb, DoorDash

## Problem

Koko has `n` piles of bananas; pile `i` has `piles[i]` bananas. The guards return in `h` hours.

Koko chooses an eating speed `k` (bananas per hour). Each hour she picks **one** pile and eats `k` bananas from it. If the pile has fewer than `k` bananas, she eats them all and **does not** move on to another pile during that hour.

Return the **minimum integer `k`** such that she can eat every banana within `h` hours.

## Examples

```
Input:  piles = [3,6,7,11], h = 8
Output: 4
Explanation: At k=4 the hours are ceil(3/4)+ceil(6/4)+ceil(7/4)+ceil(11/4)
             = 1 + 2 + 2 + 3 = 8 <= 8.  At k=3 it would take 10 hours.
```

```
Input:  piles = [30,11,23,4,20], h = 5
Output: 30
Explanation: One hour per pile, so k must be at least the biggest pile.
```

```
Input:  piles = [30,11,23,4,20], h = 6
Output: 23
```

## Constraints

- `1 <= len(piles) <= 10^4`
- `len(piles) <= h <= 10^9`
- `1 <= piles[i] <= 10^9`

## Function Signature

```go
func MinEatingSpeed(piles []int, h int) int
```

## Hints

<details><summary>Hint 1</summary>

What's the smallest `k` that could possibly work? The largest `k` you'd ever need? (Think: `1` and `max(piles)`.)

</details>

<details><summary>Hint 2</summary>

If speed `k` works, does `k+1` also work? That **monotonic** yes/no property means you can binary search over `k` instead of over an array.

</details>

<details><summary>Hint 3</summary>

Write `hoursNeeded(k)` = sum of `ceil(p / k)`. In integer math, `ceil(p/k) == (p + k - 1) / k`. Find the smallest `k` in `[1, max]` with `hoursNeeded(k) <= h`.

</details>

## Approach Guide

**Brute force:** Try `k = 1, 2, 3, ...` until `hoursNeeded(k) <= h`. O(max · n) — up to 10^13 operations. Too slow.

**Optimal (binary search on the answer):**
1. `lo = 1`, `hi = max(piles)`.
2. While `lo < hi`: `k = lo + (hi-lo)/2`.
3. If `hoursNeeded(k) <= h`, the answer is `k` or smaller → `hi = k`.
4. Otherwise `k` is too slow → `lo = k + 1`.
5. Return `lo`.

This "find the smallest value satisfying a monotonic predicate" template solves a whole family of problems.

## Target Complexity

- Time: **O(n · log(max(piles)))**
- Space: **O(1)**

## Go Concepts You'll Practise

- Integer ceiling division without `math.Ceil` and float conversion.
- `int` is 64-bit on modern platforms — why the sum of hours won't overflow here, and when it might on 32-bit.
- `slices.Max` (Go 1.21+) for the upper bound.
- `sort.Search(n, func(i int) bool)` — the stdlib's "smallest index where predicate is true"; try rewriting your loop with it.

## Interview Follow-ups

1. Capacity To Ship Packages Within D Days — same template, different predicate.
2. Split Array Largest Sum.
3. Can you tighten the lower bound? (Hint: `ceil(sum / h)`.)
4. How would the answer change if Koko could move to the next pile in the same hour?

## Run the Tests

```bash
cd 25-koko-eating-bananas
go test -v
```
