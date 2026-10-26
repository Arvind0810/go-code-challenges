# 11 · Best Time to Buy and Sell Stock

**Difficulty:** Easy · **Topics:** Arrays, Greedy, One Pass · **Asked at:** Amazon, Google, Microsoft, Bloomberg

## Problem

You are given a slice `prices` where `prices[i]` is the price of a stock on day `i`.

You may complete **at most one transaction**: choose a single day to buy one share and a **later** day to sell it. Return the maximum profit you can achieve. If no profit is possible, return `0`.

## Examples

```
Input:  prices = [7, 1, 5, 3, 6, 4]
Output: 5
Explanation: Buy on day 1 (price 1), sell on day 4 (price 6). Profit = 6 - 1 = 5.
Selling on day 0 (price 7) is not allowed — you must buy before you sell.
```

```
Input:  prices = [7, 6, 4, 3, 1]
Output: 0
Explanation: Prices only go down, so the best choice is not to trade.
```

```
Input:  prices = [2, 4, 1]
Output: 2
Explanation: Buy at 2, sell at 4. The later low of 1 has no higher price after it.
```

## Constraints

- `0 <= len(prices) <= 10^5`
- `0 <= prices[i] <= 10^4`
- An empty or single-element slice yields `0`.

## Function Signature

```go
func MaxProfit(prices []int) int
```

## Hints

<details><summary>Hint 1</summary>

For any sell day, which buy day gives the best profit? Only one past value matters.

</details>

<details><summary>Hint 2</summary>

Walk the slice once while remembering the cheapest price seen so far.

</details>

<details><summary>Hint 3</summary>

At each day: `profit = price - minSoFar`; update the best profit, then update `minSoFar = min(minSoFar, price)`.

</details>

## Approach Guide

**Brute force — O(n²):** try every pair `(i, j)` with `i < j` and track the largest `prices[j] - prices[i]`. Correct, but too slow for 10⁵ elements.

**Optimal — one pass, O(n):**
1. Keep `minPrice` (cheapest buy so far) and `best` (best profit so far, starting at 0).
2. For each price, compute what you would earn selling today against `minPrice` and update `best`.
3. Lower `minPrice` if today is cheaper.

This is a greedy/DP hybrid: the state you carry forward is just "the best buy price available to any future day".

## Target Complexity

- **Time:** O(n)
- **Space:** O(1)

## Go Concepts You'll Practise

- `for _, p := range prices` iteration over a slice
- Built-in `min` and `max` functions (Go 1.21+)
- Handling nil vs empty slices — `range` over nil is a safe no-op
- Choosing sensible zero values so edge cases need no special branch

## Interview Follow-ups

- What if you can make **unlimited** transactions (buy/sell multiple times, holding at most one share)?
- What if you are limited to **at most k** transactions?
- What if there is a **cooldown** day after each sale, or a transaction fee?
- How would you also return the buy and sell day indices?

## Run the Tests

```bash
cd 11-best-time-to-buy-sell-stock
go test -v
```
