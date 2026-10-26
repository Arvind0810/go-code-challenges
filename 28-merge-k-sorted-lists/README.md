# 28 · Merge k Sorted Lists

**Difficulty:** Hard · **Topics:** Linked List, Heap / Priority Queue, Divide and Conquer · **Asked at:** Amazon, Google, Meta, Microsoft, Uber, Databricks

## Problem

You are given a slice of `k` linked lists, each sorted in ascending order.

Merge all of them into **one** sorted linked list and return its head. Reuse the existing nodes — no need to allocate new ones.

## Examples

```
Input:
  1 -> 4 -> 5
  1 -> 3 -> 4
  2 -> 6
Output:
  1 -> 1 -> 2 -> 3 -> 4 -> 4 -> 5 -> 6
```

```
Input:  lists = []
Output: nil
```

```
Input:  lists = [nil]
Output: nil
Explanation: A slice containing one empty list still merges to an empty list.
```

## Constraints

- `0 <= k <= 10^4`
- `0 <= length of each list <= 500`
- `-10^4 <= Node.Val <= 10^4`
- Each list is sorted ascending; total nodes `N <= 10^4`.

## Function Signature

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeKLists(lists []*ListNode) *ListNode
```

## Hints

<details><summary>Hint 1</summary>

You already know how to merge **two** sorted lists. What happens if you merge list 0 with list 1, then the result with list 2, and so on? What's the cost?

</details>

<details><summary>Hint 2</summary>

At every step, the next node of the output is the smallest among the current **heads** of the `k` lists. Which data structure gives you the minimum of `k` items in O(log k)?

</details>

<details><summary>Hint 3</summary>

Implement `heap.Interface` on a `[]*ListNode` (`Len`, `Less`, `Swap`, `Push`, `Pop`). Push all non-nil heads, then repeatedly `heap.Pop` the smallest, append it to the result, and `heap.Push` its `Next` if non-nil.

</details>

## Approach Guide

**Brute force:** Collect every value into a slice, sort it, build a new list. O(N log N) time, O(N) space. Or merge lists one at a time: O(k · N).

**Optimal A (min-heap):**
1. Build a min-heap of the non-nil list heads, ordered by `Val`.
2. Use a dummy node + `tail` pointer for the output.
3. While the heap isn't empty: pop the smallest node, attach to `tail`, advance `tail`; if the node has a `Next`, push it.
4. Return `dummy.Next`.

**Optimal B (divide and conquer):** Pair lists up and merge each pair with the two-list merge; repeat until one remains. Also O(N log k), and O(1) extra space if done iteratively.

## Target Complexity

- Time: **O(N log k)** where `N` is the total number of nodes.
- Space: **O(k)** for the heap.

## Go Concepts You'll Practise

- `container/heap` and the `heap.Interface` contract — it embeds `sort.Interface` and adds `Push(x any)` / `Pop() any`.
- Why `Push`/`Pop` need **pointer receivers** (they change the slice length) while `Len`/`Less`/`Swap` can use value receivers.
- Calling `heap.Push(h, x)` / `heap.Pop(h)` — **not** `h.Push(x)` directly (a classic bug).
- Type assertions: `heap.Pop(h).(*ListNode)`.
- `heap.Init` to heapify in O(k).
- The dummy-head pattern for building output lists.

## Interview Follow-ups

1. Compare the heap and divide-and-conquer approaches — which has better cache behaviour? Which is easier to parallelise?
2. What if the lists are huge streams that don't fit in memory (external merge sort)?
3. Find the k-th smallest element across `k` sorted arrays.
4. Rewrite the heap with Go generics so it works for any ordered type.

## Run the Tests

```bash
cd 28-merge-k-sorted-lists
go test -v
```
