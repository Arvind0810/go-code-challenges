package linkedlistcycle

import (
	"fmt"
	"testing"
)

// buildList links vals into a list and, when pos >= 0, connects the tail to
// the node at index pos. It returns the head and the expected cycle-start node.
func buildList(vals []int, pos int) (*ListNode, *ListNode) {
	nodes := make([]*ListNode, len(vals))
	for i, v := range vals {
		nodes[i] = &ListNode{Val: v}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	if len(nodes) == 0 {
		return nil, nil
	}
	if pos < 0 {
		return nodes[0], nil
	}
	nodes[len(nodes)-1].Next = nodes[pos]
	return nodes[0], nodes[pos]
}

var cycleCases = []struct {
	name string
	vals []int
	pos  int
}{
	{"empty list", nil, -1},
	{"single node no cycle", []int{1}, -1},
	{"single node self loop", []int{1}, 0},
	{"two nodes cycle to head", []int{1, 2}, 0},
	{"two nodes no cycle", []int{1, 2}, -1},
	{"classic cycle mid list", []int{3, 2, 0, -4}, 1},
	{"tail points to itself", []int{1, 2, 3, 4, 5}, 4},
	{"long list no cycle", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, -1},
	{"long tail short loop", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 7},
	{"duplicate values cycle", []int{5, 5, 5, 5}, 2},
}

func TestHasCycle(t *testing.T) {
	for _, tc := range cycleCases {
		t.Run(tc.name, func(t *testing.T) {
			head, _ := buildList(tc.vals, tc.pos)
			want := tc.pos >= 0
			if got := HasCycle(head); got != want {
				t.Errorf("HasCycle(vals=%v, pos=%d) = %v, want %v", tc.vals, tc.pos, got, want)
			}
		})
	}
}

func TestDetectCycleStart(t *testing.T) {
	for _, tc := range cycleCases {
		t.Run(tc.name, func(t *testing.T) {
			head, want := buildList(tc.vals, tc.pos)
			got := DetectCycleStart(head)
			if got != want {
				t.Errorf("DetectCycleStart(vals=%v, pos=%d) returned wrong node (got %s, want %s)",
					tc.vals, tc.pos, describe(got), describe(want))
			}
		})
	}
}

func describe(n *ListNode) string {
	if n == nil {
		return "nil"
	}
	return fmt.Sprintf("node{Val: %d}", n.Val)
}
