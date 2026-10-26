# 44 · Combination Sum

**Difficulty:** Medium · **Topics:** Backtracking, Recursion, Sorting · **Asked at:** Amazon, Meta, Airbnb, Uber, Adobe

## Problem

Given a slice of **distinct positive** integers `candidates` and a target integer `target`, return all **unique combinations** of `candidates` where the chosen numbers sum to `target`.

- The same number may be chosen **any number of times**.
- Two combinations are the same if they contain the same numbers with the same frequencies — `[2, 2, 3]` and `[3, 2, 2]` are one combination.
- You may return the combinations in any order, and the numbers inside a combination in any order (the tests normalise both).
- If no combination exists, return an empty result.
- Do not modify the caller's `candidates` slice in a way that matters to them (sort a copy if you need sorted input).

## Examples

```text
Input:  candidates = [2, 3, 6, 7], target = 7
Output: [[2, 2, 3], [7]]
```

```text
Input:  candidates = [2, 3, 5], target = 8
Output: [[2, 2, 2, 2], [2, 3, 3], [3, 5]]
```

```text
Input:  candidates = [2], target = 1
Output: []
```

## Constraints

- `1 <= len(candidates) <= 30`
- `2 <= candidates[i] <= 40` (the tests also use `1`)
- All elements of `candidates` are distinct.
- `1 <= target <= 40`
- The number of unique combinations is less than 150 for the given inputs.

## Function Signature

```go
func CombinationSum(candidates []int, target int) [][]int
```

## Hints

<details><summary>Hint 1</summary>

Think of a decision tree: at each level, pick the next number to add to the current combination and subtract it from the remaining target. Stop when the remainder is 0 (record) or negative (abandon).

</details>

<details><summary>Hint 2</summary>

To avoid duplicates like `[2,3,2]` and `[3,2,2]`, only allow choosing candidates at index `>= start`, where `start` is the index of the last number chosen. Passing `i` (not `i+1`) into the recursion allows reuse of the same number.

</details>

<details><summary>Hint 3</summary>

Sort a copy of the candidates first — then once `cs[i] > remaining` you can `break` out of the loop instead of continuing. When you record a result, **copy** the path (`slices.Clone(path)`) — the backing array will be overwritten by later appends.

</details>

## Approach Guide

**Naive:** generate every multiset of candidates up to `target` (e.g. by building all sequences), then de-duplicate by sorting each and using a set of keys. Works, but explores many permutations of the same combination.

**Backtracking with a start index:**

```text
sort copy of candidates
dfs(start, remaining):
  if remaining == 0: record copy of path; return
  for i from start while cs[i] <= remaining:
    push cs[i]
    dfs(i, remaining - cs[i])   // i, not i+1: reuse allowed
    pop
dfs(0, target)
```

**Alternative (DP):** `ways[t]` = list of combinations summing to `t`, built like unbounded knapsack by iterating candidates in the outer loop. Uses more memory but no recursion.

## Target Complexity

- **Time:** O(N^(T/M + 1)) in the worst case, where `N` = number of candidates, `T` = target, `M` = smallest candidate (the depth of the tree is at most `T/M`).
- **Space:** O(T/M) recursion depth, plus the output.

## Go Concepts You'll Practise

- Recursive closures (`var dfs func(int, int)` then assign) capturing outer state.
- The **slice aliasing trap**: `append(res, path)` stores a view that later mutates — use `slices.Clone`.
- `slices.Sort`, `slices.Clone` and treating input as read-only.
- Push/pop on a slice with `append` and `path[:len(path)-1]`.

## Interview Follow-ups

1. Combination Sum II: each candidate used at most once, and the input may contain duplicates. How do you skip duplicate branches?
2. Only return the **count** of combinations — can you do it in O(N·T) with DP?
3. What if candidates can be negative? Why does the problem become ill-defined?
4. Return combinations lazily via a channel or an iterator (`iter.Seq[[]int]`).

## Run the Tests

```bash
cd 44-combination-sum
go test -v
```
