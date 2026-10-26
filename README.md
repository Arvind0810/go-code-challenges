# Go Coding & DSA Challenges — Zero to Hero

50 interview-focused challenges in Go, ordered from language basics to concurrency and system-design-style problems.

## How each challenge works

Every folder contains:

| File | Purpose |
|---|---|
| `README.md` | Problem statement, examples, constraints, progressive hints, approach guide, complexity target, interview follow-ups |
| `solution.go` | Starter stub with the function signature — **write your code here** |
| `solution_test.go` | Table-driven tests — they fail until you solve the problem |

```bash
cd 01-fizzbuzz
go test -v            # run tests for one challenge
go test ./...         # from the root: run everything
go test -race ./...   # required habit for the concurrency phase (46–50)
```

## Suggested workflow

1. Read the problem. Write down inputs, outputs and edge cases **before** coding.
2. Solve it with brute force first, then optimize.
3. Open hints one at a time — only when stuck for 15+ minutes.
4. Say the time/space complexity out loud, like in a real interview.
5. Answer the follow-up questions at the bottom of each README.
6. Re-solve the challenge a week later without looking.

## Roadmap

### Phase 1 — Go Fundamentals (warm-up)
- [ ] 01 · [FizzBuzz](01-fizzbuzz) — loops, conditionals, slices
- [ ] 02 · [Reverse String](02-reverse-string) — runes vs bytes, Unicode
- [ ] 03 · [Valid Palindrome](03-valid-palindrome) — two pointers, `unicode` package
- [ ] 04 · [Word Frequency](04-word-frequency) — maps, sorting
- [ ] 05 · [Two Sum](05-two-sum) — hash map lookup
- [ ] 06 · [Valid Anagram](06-valid-anagram) — counting
- [ ] 07 · [Maximum Subarray](07-maximum-subarray) — Kadane's algorithm
- [ ] 08 · [Roman to Integer](08-roman-to-integer) — parsing
- [ ] 09 · [Valid Parentheses](09-valid-parentheses) — stack
- [ ] 10 · [Merge Sorted Arrays](10-merge-sorted-arrays) — in-place, back-to-front

### Phase 2 — Arrays, Strings, Two Pointers, Sliding Window
- [ ] 11 · [Best Time to Buy and Sell Stock](11-best-time-to-buy-sell-stock)
- [ ] 12 · [Product of Array Except Self](12-product-of-array-except-self)
- [ ] 13 · [Container With Most Water](13-container-with-most-water)
- [ ] 14 · [Three Sum](14-three-sum)
- [ ] 15 · [Longest Substring Without Repeating Characters](15-longest-substring-without-repeating)
- [ ] 16 · [Group Anagrams](16-group-anagrams)
- [ ] 17 · [Top K Frequent Elements](17-top-k-frequent-elements)
- [ ] 18 · [Merge Intervals](18-merge-intervals)
- [ ] 19 · [Rotate Matrix](19-rotate-matrix)
- [ ] 20 · [Longest Consecutive Sequence](20-longest-consecutive-sequence)

### Phase 3 — Linked Lists, Stacks, Binary Search, Heaps
- [ ] 21 · [Reverse Linked List](21-reverse-linked-list)
- [ ] 22 · [Linked List Cycle](22-linked-list-cycle)
- [ ] 23 · [Min Stack](23-min-stack)
- [ ] 24 · [Search in Rotated Sorted Array](24-search-rotated-sorted-array)
- [ ] 25 · [Koko Eating Bananas](25-koko-eating-bananas) — binary search on the answer
- [ ] 26 · [Daily Temperatures](26-daily-temperatures) — monotonic stack
- [ ] 27 · [LRU Cache](27-lru-cache) — map + doubly linked list
- [ ] 28 · [Merge K Sorted Lists](28-merge-k-sorted-lists) — `container/heap`

### Phase 4 — Trees & Graphs
- [ ] 29 · [Binary Tree Level Order Traversal](29-binary-tree-level-order) — BFS
- [ ] 30 · [Validate BST](30-validate-bst)
- [ ] 31 · [Lowest Common Ancestor](31-lowest-common-ancestor)
- [ ] 32 · [Serialize & Deserialize Binary Tree](32-serialize-deserialize-tree)
- [ ] 33 · [Number of Islands](33-number-of-islands) — grid DFS/BFS
- [ ] 34 · [Course Schedule](34-course-schedule) — topological sort
- [ ] 35 · [Rotting Oranges](35-rotting-oranges) — multi-source BFS
- [ ] 36 · [Network Delay Time](36-network-delay-time) — Dijkstra
- [ ] 37 · [Redundant Connection](37-redundant-connection) — Union-Find
- [ ] 38 · [Implement Trie](38-implement-trie)

### Phase 5 — Dynamic Programming & Backtracking
- [ ] 39 · [House Robber](39-house-robber) — 1D DP
- [ ] 40 · [Coin Change](40-coin-change) — unbounded knapsack
- [ ] 41 · [Longest Increasing Subsequence](41-longest-increasing-subsequence)
- [ ] 42 · [Edit Distance](42-edit-distance) — 2D DP
- [ ] 43 · [Word Break](43-word-break)
- [ ] 44 · [Combination Sum](44-combination-sum) — backtracking
- [ ] 45 · [N-Queens](45-n-queens)

### Phase 6 — Go Concurrency & System Design (hero level)
- [ ] 46 · [Worker Pool](46-worker-pool) — goroutines, channels, `sync.WaitGroup`
- [ ] 47 · [Rate Limiter](47-rate-limiter) — token bucket, injectable clock
- [ ] 48 · [Concurrent Pipeline](48-concurrent-pipeline) — fan-out/fan-in, `context` cancellation
- [ ] 49 · [Thread-Safe TTL Cache](49-ttl-cache) — `sync.RWMutex`, expiry
- [ ] 50 · [Concurrent Web Crawler](50-concurrent-web-crawler) — capstone

## Patterns cheat sheet

| Pattern | Recognise it when… | Challenges |
|---|---|---|
| Hash map | "find pair / count / seen before" | 05, 06, 16, 20 |
| Two pointers | sorted input, pairs, palindromes | 03, 10, 13, 14 |
| Sliding window | "longest/shortest substring/subarray with…" | 15 |
| Stack / monotonic stack | matching brackets, "next greater" | 09, 23, 26 |
| Binary search | sorted data, or "minimum X such that…" | 24, 25 |
| Heap | "top K", "merge K", "k-th smallest" | 17, 28, 36 |
| BFS | shortest path in unweighted graph, levels | 29, 35 |
| DFS / backtracking | enumerate all, explore regions | 33, 44, 45 |
| Topological sort | dependencies, ordering | 34 |
| Union-Find | connectivity, cycle detection | 37 |
| DP | "number of ways", "min/max", overlapping subproblems | 39–43 |
