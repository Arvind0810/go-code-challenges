# 01 · FizzBuzz

**Difficulty:** Easy · **Topics:** Loops, Conditionals, Strings · **Asked at:** Amazon, Microsoft, Bloomberg, many phone screens

## Problem

Given an integer `n`, return a slice of strings for the numbers `1` through `n` (inclusive) where:

- numbers divisible by both 3 and 5 become `"FizzBuzz"`,
- numbers divisible by 3 only become `"Fizz"`,
- numbers divisible by 5 only become `"Buzz"`,
- every other number becomes its decimal string form (e.g. `"7"`).

If `n` is less than 1, return an empty slice.

## Examples

```text
Input:  n = 3
Output: ["1", "2", "Fizz"]
Explanation: 3 is divisible by 3.
```

```text
Input:  n = 5
Output: ["1", "2", "Fizz", "4", "Buzz"]
Explanation: 5 is divisible by 5.
```

```text
Input:  n = 15
Output: ["1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"]
Explanation: 15 is divisible by both 3 and 5.
```

## Constraints

- `-10^4 <= n <= 10^4`
- For `n < 1` the result is empty.

## Function Signature

```go
func FizzBuzz(n int) []string
```

## Hints

<details><summary>Hint 1</summary>

The order in which you check the conditions matters. What happens to 15 if you check "divisible by 3" first?

</details>

<details><summary>Hint 2</summary>

Check the combined case (divisible by 15) before the single cases. Use the `%` operator to test divisibility.

</details>

<details><summary>Hint 3</summary>

Pre-allocate the result with `make([]string, 0, n)`, loop `i` from 1 to `n`, use a `switch` with no tag (`switch { case i%15 == 0: ... }`), and convert plain numbers with `strconv.Itoa`.

</details>

## Approach Guide

**Brute force / straightforward:** loop from 1 to `n`, test each number against 15, 3 and 5 in that order, and append the matching string. This is already optimal in time — the problem is about writing clean code, not clever algorithms.

**Cleaner variant (extensible):** build the output by concatenation instead of special-casing 15:

1. Start with an empty string `s`.
2. If `i % 3 == 0`, append `"Fizz"`; if `i % 5 == 0`, append `"Buzz"`.
3. If `s` is still empty, use the number as a string.

This scales nicely when the interviewer adds a new rule (e.g. 7 → `"Bazz"`).

## Target Complexity

- **Time:** O(n)
- **Space:** O(n) for the output (O(1) extra)

## Go Concepts You'll Practise

- `for` loops — Go's only loop keyword.
- Tagless `switch` as a cleaner `if / else if` chain.
- `strconv.Itoa` vs `fmt.Sprint` (and why `string(rune(i))` is a bug — it produces a code point, not digits).
- Slice pre-allocation with `make([]T, 0, capacity)` to avoid repeated growth.
- Guarding against invalid input (`n < 1`) before allocating.

## Interview Follow-ups

1. How would you make the rules configurable (e.g. a slice of `{divisor, word}` pairs) without changing the loop?
2. What does `make([]string, 0, n)` panic on if you forget the `n < 1` guard?
3. Can you avoid the modulo operator entirely (hint: counters)? Why might that matter on some hardware?
4. How would you stream the output (e.g. via a channel or `io.Writer`) for very large `n`?

## Run the Tests

```bash
cd 01-fizzbuzz
go test -v
```
