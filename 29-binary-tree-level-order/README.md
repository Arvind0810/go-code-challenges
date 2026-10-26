# 29 · Binary Tree Level Order Traversal

**Difficulty:** Medium · **Topics:** Binary Tree, BFS, Queue · **Asked at:** Amazon, Meta, Microsoft, LinkedIn, Bloomberg

## Problem

Given the `root` of a binary tree, return the **level order traversal** of its node values: each level as its own slice, levels from top to bottom, values within a level from left to right.

An empty tree returns an empty result.

## Examples

```
Input:
        3
       / \
      9   20
         /  \
        15   7

Output: [[3], [9, 20], [15, 7]]
```

```
Input:
        1
       / \
      2   3
       \  /
       4 5

Output: [[1], [2, 3], [4, 5]]
```

```
Input:  nil
Output: []
```

## Constraints

- The number of nodes is in the range `[0, 2000]`.
- `-1000 <= Node.Val <= 1000`

## Function Signature

```go
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func LevelOrder(root *TreeNode) [][]int
```

## Hints

<details><summary>Hint 1</summary>

Breadth-first search visits nodes level by level. BFS uses a **queue** — in Go, a slice where you read from the front and append to the back.

</details>

<details><summary>Hint 2</summary>

The trick is knowing where one level ends. At the start of each round, the queue contains *exactly* the nodes of the current level — record `len(queue)` before you start adding children.

</details>

<details><summary>Hint 3</summary>

A clean Go pattern: keep `current []*TreeNode`; build `next []*TreeNode` from their children while collecting values; then `current = next`. No index bookkeeping needed.

</details>

## Approach Guide

**Alternative (DFS):** Recurse with a `depth` parameter; if `depth == len(result)`, append a new empty level; append the node's value to `result[depth]`; recurse left then right. Also O(n) — worth knowing as a second solution.

**Optimal (BFS):**
1. If `root == nil`, return an empty result.
2. `queue = [root]`.
3. While the queue is non-empty:
   - `size = len(queue)`; create `level` with capacity `size`.
   - Process `size` nodes: append the value; enqueue non-nil children.
   - Append `level` to the result.
4. Return the result.

## Target Complexity

- Time: **O(n)**
- Space: **O(w)** for the queue, where `w` is the maximum width of the tree (up to `n/2`), plus the output.

## Go Concepts You'll Practise

- Slices as queues: `queue[0]`, `queue = queue[1:]` — and the memory caveat (the backing array isn't freed until the slice is dropped).
- Slices of slices (`[][]int`) and building them with `append`.
- Pre-allocating with `make([]int, 0, size)`.
- Recursive struct types with pointer children; nil checks before enqueueing.
- Reading the test helper `buildTree` — converting LeetCode's `[3,9,20,null,null,15,7]` notation into Go, using `[]*int` so `nil` can mean "no node".

## Interview Follow-ups

1. Return the levels bottom-up (Level Order II).
2. Zigzag level order — alternate left→right and right→left.
3. Binary Tree Right Side View — return the last value of each level.
4. Average / maximum of each level; or the minimum depth of the tree using BFS early exit.

## Run the Tests

```bash
cd 29-binary-tree-level-order
go test -v
```
