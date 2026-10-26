# 21 · Reverse Linked List

**Difficulty:** Easy · **Topics:** Linked List, Pointers, Recursion · **Asked at:** Amazon, Microsoft, Apple, Bloomberg, most product companies

## Problem

Given the `head` of a singly linked list, reverse the list **in place** and return the new head.

Implement it twice:

1. `ReverseList` — iteratively.
2. `ReverseListRecursive` — recursively (the classic follow-up).

Do not allocate new nodes — rewire the existing `Next` pointers.

## Examples

```
Input:  1 -> 2 -> 3 -> 4 -> 5 -> nil
Output: 5 -> 4 -> 3 -> 2 -> 1 -> nil
```

```
Input:  1 -> 2 -> nil
Output: 2 -> 1 -> nil
```

```
Input:  nil
Output: nil
Explanation: An empty list reversed is still empty.
```

## Constraints

- The number of nodes is in the range `[0, 5000]`.
- `-5000 <= Node.Val <= 5000`

## Function Signature

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode
func ReverseListRecursive(head *ListNode) *ListNode
```

## Hints

<details><summary>Hint 1</summary>

Draw three boxes: `prev`, `curr`, `next`. At each step you only need to flip **one** arrow.

</details>

<details><summary>Hint 2</summary>

Before you overwrite `curr.Next`, save it — otherwise you lose the rest of the list. Then point `curr.Next` at `prev` and slide both pointers forward.

</details>

<details><summary>Hint 3</summary>

Recursive version: assume `ReverseListRecursive(head.Next)` already reversed everything after `head` and returned the new head. Now `head.Next` is the *tail* of that reversed list — make it point back to `head`, and cut `head.Next`.

</details>

## Approach Guide

**Brute force:** Walk the list, push values into a slice, then walk again writing values back in reverse order. O(n) time, O(n) extra space — and it mutates values, not links (interviewers usually reject this).

**Optimal (iterative):**
1. `prev = nil`, `curr = head`.
2. While `curr != nil`: save `next = curr.Next`, set `curr.Next = prev`, advance `prev = curr`, `curr = next`.
3. Return `prev`.

**Recursive:**
1. Base case: `head == nil || head.Next == nil` → return `head`.
2. `newHead = reverse(head.Next)`.
3. `head.Next.Next = head`; `head.Next = nil`.
4. Return `newHead`.

## Target Complexity

- Iterative: **O(n)** time, **O(1)** space.
- Recursive: **O(n)** time, **O(n)** space (call stack).

## Go Concepts You'll Practise

- Pointer-to-struct types (`*ListNode`) and the zero value `nil`.
- Multiple assignment for pointer juggling: `prev, curr, curr.Next = curr, curr.Next, prev` — and why the evaluation order makes it safe (or not!).
- Recursion depth: Go has growable goroutine stacks, so deep recursion rarely crashes, but it still costs O(n) memory.
- The "dummy head" pattern used in the test helper `fromSlice`.

## Interview Follow-ups

1. Reverse only the sublist between positions `left` and `right` (Reverse Linked List II).
2. Reverse the list in groups of `k` nodes.
3. Check whether a linked list is a palindrome in O(1) extra space (hint: reverse the second half).
4. Which version would you ship to production, and why?

## Run the Tests

```bash
cd 21-reverse-linked-list
go test -v
```
