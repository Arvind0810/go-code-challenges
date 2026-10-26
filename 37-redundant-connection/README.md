# 37 · Redundant Connection

**Difficulty:** Medium · **Topics:** Graph, Union-Find (Disjoint Set Union), Cycle Detection · **Asked at:** Google, Amazon, Microsoft, Uber

## Problem

A tree is an undirected, connected graph with no cycles. You start with a tree of `n` nodes labelled `1` to `n`, and then **one extra edge** is added between two different nodes. The result is given as `edges`, a list of `n` undirected edges `[a, b]`.

Return an edge that can be removed so the remaining graph is a tree of `n` nodes. If there are multiple answers, return the one that occurs **last** in `edges`.

## Examples

```text
    1
   / \
  2 ─ 3

Input:  edges = [[1,2],[1,3],[2,3]]
Output: [2,3]
Explanation: any edge of the triangle breaks the cycle; [2,3] appears last.
```

```text
  5 ─ 1 ─ 2
      │   │
      4 ─ 3

Input:  edges = [[1,2],[2,3],[3,4],[1,4],[1,5]]
Output: [1,4]
Explanation: the cycle is 1-2-3-4; of its edges, [1,4] appears last in the input.
```

```text
Input:  edges = [[1,2],[2,1]]
Output: [2,1]
```

## Constraints

- `n == len(edges)`
- `3 <= n <= 1000` (the tests also include a 2-node case)
- `edges[i] = [a, b]` with `1 <= a, b <= n` and `a != b`; the order inside a pair is not guaranteed.
- The graph is connected and contains exactly one cycle.

## Function Signature

```go
func FindRedundantConnection(edges [][]int) []int
```

## Hints

<details><summary>Hint 1</summary>

Add the edges one by one. Which edge is the first to connect two nodes that were **already connected**? Why is that edge the right answer given the "last in input" rule?

</details>

<details><summary>Hint 2</summary>

You need a structure that can answer "are `a` and `b` already in the same component?" and "merge the components of `a` and `b`" quickly. That is exactly Union-Find (Disjoint Set Union).

</details>

<details><summary>Hint 3</summary>

Keep `parent[i]` (initially `i`) and `rank[i]`. `find(x)` follows parents to the root and **compresses** the path by pointing every visited node straight at the root. `union` attaches the lower-rank root under the higher-rank root. If `find(a) == find(b)` before a union, return `[a, b]`.

</details>

## Approach Guide

**Brute force — try removing each edge:** for each edge from last to first, remove it and check with DFS/BFS whether the rest is connected and acyclic. O(n²).

**Better — DFS per edge:** before adding edge `[a, b]` to an adjacency list, DFS from `a` to see if `b` is already reachable. If it is, this edge closes the cycle. Still O(n²) worst case.

**Optimal — Union-Find:**

1. `parent := make([]int, n+1)` with `parent[i] = i`; `rank := make([]int, n+1)`.
2. `find(x)`: if `parent[x] != x`, set `parent[x] = find(parent[x])`; return `parent[x]`.
3. For each edge `[a, b]` in order:
   - `ra, rb := find(a), find(b)`.
   - If `ra == rb`, return this edge.
   - Otherwise union by rank: attach the shorter tree under the taller; if equal, pick one and increment its rank.

Because edges are processed in order, the first edge that closes a cycle is the last cycle edge in the input — exactly what the problem asks for.

## Target Complexity

- **Time:** O(n · α(n)) — effectively linear; α is the inverse Ackermann function.
- **Space:** O(n) for `parent` and `rank`.

## Go Concepts You'll Practise

- Modelling a reusable data structure: a `type DSU struct { parent, rank []int }` with `Find` / `Union` methods on a pointer receiver.
- Recursive closures (`var find func(int) int`) vs an iterative two-pass path compression.
- Tagless `switch` for the three union-by-rank cases.
- Returning a sub-slice of the input (`return edge`) — and when you'd copy it instead to avoid aliasing.
- `slices.Equal` (Go 1.21+) for comparing results in tests.

## Interview Follow-ups

1. **Redundant Connection II:** the graph is **directed** (a rooted tree plus one edge). What new cases appear? (A node with two parents, a cycle, or both.)
2. Why do you need both path compression and union by rank to get near-constant time? What is the bound with only one of them?
3. Use the same DSU to count the **number of connected components** in an undirected graph.
4. How would Union-Find solve **Number of Islands II** (cells added over time)?

## Run the Tests

```bash
cd 37-redundant-connection
go test -v
```
