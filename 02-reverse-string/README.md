# 02 · Reverse String

**Difficulty:** Easy · **Topics:** Strings, Unicode, Two Pointers · **Asked at:** Google, Amazon, Uber, general Go screens

## Problem

Given a string `s`, return a new string with its characters in reverse order.

The input is UTF-8 and may contain multi-byte characters such as accented letters (`é`), CJK characters (`世`) and emoji (`🌍`). Each Unicode code point must stay intact — reversing the raw bytes would corrupt them.

## Examples

```text
Input:  s = "hello"
Output: "olleh"
```

```text
Input:  s = "héllo"
Output: "olléh"
Explanation: "é" is 2 bytes in UTF-8; it must be moved as one unit.
```

```text
Input:  s = "Hi 👋🌍"
Output: "🌍👋 iH"
Explanation: each emoji is a single 4-byte code point.
```

## Constraints

- `0 <= len(s) <= 10^5` bytes
- `s` is valid UTF-8.
- Combining sequences (e.g. `e` + `U+0301`) and multi-code-point emoji (flags, skin tones) are out of scope — reverse by code point (rune).

## Function Signature

```go
func Reverse(s string) string
```

## Hints

<details><summary>Hint 1</summary>

What does `len("é")` return in Go? What does indexing `s[0]` give you — a character or a byte?

</details>

<details><summary>Hint 2</summary>

Convert the string to a `[]rune` so each element is one Unicode code point, then reverse that slice.

</details>

<details><summary>Hint 3</summary>

Use two indices `i := 0`, `j := len(r)-1`; while `i < j`, swap `r[i], r[j] = r[j], r[i]`, then move both inward. Finish with `string(r)`.

</details>

## Approach Guide

**Naive (buggy) approach:** iterate bytes from the end and build a new string. This works for ASCII but splits multi-byte characters, producing invalid UTF-8 for `"héllo"`.

**Correct approach:**

1. Convert `s` to `[]rune` — Go decodes UTF-8 for you.
2. Reverse the rune slice in place with two pointers swapping from the ends toward the middle.
3. Convert back with `string(runes)`.

**Alternative:** walk the string with `utf8.DecodeLastRuneInString`, writing each rune into a `strings.Builder` (pre-sized with `Grow(len(s))`). This avoids the intermediate `[]rune` (4 bytes per rune) allocation.

## Target Complexity

- **Time:** O(n)
- **Space:** O(n) for the result (strings are immutable in Go, so a copy is unavoidable)

## Go Concepts You'll Practise

- Strings are immutable byte sequences; `len(s)` counts bytes, not characters.
- `rune` is an alias for `int32` and represents one Unicode code point.
- `[]rune(s)` / `string(runes)` conversions and their allocation cost.
- `for i, r := range s` iterates by rune, yielding byte offsets.
- Tuple assignment for swaps: `a, b = b, a`.
- `unicode/utf8` and `strings.Builder` for allocation-aware string building.

## Interview Follow-ups

1. How would you handle grapheme clusters like `"🇮🇳"` or `"é"` written as `e` + combining accent?
2. Can you reverse a `[]byte` that holds ASCII in place with O(1) extra space? Why can't you do the same with a Go `string`?
3. Reverse the order of words in a sentence instead of characters.
4. What is the memory overhead of `[]rune(s)` compared with the original string?

## Run the Tests

```bash
cd 02-reverse-string
go test -v
```
