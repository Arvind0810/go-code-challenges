# 04 · Word Frequency

**Difficulty:** Easy · **Topics:** Hash Map, Sorting, Strings · **Asked at:** Amazon, Google, Yelp, Bloomberg

## Problem

Given a block of text and an integer `n`, return the `n` most frequent words as a slice of `WordCount` values.

Rules:

- A **word** is a maximal run of letters or digits (`unicode.IsLetter` / `unicode.IsDigit`). Every other character — spaces, punctuation, apostrophes, hyphens — is a separator. So `"don't"` produces the words `"don"` and `"t"`.
- Counting is **case-insensitive**; report words in lowercase.
- Sort by **count descending**. Break ties **alphabetically ascending** by word.
- If `n` is greater than the number of distinct words, return all of them. If `n <= 0` or the text has no words, return an empty slice.

```go
type WordCount struct {
    Word  string
    Count int
}
```

## Examples

```text
Input:  text = "the cat and the hat", n = 2
Output: [{the 2} {and 1}]
Explanation: "the" appears twice. "and", "cat" and "hat" tie at 1; "and" wins alphabetically.
```

```text
Input:  text = "Go, go, GO! Gophers love Go.", n = 3
Output: [{go 4} {gophers 1} {love 1}]
Explanation: all spellings of "go" are counted together after lowercasing.
```

```text
Input:  text = "!!! ...", n = 5
Output: []
Explanation: there are no words.
```

## Constraints

- `0 <= len(text) <= 10^6` bytes
- `-10 <= n <= 10^4`
- `text` is valid UTF-8.

## Function Signature

```go
type WordCount struct {
	Word  string
	Count int
}

func TopWords(text string, n int) []WordCount
```

## Hints

<details><summary>Hint 1</summary>

Split the problem into three steps: tokenise, count, rank.

</details>

<details><summary>Hint 2</summary>

`strings.FieldsFunc` can split on "anything that is not a letter or digit". Count lowercased words in a `map[string]int`.

</details>

<details><summary>Hint 3</summary>

Copy the map into a `[]WordCount`, then sort with a comparator: `if a.Count != b.Count { return a.Count > b.Count }; return a.Word < b.Word`. Return `list[:min(n, len(list))]`.

</details>

## Approach Guide

**Straightforward approach:**

1. Tokenise: split `text` into words using a predicate that returns `true` for separator runes.
2. Count: lowercase each word and increment `counts[word]`.
3. Collect: move map entries into a slice (map iteration order is random — you *must* sort).
4. Rank: sort by count desc, word asc.
5. Truncate to `n` (guarding `n <= 0` and `n > len`).

**Optimisation for large vocabularies with small `n`:** instead of sorting all `k` distinct words (O(k log k)), keep a min-heap of size `n` with `container/heap` for O(k log n). Mention this, but implement the sort first.

## Target Complexity

- **Time:** O(L + k log k), where `L` is the text length and `k` the number of distinct words
- **Space:** O(k)

## Go Concepts You'll Practise

- `map[string]int` and the zero-value trick: `counts[w]++` works without initialising the key.
- Randomised map iteration order and why tests would be flaky without sorting.
- `sort.Slice` (or `slices.SortFunc` with `cmp.Compare`) with a multi-key comparator.
- `strings.FieldsFunc`, `strings.ToLower`, `unicode.IsLetter`, `unicode.IsDigit`.
- Struct types and slice-of-struct literals.
- The built-in `min` function (Go 1.21+).

## Interview Follow-ups

1. Implement the top-`n` selection with a heap. When is it better than a full sort?
2. The text is a 50 GB log file. How would you count words with bounded memory?
3. How would you parallelise counting across goroutines and merge the results safely?
4. How would you treat `"don't"` as a single word while still stripping quotes around `'quoted'` words?

## Run the Tests

```bash
cd 04-word-frequency
go test -v
```
