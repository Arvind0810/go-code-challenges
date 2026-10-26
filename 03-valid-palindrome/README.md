# 03 · Valid Palindrome

**Difficulty:** Easy · **Topics:** Strings, Two Pointers, Unicode · **Asked at:** Meta, Microsoft, Amazon, Apple

## Problem

A phrase is a **palindrome** if, after converting all letters to lowercase and removing every character that is not a letter or a digit, it reads the same forward and backward.

Given a string `s`, return `true` if it is a palindrome, or `false` otherwise.

Letters and digits are defined by Go's `unicode.IsLetter` and `unicode.IsDigit`, so non-ASCII letters such as `é` count as letters.

## Examples

```text
Input:  s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
```

```text
Input:  s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
```

```text
Input:  s = " "
Output: true
Explanation: after cleaning, the string is empty, and an empty string reads the same both ways.
```

## Constraints

- `0 <= len(s) <= 2 * 10^5` bytes
- `s` is valid UTF-8.

## Function Signature

```go
func IsPalindrome(s string) bool
```

## Hints

<details><summary>Hint 1</summary>

You could build a cleaned copy and compare it with its reverse — but can you do it without building any new string?

</details>

<details><summary>Hint 2</summary>

Use two pointers, one at each end. Skip characters that are not letters or digits, then compare the two characters case-insensitively.

</details>

<details><summary>Hint 3</summary>

Convert to `[]rune` first so indices point at whole characters. Loop while `i < j`: advance `i` while `r[i]` is not alphanumeric, retreat `j` likewise, compare `unicode.ToLower(r[i])` with `unicode.ToLower(r[j])`, return `false` on mismatch, then `i++` and `j--`.

</details>

## Approach Guide

**Brute force:** build a filtered, lowercased `[]rune`, then check it equals its reverse. Simple and correct, but uses O(n) extra memory twice (filtered copy + reverse).

**Optimal (two pointers):**

1. Put `left` at the start and `right` at the end.
2. Move `left` forward past any non-alphanumeric character; move `right` backward likewise (never letting them cross).
3. Compare the lowercased characters. If they differ, it is not a palindrome.
4. Move both pointers inward and repeat until they meet.

For pure ASCII input you can work directly on the bytes for O(1) extra space. For Unicode, either convert to `[]rune` or decode from both ends with `utf8.DecodeRuneInString` / `utf8.DecodeLastRuneInString`.

## Target Complexity

- **Time:** O(n)
- **Space:** O(1) extra with byte/rune decoding from both ends; O(n) if you convert to `[]rune`

## Go Concepts You'll Practise

- The `unicode` package: `IsLetter`, `IsDigit`, `ToLower`.
- Bytes vs runes when indexing strings.
- Nested `for` loops with guard conditions (`for i < j && !ok(r[i])`).
- Why `strings.ToLower` + `strings.Map` is convenient but allocates.

## Interview Follow-ups

1. **Valid Palindrome II:** return `true` if the string can become a palindrome after deleting at most one character.
2. How would you do this in O(1) extra space for full Unicode input, without `[]rune`?
3. What changes if the input is a stream you can only read once?
4. Why can't you simply compare `strings.ToLower(s)` with its reverse?

## Run the Tests

```bash
cd 03-valid-palindrome
go test -v
```
