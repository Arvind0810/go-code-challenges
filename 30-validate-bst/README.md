# 30 · Validate Binary Search Tree

**Difficulty:** Medium · **Topics:** Binary Tree, BST, DFS, Recursion · **Asked at:** Amazon, Meta, Microsoft, Bloomberg, Oracle

## Problem

Given the `root` of a binary tree, determine whether it is a **valid binary search tree (BST)**.

A valid BST satisfies, for **every** node:

- all values in its **left** subtree are **strictly less** than the node's value,
- all values in its **right** subtree are **strictly greater** than the node's value,
- both subtrees are themselves valid BSTs.

Duplicates are therefore **not** allowed. An empty tree is a valid BST.

## Examples

```
Input:
      2
     / \
    1   3
Output: true
```

```
Input:
      5
     / \
    1   4
       / \
      3   6
Output: false
Explanation: The right child 4 is smaller than the root 5.
```

```
Input:
      5
     / \
    4   6
       / \
      3   7
Output: false
Explanation: Every parent/child pair looks fine locally, but 3 sits in the
             RIGHT subtree of 5 and 3 < 5. Local checks are not enough!
```

## Constraints

- The number of nodes is in the range `[0, 10^4]`.
- `math.MinInt64 <= Node.Val <= math.MaxInt64` — node values may be the extreme `int` values.

## Function Signature

```go
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func IsValidBST(root *TreeNode) bool
```

## Hints

<details><summary>Hint 1</summary>

Comparing each node only with its direct children is wrong — see example 3. Each node must respect a **range** inherited from all its ancestors.

</details>

<details><summary>Hint 2</summary>

Recurse with `(node, low, high)`. Going left tightens `high` to `node.Val`; going right tightens `low` to `node.Val`. But what initial `low`/`high` do you use when a node can legitimately equal `math.MinInt64`? Sentinels like `math.MinInt64` break — use "no bound" instead (e.g. `*int` that may be `nil`, or a `bool` flag).

</details>

<details><summary>Hint 3</summary>

Alternative: an **in-order** traversal of a BST visits values in strictly increasing order. Track the previously visited value (plus a "have we seen one yet?" flag, again because of `MinInt64`) and fail if the current value is `<=` the previous one.

</details>

## Approach Guide

**Brute force:** For each node, compute the max of its left subtree and the min of its right subtree and compare. O(n²) in the worst case (skewed tree).

**Optimal A (bounds recursion):**
1. `validate(node, low, high)` where `low`/`high` are optional bounds.
2. `nil` node → `true`.
3. If `low` exists and `node.Val <= low`, or `high` exists and `node.Val >= high` → `false`.
4. Return `validate(left, low, node.Val) && validate(right, node.Val, high)`.

**Optimal B (in-order):**
1. Traverse left → node → right (recursively or with an explicit stack).
2. Keep `prev` and `hasPrev`. If `hasPrev && node.Val <= prev`, return `false`.
3. Update `prev = node.Val`, `hasPrev = true`.

## Target Complexity

- Time: **O(n)**
- Space: **O(h)** for recursion / stack, where `h` is the tree height (O(n) worst case for a skewed tree).

## Go Concepts You'll Practise

- `math.MinInt64` / `math.MaxInt64` and why sentinel values are a bug magnet when the data range covers them.
- Using `*int` (nil = "unbounded") to express optional values — Go has no `Optional<T>`.
- Recursive closures: declare `var walk func(...) bool` first, then assign, so the closure can call itself.
- Short-circuit evaluation (`&&`) to stop early on the first violation.
- Iterative in-order traversal with a slice-backed stack.

## Interview Follow-ups

1. What changes if duplicates are allowed on one side (e.g. `left <= node < right`)?
2. Recover a BST where exactly two nodes were swapped by mistake (Recover BST).
3. Find the k-th smallest element in a BST.
4. Do it with **O(1)** extra space using Morris traversal.

## Run the Tests

```bash
cd 30-validate-bst
go test -v
```
