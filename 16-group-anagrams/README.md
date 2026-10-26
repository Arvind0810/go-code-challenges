# 16 · Group Anagrams

**Difficulty:** Medium · **Topics:** Strings, Hash Map, Sorting · **Asked at:** Amazon, Meta, Uber, Goldman Sachs

## Problem

Given a slice of strings `strs`, group the anagrams together and return the groups.

Two strings are anagrams if one is a rearrangement of the other's characters (same characters, same counts). Every input string must appear in exactly one group — including duplicates, which go in the same group.

The order of the groups, and the order of strings within each group, does not matter.

## Examples

```
Input:  strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
Output: [["bat"], ["nat", "tan"], ["ate", "eat", "tea"]]
Explanation: "eat", "tea", "ate" share the letters a, e, t; "tan" and "nat" share a, n, t.
```

```
Input:  strs = [""]
Output: [[""]]
Explanation: The empty string is its own group.
```

```
Input:  strs = ["a"]
Output: [["a"]]
```

## Constraints

- `0 <= len(strs) <= 10^4`
- `0 <= len(strs[i]) <= 100`
- `strs[i]` consists of lowercase English letters.

## Function Signature

```go
func GroupAnagrams(strs []string) [][]string
```

## Hints

<details><summary>Hint 1</summary>

All anagrams share a *canonical form*. If you could compute one key per word, grouping is just a map lookup.

</details>

<details><summary>Hint 2</summary>

Two candidate keys: the word's letters sorted (`"eat"` → `"aet"`), or a count of each of the 26 letters.

</details>

<details><summary>Hint 3</summary>

A `[26]int` array is **comparable** in Go, so it can be a map key directly: `map[[26]int][]string`. Fill counts per word, append the word under that key, then collect the map's values.

</details>

## Approach Guide

**Brute force — O(n² · k):** for each word, compare it against the representative of every existing group using an anagram check. Quadratic in the number of words.

**Sorted-key map — O(n · k log k):**
1. For each word, sort its bytes to get a key.
2. `groups[key] = append(groups[key], word)`.
3. Return all values of the map.

**Count-key map — O(n · k):**
1. For each word, build `var count [26]int` by incrementing `count[ch-'a']`.
2. Use the array itself as the map key.
3. Collect the groups.

(`k` is the maximum word length.)

## Target Complexity

- **Time:** O(n · k) with a count key (O(n · k log k) with a sorted key)
- **Space:** O(n · k)

## Go Concepts You'll Practise

- Arrays are value types and comparable — usable as map keys; slices are not
- Converting `string` ↔ `[]byte` and sorting with `slices.Sort`
- Map iteration order is randomised — never rely on it
- Pre-sizing the result with `make([][]string, 0, len(groups))`

## Interview Follow-ups

- What if the strings can contain any Unicode characters?
- Which key is faster in practice for short words vs long words, and why?
- How would you do this for billions of words that don't fit in memory?
- Return the groups sorted by size, largest first.

## Run the Tests

```bash
cd 16-group-anagrams
go test -v
```
