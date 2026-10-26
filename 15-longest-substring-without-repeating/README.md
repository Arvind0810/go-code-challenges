# 15 · Longest Substring Without Repeating Characters

**Difficulty:** Medium · **Topics:** Strings, Sliding Window, Hash Map · **Asked at:** Amazon, Google, Bloomberg, Adobe

## Problem

Given a string `s`, return the length of the longest **substring** (contiguous run of characters) that contains no repeated characters.

Treat `s` as a sequence of Unicode characters (runes), not bytes — `"héllo"` has 5 characters.

## Examples

```
Input:  s = "abcabcbb"
Output: 3
Explanation: "abc" is the longest window with all distinct characters.
```

```
Input:  s = "bbbbb"
Output: 1
Explanation: Every substring longer than 1 repeats 'b'.
```

```
Input:  s = "pwwkew"
Output: 3
Explanation: "wke". Note that "pwke" is a subsequence, not a substring.
```

## Constraints

- `0 <= len(s) <= 5 * 10^4`
- `s` may contain letters, digits, symbols, spaces and non-ASCII runes.

## Function Signature

```go
func LengthOfLongestSubstring(s string) int
```

## Hints

<details><summary>Hint 1</summary>

If `s[i..j]` has no repeats, then any substring inside it also has no repeats. You never need to move the left edge backwards.

</details>

<details><summary>Hint 2</summary>

Maintain a window `[left, right]` and a map from character to its most recent index.

</details>

<details><summary>Hint 3</summary>

For each new rune at `right`: if it was last seen at index `≥ left`, jump `left` to `lastSeen + 1`. Record the new index, then update `best = max(best, right - left + 1)`.

</details>

## Approach Guide

**Brute force — O(n³):** enumerate every substring and check whether its characters are unique with a set.

**Sliding window with a set — O(n):** extend `right`; while the new character is already in the set, remove `s[left]` and advance `left`. Each character enters and leaves the window at most once.

**Optimal — sliding window with last-seen index — O(n):**
1. Convert to `[]rune` (or iterate with care) so indices count characters.
2. Keep `lastSeen map[rune]int`, `left := 0`, `best := 0`.
3. For each `right`, if `lastSeen[r] >= left`, set `left = lastSeen[r] + 1`.
4. Store `lastSeen[r] = right` and update `best`.

## Target Complexity

- **Time:** O(n)
- **Space:** O(min(n, alphabet size))

## Go Concepts You'll Practise

- Strings vs `[]rune` — `len(s)` counts bytes, not characters
- `for i, r := range s` yields byte offsets, not rune indices
- `map[rune]int` and the comma-ok idiom `idx, ok := m[r]`
- Optional optimisation: a `[128]int` array for pure-ASCII input

## Interview Follow-ups

- Return the substring itself, not just its length.
- Longest substring with **at most k distinct** characters.
- How would you speed this up if the input is guaranteed ASCII?
- What changes if the input is a stream you can only read once?

## Run the Tests

```bash
cd 15-longest-substring-without-repeating
go test -v
```
