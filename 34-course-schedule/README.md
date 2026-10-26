# 34 · Course Schedule

**Difficulty:** Medium · **Topics:** Graph, Topological Sort, BFS (Kahn's Algorithm), DFS, Cycle Detection · **Asked at:** Amazon, Google, Meta, Microsoft, Airbnb

## Problem

There are `numCourses` courses labelled `0` to `numCourses-1`. Each pair `prerequisites[i] = [a, b]` means you must take course `b` **before** course `a`.

1. `CanFinish` — return `true` if it is possible to finish every course, otherwise `false`.
2. `FindOrder` (follow-up) — return **any** valid order in which to take all courses. If it is impossible, return an empty slice.

Multiple valid orders may exist; the tests accept any of them.

## Examples

```text
Input:  numCourses = 2, prerequisites = [[1, 0]]
Graph:  0 ──▶ 1
Output: CanFinish = true, FindOrder = [0, 1]
```

```text
Input:  numCourses = 2, prerequisites = [[1, 0], [0, 1]]
Graph:  0 ◀──▶ 1   (cycle)
Output: CanFinish = false, FindOrder = []
```

```text
Input:  numCourses = 4, prerequisites = [[1, 0], [2, 0], [3, 1], [3, 2]]
Graph:
          ┌──▶ 1 ──┐
        0 ┤        ├──▶ 3
          └──▶ 2 ──┘
Output: CanFinish = true, FindOrder = [0, 1, 2, 3] or [0, 2, 1, 3]
```

## Constraints

- `1 <= numCourses <= 2000`
- `0 <= len(prerequisites) <= 5000`
- `prerequisites[i] = [a, b]` with `0 <= a, b < numCourses`
- Pairs may repeat; a course may (pathologically) depend on itself.

## Function Signature

```go
func CanFinish(numCourses int, prerequisites [][]int) bool

func FindOrder(numCourses int, prerequisites [][]int) []int
```

## Hints

<details><summary>Hint 1</summary>

Model courses as nodes and each prerequisite `[a, b]` as a directed edge `b → a`. When is it impossible to finish? Think about what a cycle means for the "take before" relation.

</details>

<details><summary>Hint 2</summary>

A course with **in-degree 0** (no remaining prerequisites) can be taken right now. Taking it "removes" its outgoing edges, which may free up other courses.

</details>

<details><summary>Hint 3</summary>

Kahn's algorithm: queue all in-degree-0 nodes; pop one, append it to the order, decrement the in-degree of each neighbour, and enqueue any that reach 0. If the final order has fewer than `numCourses` nodes, a cycle blocked the rest.

</details>

## Approach Guide

**Brute force:** repeatedly scan all courses looking for one whose prerequisites are all taken; take it; stop when none can be taken. Correct, but O(V·E) or worse.

**Optimal A — Kahn's algorithm (BFS topological sort):**

1. Build an adjacency list `adj[b] = append(adj[b], a)` and an `indegree[a]++` for each pair.
2. Queue every node with `indegree == 0`.
3. While the queue is non-empty: pop `c`, append to `order`, for each neighbour `n`: `indegree[n]--`; if it hits 0, enqueue `n`.
4. `FindOrder`: return `order` if `len(order) == numCourses`, else an empty slice.
5. `CanFinish`: `len(order) == numCourses`.

**Optimal B — DFS with three colours:**

- `white` (unvisited), `grey` (on the current DFS path), `black` (fully done).
- Meeting a grey node during DFS means a back edge → cycle.
- Appending nodes to a list when they turn black gives a **reverse** topological order.

## Target Complexity

- **Time:** O(V + E) where V = `numCourses`, E = `len(prerequisites)`.
- **Space:** O(V + E) for the adjacency list, in-degree array, and queue.

## Go Concepts You'll Practise

- Building an adjacency list as `[][]int` with `make([][]int, n)` and `append` on nil inner slices.
- Slices as FIFO queues and the cost of `q = q[1:]` (the backing array is not freed until the slice is dropped).
- Returning an empty-but-non-nil slice (`[]int{}`) vs `nil` — and why callers should check `len` rather than `== nil`.
- Using a small `const` set or `iota` for DFS colours (`white`, `grey`, `black`).
- Reusing one function from another (`CanFinish` can delegate to `FindOrder`).

## Interview Follow-ups

1. Implement the DFS three-colour version. How do you produce the order from it?
2. **Course Schedule III / IV:** answer "is course `u` a prerequisite of `v`?" for many queries efficiently.
3. Return the **minimum number of semesters** if you can take any number of available courses per semester. (Kahn's level by level.)
4. If there is a cycle, how would you report the courses involved in it?

## Run the Tests

```bash
cd 34-course-schedule
go test -v
```
