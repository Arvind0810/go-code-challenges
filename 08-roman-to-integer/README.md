# 08 · Roman to Integer

**Difficulty:** Easy · **Topics:** Strings, Hash Map, Parsing · **Asked at:** Amazon, Microsoft, Adobe, Bloomberg

## Problem

Roman numerals use seven symbols:

| Symbol | Value |
|---|---|
| `I` | 1 |
| `V` | 5 |
| `X` | 10 |
| `L` | 50 |
| `C` | 100 |
| `D` | 500 |
| `M` | 1000 |

Numerals are usually written largest to smallest, left to right, and their values are added: `XII` = 10 + 1 + 1 = 12.

Six **subtractive** pairs place a smaller symbol before a larger one to subtract it:

- `IV` = 4, `IX` = 9
- `XL` = 40, `XC` = 90
- `CD` = 400, `CM` = 900

Given a valid Roman numeral `s`, return its integer value.

## Examples

```text
Input:  s = "III"
Output: 3
Explanation: 1 + 1 + 1.
```

```text
Input:  s = "LVIII"
Output: 58
Explanation: L = 50, V = 5, III = 3.
```

```text
Input:  s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90, IV = 4.
```

## Constraints

- `1 <= len(s) <= 15`
- `s` contains only `I, V, X, L, C, D, M`.
- `s` is a valid Roman numeral in the range `[1, 3999]`.

## Function Signature

```go
func RomanToInt(s string) int
```

## Hints

<details><summary>Hint 1</summary>

When is a symbol subtracted instead of added? Compare it with the symbol to its right.

</details>

<details><summary>Hint 2</summary>

Walk the string once. If the current symbol's value is smaller than the next symbol's value, subtract it; otherwise add it.

</details>

<details><summary>Hint 3</summary>

Store values in a lookup (`map[byte]int` or a `switch`). For `i` in `0..len(s)-1`: `v := val(s[i])`; if `i+1 < len(s) && v < val(s[i+1])` then `total -= v` else `total += v`.

</details>

## Approach Guide

**Brute force / explicit pairs:** scan left to right; at each position first check whether the next two characters form one of the six subtractive pairs (`"CM"`, `"IV"`, …). If so, add the pair's value and skip two; otherwise add the single symbol. Correct, but needs a second lookup table and more branching.

**Optimal (compare with neighbour):**

1. Build a symbol → value lookup.
2. Iterate over indices of `s`.
3. If the current value is less than the next value, it is part of a subtractive pair — subtract it.
4. Otherwise add it.
5. Return the total.

An equivalent variant walks **right to left**, keeping track of the previous (right-hand) value, and subtracts whenever the current value is smaller than it.

## Target Complexity

- **Time:** O(n) — and since `n <= 15`, effectively O(1)
- **Space:** O(1)

## Go Concepts You'll Practise

- Indexing a string yields a `byte`, which is perfect for ASCII-only input.
- `map[byte]int` literals vs a `switch` on a byte vs an array lookup `[256]int` — trade-offs in speed and readability.
- Bounds checks (`i+1 < len(s)`) to avoid `index out of range` panics.
- Package-level `var` for lookup tables so they are built once.

## Interview Follow-ups

1. Implement the reverse: **Integer to Roman**.
2. How would you **validate** input and return an `error` for strings like `"IIII"`, `"VX"` or `"IC"`?
3. Which is faster here: a `map[byte]int`, a `switch`, or a `[256]int` array? Why?
4. How would you extend this to numbers above 3999 (vinculum/overline notation)?

## Run the Tests

```bash
cd 08-roman-to-integer
go test -v
```
