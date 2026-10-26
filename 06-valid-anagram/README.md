# 06 · Valid Anagram

**Difficulty:** Easy · **Topics:** Hash Map, Counting, Strings · **Asked at:** Amazon, Bloomberg, Microsoft, Uber

## Problem

Two strings are **anagrams** if one can be formed by rearranging the characters of the other, using every character exactly once.

Implement two functions:

1. `IsAnagram(s, t)` — both strings contain **only lowercase English letters** `a`–`z`. Return `true` if `t` is an anagram of `s`.
2. `IsAnagramUnicode(s, t)` — the strings may contain **any Unicode characters** (accents, CJK, emoji, spaces, punctuation). Compare rune by rune, case-sensitively, with no normalisation. Return `true` if `t` is an anagram of `s`.

## Examples

```text
Input:  s = "anagram", t = "nagaram"
Output: true
```

```text
Input:  s = "rat", t = "car"
Output: false
Explanation: 't' is not in "car" and 'c' is not in "rat".
```

```text
Input:  s = "été", t = "éét"   (IsAnagramUnicode)
Output: true
Explanation: both contain two 'é' and one 't'.
```

## Constraints

- `0 <= len(s), len(t) <= 5 * 10^4` bytes
- `IsAnagram`: `s` and `t` contain only `'a'`–`'z'`.
- `IsAnagramUnicode`: `s` and `t` are valid UTF-8.

## Function Signature

```go
func IsAnagram(s, t string) bool

func IsAnagramUnicode(s, t string) bool
```

## Hints

<details><summary>Hint 1</summary>

If the lengths differ, can they possibly be anagrams? What do two anagrams have in common if you count each character?

</details>

<details><summary>Hint 2</summary>

With only 26 possible letters, a fixed-size array `[26]int` indexed by `c - 'a'` is a perfect counter. For Unicode, use a `map[rune]int`.

</details>

<details><summary>Hint 3</summary>

Increment the counter for every character in `s` and decrement for every character in `t`. They are anagrams if and only if every count ends at zero. For the array version, you can compare `[26]int` values directly with `==`.

</details>

## Approach Guide

**Brute force:** sort both strings (as `[]byte` or `[]rune`) and compare. O(n log n) time, O(n) space. Easy to explain, good as a first answer.

**Optimal for `a`–`z`:**

1. If `len(s) != len(t)`, return `false`.
2. Create `var counts [26]int`.
3. For each byte `c` in `s`, `counts[c-'a']++`; for each byte in `t`, `counts[c-'a']--`.
4. Return `true` only if all counts are zero.

**Unicode version:** the same idea, but iterate with `for _, r := range s` (by rune) and count in a `map[rune]int`. Decrement for `t`; if any count goes negative, return `false` early. Note that the byte-length check is still a valid shortcut — anagrams have the same multiset of runes, hence the same UTF-8 byte length.

## Target Complexity

- **Time:** O(n)
- **Space:** O(1) for `[26]int` (fixed size); O(k) for the map, where `k` is the number of distinct runes

## Go Concepts You'll Practise

- Arrays vs slices: `[26]int` is a value type, zero-initialised, and comparable with `==`.
- Byte arithmetic: `s[i] - 'a'` yields a `byte` index.
- `range` over a string yields runes; indexing yields bytes.
- `map[rune]int` with zero-value increments.
- Early return to keep nesting shallow.

## Interview Follow-ups

1. Group a list of words into anagram groups (Group Anagrams). What makes a good map key?
2. Should `"é"` (precomposed) and `"e" + U+0301` be considered equal? How would you handle that? (`golang.org/x/text/unicode/norm`)
3. Find all start indices of anagrams of `p` inside `s` (sliding window).
4. Why is `[26]int` faster than `map[byte]int` in practice?

## Run the Tests

```bash
cd 06-valid-anagram
go test -v
```
