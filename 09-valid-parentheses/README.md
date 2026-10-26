# 09 · Valid Parentheses

**Difficulty:** Easy · **Topics:** Stack, Strings · **Asked at:** Google, Amazon, Meta, Microsoft, Bloomberg

## Problem

Given a string `s` containing only the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine whether it is **valid**.

A string is valid if:

1. Every opening bracket is closed by a bracket of the **same type**.
2. Brackets are closed in the **correct order** (the most recently opened bracket closes first).
3. Every closing bracket has a matching opening bracket.

The empty string is valid.

## Examples

```text
Input:  s = "()[]{}"
Output: true
```

```text
Input:  s = "([)]"
Output: false
Explanation: '[' is opened after '(' so it must close before ')'.
```

```text
Input:  s = "{[()]}"
Output: true
Explanation: each bracket closes in reverse order of opening.
```

## Constraints

- `0 <= len(s) <= 10^4`
- `s` consists only of `()[]{}`.

## Function Signature

```go
func IsValid(s string) bool
```

## Hints

<details><summary>Hint 1</summary>

When you meet a closing bracket, which opening bracket must it match? The first one you saw, or the most recent unmatched one?

</details>

<details><summary>Hint 2</summary>

"Most recent unmatched" is last-in-first-out — use a stack. In Go, a slice works as a stack: `append` to push, re-slice to pop.

</details>

<details><summary>Hint 3</summary>

Map each closer to its opener (`')': '('`, …). For each char: if it is an opener, push it. If it is a closer, return `false` if the stack is empty or the top isn't the matching opener; otherwise pop. At the end, return `len(stack) == 0`.

</details>

## Approach Guide

**Brute force:** repeatedly remove adjacent pairs `"()"`, `"[]"`, `"{}"` with `strings.ReplaceAll` until nothing changes; valid if the string ends empty. O(n²) time and lots of allocation — mention it, then improve.

**Optimal (stack):**

1. Create an empty stack (`[]byte` with capacity `len(s)`).
2. For each character:
   - opener → push;
   - closer → the stack must be non-empty and its top must be the matching opener; pop it, else return `false`.
3. After the loop, any leftover openers mean the string is invalid.

**Quick win:** an odd-length string can never be valid — return `false` immediately.

## Target Complexity

- **Time:** O(n)
- **Space:** O(n) for the stack in the worst case (e.g. `"(((((("`)

## Go Concepts You'll Practise

- Slices as stacks: `stack = append(stack, c)` to push, `stack[len(stack)-1]` to peek, `stack = stack[:len(stack)-1]` to pop.
- Guarding against `index out of range` when peeking an empty slice.
- `map[byte]byte` literals for pair lookup, or a `switch` with multiple case values.
- Pre-allocating capacity with `make([]byte, 0, len(s))`.

## Interview Follow-ups

1. If only `'('` and `')'` are allowed, can you solve it in O(1) space? (A counter.)
2. Return the **minimum number of insertions** needed to make the string valid.
3. Find the **length of the longest valid** parentheses substring.
4. How would you ignore other characters (e.g. validate brackets in source code), and what about brackets inside string literals?

## Run the Tests

```bash
cd 09-valid-parentheses
go test -v
```
