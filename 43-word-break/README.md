# 43 · Word Break

**Difficulty:** Medium · **Topics:** Dynamic Programming, Hash Set, Strings · **Asked at:** Amazon, Google, Meta, Uber, Microsoft

## Problem

Given a string `s` and a dictionary of strings `wordDict`, return `true` if `s` can be segmented into a sequence of one or more dictionary words.

- The same dictionary word may be reused any number of times.
- Matching is exact and case-sensitive.

## Examples

```text
Input:  s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: "leetcode" = "leet" + "code".
```

```text
Input:  s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: "apple" + "pen" + "apple" — words can be reused.
```

```text
Input:  s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
```

## Constraints

- `1 <= len(s) <= 300`
- `0 <= len(wordDict) <= 1000`
- `1 <= len(wordDict[i]) <= 20`
- All words in `wordDict` are unique.
- The tests include an adversarial input (`"aaaa...ab"`) that makes un-memoised recursion run effectively forever.

## Function Signature

```go
func WordBreak(s string, wordDict []string) bool
```

## Hints

<details><summary>Hint 1</summary>

Greedy (always take the longest matching word) fails — try `"goalspecial"` with `["go", "goal", "goals", "special"]`. You need to explore alternatives.

</details>

<details><summary>Hint 2</summary>

The question "can `s[i:]` be segmented?" is asked over and over with the same `i`. There are only `len(s)+1` distinct sub-problems — cache them.

</details>

<details><summary>Hint 3</summary>

Bottom-up: `dp[i]` = "the prefix `s[:i]` can be segmented". `dp[0] = true`. `dp[i]` is true if there is a `j < i` with `dp[j]` true and `s[j:i]` in the dictionary (a `map[string]bool`). Only check `j >= i - maxWordLen`.

</details>

## Approach Guide

**Brute force (recursion):** at position `i`, try every word that is a prefix of `s[i:]` and recurse on the rest. Worst case O(2ⁿ) — the `"aaa...ab"` input explodes.

**Memoised recursion:** same recursion, but store `memo[i]` (true/false/unknown). Each start index is solved once.

**Bottom-up DP:**

```text
words := set(wordDict); maxLen := longest word
dp[0] = true
for i in 1..n:
  for j in max(0, i-maxLen) .. i-1:
    if dp[j] and s[j:i] in words: dp[i] = true; break
return dp[n]
```

**Alternative:** BFS over start indices with a visited set, or a trie of words to walk forward from each reachable index without building substrings.

## Target Complexity

- **Time:** O(n · L) substring checks, where `L` = max word length (each check costs O(L) to hash) → O(n · L²)
- **Space:** O(n + total dictionary size)

## Go Concepts You'll Practise

- `map[string]bool` / `map[string]struct{}` as a set.
- String slicing `s[j:i]` (shares the underlying bytes — no copy).
- `[]bool` DP tables and early `break`.
- Built-in `max` for bounds clamping.

## Interview Follow-ups

1. Word Break II: return **all** possible sentences (LeetCode 140). How does output size affect complexity?
2. How would a trie speed this up when the dictionary is huge?
3. Return the segmentation with the **fewest** words.
4. If the dictionary is a remote service (each lookup is an RPC), how do you minimise calls?

## Run the Tests

```bash
cd 43-word-break
go test -v
```
