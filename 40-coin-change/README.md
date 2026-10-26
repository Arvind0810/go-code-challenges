# 40 · Coin Change

**Difficulty:** Medium · **Topics:** Dynamic Programming, Unbounded Knapsack, BFS · **Asked at:** Amazon, Google, Microsoft, Goldman Sachs, Walmart

## Problem

You are given a slice of coin denominations `coins` and a target `amount`. You have an **unlimited** supply of each coin.

Return the **fewest** number of coins needed to make up exactly `amount`. If it cannot be made with any combination of the coins, return `-1`. An `amount` of `0` needs `0` coins.

## Examples

```text
Input:  coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1.
```

```text
Input:  coins = [2], amount = 3
Output: -1
Explanation: only even amounts can be formed.
```

```text
Input:  coins = [1, 3, 4], amount = 6
Output: 2
Explanation: greedy (largest coin first) gives 4 + 1 + 1 = 3 coins,
but 3 + 3 = 2 coins is optimal — greedy does not work here.
```

## Constraints

- `1 <= len(coins) <= 12`
- `1 <= coins[i] <= 2^31 - 1`
- `0 <= amount <= 10^4`
- `coins` is not necessarily sorted.

## Function Signature

```go
func CoinChange(coins []int, amount int) int
```

## Hints

<details><summary>Hint 1</summary>

Greedy (always take the largest coin that fits) fails for `[1, 3, 4]` and `6`. Think about the **last** coin used in an optimal solution for `amount`.

</details>

<details><summary>Hint 2</summary>

If the last coin is `c`, then the rest is an optimal solution for `amount - c`. So `minCoins(a) = 1 + min over c of minCoins(a - c)`, for every coin `c <= a`.

</details>

<details><summary>Hint 3</summary>

Build a table `dp[0..amount]` bottom-up with `dp[0] = 0` and every other entry set to a sentinel "infinity" such as `amount + 1` (you never need more than `amount` coins). At the end, if `dp[amount]` is still the sentinel, return `-1`.

</details>

## Approach Guide

**State:** `dp[a]` = minimum number of coins that sum to exactly `a`.

**Transition:** `dp[a] = min(dp[a - c] + 1)` over all coins `c` with `c <= a` and `dp[a - c]` reachable.

**Base case:** `dp[0] = 0`. Unreachable amounts hold a sentinel (`amount + 1` or `math.MaxInt` — careful with overflow when adding 1).

**Answer:** `dp[amount]`, or `-1` if it is still the sentinel.

**1. Plain recursion (brute force):** `solve(a)`: if `a == 0` return 0; if `a < 0` return "impossible"; otherwise try every coin and take `1 + min(solve(a - c))`. Exponential — roughly O(k^(amount / smallest coin)) for k coins.

**2. Memoisation (top-down):** cache `solve(a)` in a slice of size `amount + 1` (use a separate "computed" marker or `-2` because `-1` is a valid answer). O(amount · k) time, O(amount) space plus recursion depth up to `amount`.

**3. Tabulation (bottom-up):** loop `a` from `1` to `amount`, inner loop over coins, apply the transition. O(amount · k) time, O(amount) space, no recursion.

**4. Space:** the table is already 1D and every `dp[a]` may depend on any earlier entry, so O(amount) is the floor here. (Contrast with House Robber, where only two previous values matter.)

**Alternative — BFS:** treat each amount as a node and each coin as an edge; the shortest path from `0` to `amount` (or `amount` down to `0`) is the answer. Same complexity, sometimes faster in practice because it stops early.

## Target Complexity

- **Time:** O(amount × len(coins))
- **Space:** O(amount)

## Go Concepts You'll Practise

- `make([]int, amount+1)` and filling with a sentinel in a loop (Go has no array-fill builtin; `slices.Repeat` or a loop are the options).
- Choosing a safe "infinity": `amount + 1` avoids the overflow of `math.MaxInt + 1`.
- Nested `for range` loops and the built-in `min` (Go 1.21+).
- Distinguishing "not computed yet" from "impossible" in a memo table.

## Interview Follow-ups

1. **Coin Change II — number of ways:** count the distinct combinations that make `amount`. (Loop coins in the **outer** loop so each combination is counted once, not each permutation: `ways[a] += ways[a - c]`.)
2. Return the actual coins used, not just the count. (Store the last coin chosen for each `dp[a]` and walk back.)
3. What if each coin can be used **at most once**? (0/1 knapsack — iterate amounts in reverse.)
4. For which coin systems does greedy give the optimal answer? (Canonical systems like `[1, 5, 10, 25]`.)

## Run the Tests

```bash
cd 40-coin-change
go test -v
```
