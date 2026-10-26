# 23 · Min Stack

**Difficulty:** Medium · **Topics:** Stack, Design · **Asked at:** Amazon, Bloomberg, Microsoft, Google, fintech companies

## Problem

Design a stack that supports `Push`, `Pop`, `Top`, and retrieving the **minimum** element — all in **O(1)** time.

Implement the `MinStack` type:

- `NewMinStack()` — creates an empty stack.
- `Push(val)` — pushes `val` onto the stack.
- `Pop()` — removes the top element.
- `Top()` — returns the top element.
- `GetMin()` — returns the smallest element currently in the stack.

`Pop`, `Top` and `GetMin` are only called on a non-empty stack.

## Examples

```
Operations: Push(-2), Push(0), Push(-3), GetMin(), Pop(), Top(), GetMin()
Output:     -,        -,       -,        -3,       -,     0,     -2

Stack (top on the right) as we go:
  [-2]           min -2
  [-2, 0]        min -2
  [-2, 0, -3]    min -3   <- GetMin() = -3
  [-2, 0]        min -2   <- Top() = 0, GetMin() = -2
```

```
Operations: Push(0), Push(1), Push(0), GetMin(), Pop(), GetMin()
Output:     -,       -,       -,       0,        -,     0
Explanation: Duplicate minimums must be handled — popping one 0 leaves another.
```

## Constraints

- `math.MinInt64 <= val <= math.MaxInt64`
- At most `3 * 10^4` calls in total.
- `Pop`, `Top`, `GetMin` are only called on non-empty stacks.

## Function Signature

```go
type MinStack struct{}

func NewMinStack() *MinStack
func (s *MinStack) Push(val int)
func (s *MinStack) Pop()
func (s *MinStack) Top() int
func (s *MinStack) GetMin() int
```

You will need to add fields to `MinStack`.

## Hints

<details><summary>Hint 1</summary>

A normal stack in Go is just a slice: push with `append`, pop with `s = s[:len(s)-1]`. The hard part is `GetMin` after a pop.

</details>

<details><summary>Hint 2</summary>

When you pop the minimum, what's the *new* minimum? It's whatever the minimum was at the moment that element was pushed. So remember it.

</details>

<details><summary>Hint 3</summary>

Store a pair `{val, minSoFar}` for every element, or keep a second stack of running minimums. `minSoFar` for a new element is `min(val, previous minSoFar)`.

</details>

## Approach Guide

**Brute force:** A plain slice; `GetMin` scans all elements. O(n) per `GetMin` — fails the requirement.

**Optimal:**
1. Each stack entry stores both the value and the minimum of the stack *including* that entry.
2. `Push(v)`: `m = v` if empty, else `min(v, top.min)`; append `{v, m}`.
3. `Pop`: reslice off the last entry.
4. `Top`: last entry's `val`. `GetMin`: last entry's `min`.

Space optimisation: a second "min stack" that only pushes when `v <= currentMin` (note the `<=` for duplicates).

## Target Complexity

- Time: **O(1)** for every operation (amortised for `Push` due to slice growth).
- Space: **O(n)**

## Go Concepts You'll Practise

- Designing a type with a constructor (`NewMinStack`) — Go has no constructors, just functions returning `*T`.
- Pointer receivers (`func (s *MinStack)`) — required because methods mutate the struct.
- Slices as stacks; amortised `append`; reslicing to pop.
- Small anonymous or named structs as slice elements (`[]struct{ val, min int }`).
- The built-in `min` function (Go 1.21+).

## Interview Follow-ups

1. Can you do it with a single stack of `int` and O(1) extra space? (Hint: store differences from the min.) What overflow risks does that create?
2. Add `GetMax` as well.
3. Make it safe for concurrent use — what's the simplest correct approach?
4. Design a **Min Queue** with O(1) amortised `GetMin`.

## Run the Tests

```bash
cd 23-min-stack
go test -v
```
