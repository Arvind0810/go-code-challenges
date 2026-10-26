package reverselist

import (
	"slices"
	"testing"
)

func fromSlice(vals []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}

func toSlice(head *ListNode) []int {
	var vals []int
	for n := head; n != nil; n = n.Next {
		vals = append(vals, n.Val)
		if len(vals) > 10_000 {
			return vals
		}
	}
	return vals
}

var reverseCases = []struct {
	name  string
	input []int
	want  []int
}{
	{"empty list", nil, nil},
	{"single node", []int{1}, []int{1}},
	{"two nodes", []int{1, 2}, []int{2, 1}},
	{"five nodes", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
	{"duplicates", []int{7, 7, 3, 3}, []int{3, 3, 7, 7}},
	{"negatives", []int{-1, 0, -5}, []int{-5, 0, -1}},
	{"already descending", []int{9, 6, 3}, []int{3, 6, 9}},
}

func TestReverseList(t *testing.T) {
	for _, tc := range reverseCases {
		t.Run(tc.name, func(t *testing.T) {
			got := toSlice(ReverseList(fromSlice(tc.input)))
			if !slices.Equal(got, tc.want) {
				t.Errorf("ReverseList(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestReverseListRecursive(t *testing.T) {
	for _, tc := range reverseCases {
		t.Run(tc.name, func(t *testing.T) {
			got := toSlice(ReverseListRecursive(fromSlice(tc.input)))
			if !slices.Equal(got, tc.want) {
				t.Errorf("ReverseListRecursive(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestReverseListReusesNodes(t *testing.T) {
	head := fromSlice([]int{1, 2, 3})
	tail := head.Next.Next
	got := ReverseList(head)
	if got != tail {
		t.Fatalf("expected the original tail node to become the new head (reverse in place, don't allocate)")
	}
	if head.Next != nil {
		t.Errorf("original head should now be the tail with Next == nil, got Next.Val = %d", head.Next.Val)
	}
}
