# 22 · Linked List Cycle

**Difficulty:** Easy / Medium · **Topics:** Linked List, Two Pointers, Floyd's Cycle Detection · **Asked at:** Amazon, Microsoft, Goldman Sachs, Adobe, Uber

## Problem

Given the `head` of a singly linked list, determine whether the list contains a **cycle** — i.e. some node can be reached again by continuously following `Next`.

Implement two functions:

1. `HasCycle(head)` — return `true` if there is a cycle, otherwise `false`.
2. `DetectCycleStart(head)` — return the **node** where the cycle begins, or `nil` if there is no cycle. (Return the node itself, not its value — values may repeat.)

Do not modify the list.

## Examples

```
Input:  3 -> 2 -> 0 -> -4
             ^          |
             +----------+
HasCycle:          true
DetectCycleStart:  the node with value 2
Explanation: The tail (-4) links back to the node at index 1.
```

```
Input:  1 -> 2
        ^    |
        +----+
HasCycle:          true
DetectCycleStart:  the node with value 1
```

```
Input:  1 -> nil
HasCycle:          false
DetectCycleStart:  nil
```

## Constraints

- The number of nodes is in the range `[0, 10^4]`.
- `-10^5 <= Node.Val <= 10^5`
- Target: **O(1)** extra memory.

## Function Signature

```go
type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool
func DetectCycleStart(head *ListNode) *ListNode
```

## Hints

<details><summary>Hint 1</summary>

The easy way: remember every node you've visited in a `map[*ListNode]bool`. Pointers are comparable in Go, so they make fine map keys. Now do it without the map.

</details>

<details><summary>Hint 2</summary>

Run two pointers: `slow` moves 1 step, `fast` moves 2 steps. If there's a cycle, `fast` eventually laps `slow` and they land on the same node. If `fast` hits `nil`, there is no cycle.

</details>

<details><summary>Hint 3</summary>

After they meet, reset one pointer to `head` and move **both** one step at a time. The node where they meet again is the cycle start. (Why? Let `a` = distance from head to cycle start, `b` = start to meeting point, `c` = cycle length. Then `2(a+b) = a+b+kc`, so `a = kc - b`.)

</details>

## Approach Guide

**Brute force:** Walk the list storing each node pointer in a set. The first node already in the set is the cycle start; reaching `nil` means no cycle. O(n) time, O(n) space.

**Optimal (Floyd's tortoise and hare):**
1. `slow, fast = head, head`.
2. Loop while `fast != nil && fast.Next != nil`: advance `slow` by 1, `fast` by 2. If `slow == fast`, a cycle exists.
3. For the start: set `p = head`; advance `p` and `slow` one step at a time until `p == slow`. Return `p`.

## Target Complexity

- Time: **O(n)**
- Space: **O(1)**

## Go Concepts You'll Practise

- Pointer identity comparison (`slow == fast` compares addresses, not values).
- Pointers as map keys (`map[*ListNode]struct{}`) for the brute-force version.
- Guarding nil dereferences: `fast != nil && fast.Next != nil` relies on `&&` short-circuiting.
- Returning a pointer vs. a value — why returning `*ListNode` is the right API when values can repeat.

## Interview Follow-ups

1. Prove why the second phase of Floyd's algorithm finds the cycle start.
2. How would you compute the **length** of the cycle?
3. Find the duplicate number in an array of `n+1` integers in `[1, n]` without modifying it (LeetCode 287 — same idea!).
4. Does it still work if `fast` moves 3 steps at a time?

## Run the Tests

```bash
cd 22-linked-list-cycle
go test -v
```
