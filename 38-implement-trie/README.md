# 38 · Implement Trie (Prefix Tree)

**Difficulty:** Medium · **Topics:** Trie, Design, String · **Asked at:** Google, Amazon, Microsoft, Uber, Twitter

## Problem

A **trie** (prefix tree) stores strings so that lookups by prefix are fast — it powers autocomplete and spell checkers.

Implement the `Trie` type:

- `NewTrie() *Trie` — create an empty trie.
- `Insert(word string)` — add `word` to the trie.
- `Search(word string) bool` — `true` if `word` was inserted before (an exact match).
- `StartsWith(prefix string) bool` — `true` if any inserted word starts with `prefix`.

Each `Trie` must be independent — no package-level state.

## Examples

```text
Operations:
  t := NewTrie()
  t.Insert("apple")
  t.Search("apple")     → true
  t.Search("app")       → false   ("app" is only a prefix so far)
  t.StartsWith("app")   → true
  t.Insert("app")
  t.Search("app")       → true
```

```text
After inserting "car", "cat", "cart":

            (root)
              │
              c
              │
              a
            ┌─┴─┐
            r*  t*
            │
            t*            * = end of a word

Search("ca")       → false
StartsWith("ca")   → true
Search("cart")     → true
Search("cars")     → false
```

```text
Operations on an empty trie:
  Search("a")      → false
  StartsWith("a")  → false
```

## Constraints

- `1 <= len(word), len(prefix) <= 2000`
- `word` and `prefix` consist only of lowercase English letters `a`–`z`.
- At most `3 * 10^4` calls in total.

## Function Signature

```go
type Trie struct {
	// TODO: add fields (e.g. a root node with children and an end-of-word flag)
}

func NewTrie() *Trie

func (t *Trie) Insert(word string)

func (t *Trie) Search(word string) bool

func (t *Trie) StartsWith(prefix string) bool
```

## Hints

<details><summary>Hint 1</summary>

Each node represents a prefix. It needs links to child nodes (one per possible next letter) and a way to say "a complete word ends here".

</details>

<details><summary>Hint 2</summary>

Define an unexported `node` struct with `children [26]*node` and `end bool`. Map a letter to an index with `word[i] - 'a'`. The `Trie` holds a pointer to the root node, created in `NewTrie`.

</details>

<details><summary>Hint 3</summary>

`Search` and `StartsWith` both walk down the trie letter by letter and fail if a child is missing. Write one helper that returns the node reached by a string (or `nil`); then `Search` checks `node != nil && node.end`, and `StartsWith` checks `node != nil`.

</details>

## Approach Guide

**Brute force — slice or set of words:** store words in a `map[string]bool`. `Search` is O(L), but `StartsWith` must scan every word with `strings.HasPrefix` — O(N·L) per call.

**Better — set of all prefixes:** on insert, add every prefix of the word to a map. Lookups become O(L), but memory is O(N·L²) and you still need a separate map for full words.

**Optimal — trie of nodes:**

1. `node` holds `children` (array of 26 pointers, or a `map[byte]*node`) and an `end` flag.
2. `Insert`: start at the root; for each letter, create the child if missing and move to it; mark the last node `end = true`.
3. `walk(s)`: start at root; for each letter, move to the child; return `nil` if it is missing.
4. `Search(w)`: `n := walk(w); return n != nil && n.end`.
5. `StartsWith(p)`: `return walk(p) != nil`.

**Array vs map children:** `[26]*node` is faster and simpler for a fixed small alphabet; `map[rune]*node` saves memory for sparse tries and supports Unicode.

## Target Complexity

- **Time:** O(L) per operation, where L is the length of the word/prefix.
- **Space:** O(total characters inserted × alphabet size) in the worst case for array children.

## Go Concepts You'll Practise

- Designing a type with an unexported helper struct (`node`) and exported methods on a pointer receiver.
- Constructor functions (`NewTrie`) that initialise internal pointers so the zero value is never used accidentally.
- Fixed-size arrays of pointers (`[26]*node`) — the zero value is 26 `nil`s, no `make` needed.
- Byte arithmetic on strings: `word[i] - 'a'` indexes bytes, not runes.
- Trade-off between arrays and `map[byte]*node` for children.

## Interview Follow-ups

1. **Autocomplete:** add `Suggest(prefix string, limit int) []string` that returns up to `limit` words with that prefix in lexicographic order. (Walk to the prefix node, then DFS children in order.)
2. Add `Delete(word string)` that also prunes nodes no longer used by any word.
3. **Word Search II:** given a letter grid and a word list, find all words in the grid — why is a trie the key optimisation?
4. How would you make the trie safe for concurrent readers and writers? (`sync.RWMutex`, or immutable copy-on-write snapshots.)

## Run the Tests

```bash
cd 38-implement-trie
go test -v
```
