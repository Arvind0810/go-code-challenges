# 32 · Serialize and Deserialize Binary Tree

**Difficulty:** Hard · **Topics:** Binary Tree, DFS, BFS, String Parsing, Design · **Asked at:** Google, Meta, Amazon, Uber, LinkedIn

## Problem

Design an encoding for binary trees. Implement:

- `Serialize(root *TreeNode) string` — turn a tree into a string.
- `Deserialize(data string) *TreeNode` — rebuild **exactly** the same tree (same shape, same values) from that string.

You choose the format. The only requirement is that `Deserialize(Serialize(t))` is structurally identical to `t` for every tree, including the empty tree, trees with negative values, and trees with duplicate values.

## Examples

```text
        1
       / \
      2   3
         / \
        4   5

Input:  root = [1, 2, 3, null, null, 4, 5]
Output: Deserialize(Serialize(root)) is identical to root
One possible encoding (pre-order, "#" for nil):  "1,2,#,#,3,4,#,#,5,#,#,"
```

```text
Input:  root = nil
Output: Deserialize(Serialize(nil)) == nil
```

```text
      1            1
     /              \
    1                1

Input:  either tree above
Output: each round-trips to its own shape
Explanation: values alone are not enough — the encoding must capture where children are missing.
```

## Constraints

- The number of nodes is in the range `[0, 10^4]`.
- `-1000 <= Node.Val <= 1000`
- Values may repeat.
- Do not use package-level state — the two functions must be independent and stateless.

## Function Signature

```go
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Serialize(root *TreeNode) string

func Deserialize(data string) *TreeNode
```

## Hints

<details><summary>Hint 1</summary>

A pre-order or in-order listing of values alone is ambiguous. What extra information do you need so a single traversal is enough to rebuild the tree?

</details>

<details><summary>Hint 2</summary>

Emit a sentinel (e.g. `#`) for every `nil` child. With nil markers, a single **pre-order** traversal uniquely determines the tree. Separate tokens with a delimiter like `,` so multi-digit and negative numbers are unambiguous.

</details>

<details><summary>Hint 3</summary>

To deserialize, split into tokens and keep an index shared by a recursive closure: read a token; if it is `#` return `nil`, otherwise create a node, then build its left subtree, then its right subtree — the same order you wrote them.

</details>

## Approach Guide

**Naive idea — two traversals:** store both pre-order and in-order sequences and rebuild with the classic "construct from pre+in order" algorithm. This breaks as soon as values repeat, so it is not a general solution.

**Optimal A — pre-order DFS with nil markers:**

1. Serialize: recursive pre-order. For `nil`, write `#`; otherwise write the value. Join tokens with `,`.
2. Deserialize: split on `,`, keep a cursor `i`.
   - `build()`: token = tokens[i]; i++. If `#`, return nil.
   - Otherwise parse the int, make a node, `node.Left = build()`, `node.Right = build()`, return node.

**Optimal B — BFS level order (LeetCode style):**

1. Serialize: queue starting at root; for each popped node write its value (or `#` for nil) and enqueue both children when non-nil.
2. Deserialize: the first token is the root; then consume tokens in pairs, attaching left/right children to nodes taken from a queue in order.
3. Optionally trim trailing `#`s.

Both are linear. Pre-order is shorter to write; BFS output is easier to read and matches how LeetCode prints trees.

## Target Complexity

- **Time:** O(n) for both functions.
- **Space:** O(n) for the output string / token slice, plus O(h) recursion stack (DFS) or O(w) queue (BFS).

## Go Concepts You'll Practise

- `strings.Builder` for efficient string assembly (avoid `s += ...` in a loop — that is O(n²)).
- `strconv.Itoa` / `strconv.Atoi` and handling the returned error.
- `strings.Split` and the trailing empty token produced by a trailing delimiter.
- Recursive closures: `var build func() *TreeNode` declared before assignment so it can call itself, capturing a shared cursor.
- Slices as queues (`queue = queue[1:]`) for the BFS variant.

## Interview Follow-ups

1. How would you make the encoding more compact? (Binary encoding with `encoding/binary`, varints, or bitmaps for the nil structure.)
2. If the tree is a **BST**, can you drop the nil markers? (Yes — pre-order plus value bounds is enough.)
3. How would you generalise to an **N-ary tree**?
4. How would you validate untrusted input in `Deserialize` and return an error instead of panicking on malformed data?

## Run the Tests

```bash
cd 32-serialize-deserialize-tree
go test -v
```
