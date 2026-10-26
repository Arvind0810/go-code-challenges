# 31 · Lowest Common Ancestor of a Binary Tree

**Difficulty:** Medium · **Topics:** Binary Tree, DFS, Recursion · **Asked at:** Meta, Amazon, Microsoft, LinkedIn

## Problem

Given the `root` of a binary tree (**not** a binary search tree) and two node pointers `p` and `q` that both exist in the tree, return their **lowest common ancestor (LCA)**.

The LCA is the deepest node that has both `p` and `q` as descendants, where a node is allowed to be a descendant of **itself**.

Compare by pointer identity (`node == p`), not by value.

## Examples

```text
            3
          /   \
         5     1
        / \   / \
       6   2 0   8
          / \
         7   4

Input:  p = 5, q = 1
Output: 3
Explanation: 5 is in the left subtree of 3, 1 is in the right subtree.
```

```text
Input:  same tree, p = 5, q = 4
Output: 5
Explanation: 5 is an ancestor of 4, and a node counts as its own descendant.
```

```text
       1
      /
     2

Input:  p = 1, q = 2
Output: 1
```

## Constraints

- The number of nodes is in the range `[2, 10^5]`.
- `-10^9 <= Node.Val <= 10^9`
- All node values are unique.
- `p` and `q` both exist in the tree (`p` may equal `q`).

## Function Signature

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode
```

## Hints

<details><summary>Hint 1</summary>

Think about what a recursive call on a subtree could *report back* to its parent. What would be useful for the parent to know?

</details>

<details><summary>Hint 2</summary>

Have the recursion return `p` or `q` if it finds either one in the subtree, or `nil` if it finds neither. If the current node *is* `p` or `q`, you can return it immediately without searching further.

</details>

<details><summary>Hint 3</summary>

At a node, if the left call returned non-nil **and** the right call returned non-nil, then `p` and `q` are split across this node — it is the LCA. Otherwise, bubble up whichever side was non-nil.

</details>

## Approach Guide

**Brute force — root-to-node paths:**

1. DFS to record the path (slice of node pointers) from `root` to `p`.
2. Do the same for `q`.
3. Walk both paths from the start; the last position where they agree is the LCA.

This is O(n) time but uses O(h) extra slices and two traversals.

**Alternative — parent pointers:** BFS/DFS once to fill a `map[*TreeNode]*TreeNode` of parents. Walk up from `p`, putting every ancestor in a set; then walk up from `q` until you hit a node in the set.

**Optimal — single post-order recursion:**

1. Base case: if the node is `nil`, `p`, or `q`, return the node.
2. Recurse left and right.
3. If both results are non-nil, return the current node.
4. Otherwise return the non-nil result (or `nil`).

Because `p` and `q` are guaranteed to exist, whatever reaches the top is the answer.

## Target Complexity

- **Time:** O(n) — each node is visited at most once.
- **Space:** O(h) recursion stack, where `h` is the tree height (O(n) for a skewed tree).

## Go Concepts You'll Practise

- Recursive functions over pointer-based structs; `nil` as a natural base case.
- Pointer identity comparison (`root == p`) vs comparing `Val`.
- Using `map[*TreeNode]*TreeNode` — pointers are valid, comparable map keys.
- Understanding Go's lack of tail-call optimisation and goroutine stack growth for deep recursion.

## Interview Follow-ups

1. How does the solution change if the tree is a **BST**? (Hint: you can decide left/right by comparing values, O(h) time, O(1) space iteratively.)
2. What if `p` or `q` might **not** exist in the tree? How do you avoid returning a wrong answer?
3. What if every node had a `Parent` pointer and you were *not* given the root? (Think "intersection of two linked lists".)
4. How would you answer many LCA queries on the same static tree efficiently? (Binary lifting / Euler tour + RMQ.)

## Run the Tests

```bash
cd 31-lowest-common-ancestor
go test -v
```
