# 42 · Edit Distance

**Difficulty:** Hard · **Topics:** Dynamic Programming (2D), Strings · **Asked at:** Google, Amazon, Microsoft, Bloomberg, search/NLP teams

## Problem

Given two strings `word1` and `word2`, return the minimum number of operations required to convert `word1` into `word2` (the **Levenshtein distance**).

You may perform three operations, each costing 1:

- **Insert** a character,
- **Delete** a character,
- **Replace** a character.

Compare strings **rune by rune** (Unicode code points), not byte by byte — `"héllo"` → `"hello"` is a single replacement.

## Examples

```text
Input:  word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose  (delete 'r')
rose  -> ros   (delete 'e')
```

```text
Input:  word1 = "intention", word2 = "execution"
Output: 5
```

```text
Input:  word1 = "", word2 = "abc"
Output: 3
Explanation: insert three characters.
```

## Constraints

- `0 <= len(word1), len(word2) <= 500` (in runes)
- Strings may contain any valid UTF-8.
- Comparison is case-sensitive.

## Function Signature

```go
func MinDistance(word1, word2 string) int
```

## Hints

<details><summary>Hint 1</summary>

Think about the **last** characters. If they are equal, you can ignore both. If they differ, one of three operations must have produced the last character of `word2` (or removed the last character of `word1`).

</details>

<details><summary>Hint 2</summary>

Let `dp[i][j]` = edit distance between the first `i` runes of `word1` and the first `j` runes of `word2`. Base cases: `dp[i][0] = i`, `dp[0][j] = j`.

</details>

<details><summary>Hint 3</summary>

If `a[i-1] == b[j-1]`: `dp[i][j] = dp[i-1][j-1]`. Otherwise `dp[i][j] = 1 + min(dp[i-1][j-1] /*replace*/, dp[i-1][j] /*delete*/, dp[i][j-1] /*insert*/)`. Each row only depends on the previous row — two rows of length `m+1` are enough.

</details>

## Approach Guide

**Brute force (recursion):** `solve(i, j)` tries all three operations when characters differ. This branches 3 ways at each step — O(3^(n+m)).

**Top-down memoisation:** cache `solve(i, j)` in an `(n+1) x (m+1)` table. Each state is computed once → O(n·m).

**Bottom-up tabulation:**

```text
convert both strings to []rune
dp is (n+1) x (m+1)
dp[i][0] = i; dp[0][j] = j
for i in 1..n:
  for j in 1..m:
    if a[i-1] == b[j-1]: dp[i][j] = dp[i-1][j-1]
    else: dp[i][j] = 1 + min(diag, up, left)
return dp[n][m]
```

**Space optimisation:** keep only `prev` and `cur` rows and swap them after each `i`. Bonus: make the shorter string the column dimension so space is O(min(n, m)).

## Target Complexity

- **Time:** O(n · m)
- **Space:** O(min(n, m)) with rolling rows (O(n · m) for the full table is fine as a first answer)

## Go Concepts You'll Practise

- `[]rune(s)` conversion and why indexing a `string` gives bytes.
- Allocating 2D slices (`make([][]int, n+1)` + inner `make`) vs a flat slice.
- The variadic built-in `min(a, b, c)` (Go 1.21+).
- Swapping slices with `prev, cur = cur, prev` to reuse memory.

## Interview Follow-ups

1. Reconstruct the actual sequence of edits (backtrack through the table).
2. What if insert, delete and replace have different costs?
3. Only answer "is the distance at most `k`?" — can you restrict the DP to a diagonal band for O(n·k)?
4. How would you power a "did you mean?" feature over a dictionary of millions of words (BK-trees, tries with DP rows)?

## Run the Tests

```bash
cd 42-edit-distance
go test -v
```
