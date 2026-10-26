package mergeklists

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
		if len(vals) > 100_000 {
			return vals
		}
	}
	return vals
}

func TestMergeKLists(t *testing.T) {
	cases := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"leetcode example", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"no lists", [][]int{}, nil},
		{"one empty list", [][]int{{}}, nil},
		{"several empty lists", [][]int{{}, {}, {}}, nil},
		{"single list", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"mix of empty and non-empty", [][]int{{}, {2}, {}, {1, 3}}, []int{1, 2, 3}},
		{"negatives and duplicates", [][]int{{-10, -5, 0}, {-5, -5}, {0, 7}}, []int{-10, -5, -5, -5, 0, 0, 7}},
		{"disjoint ranges", [][]int{{7, 8, 9}, {1, 2, 3}, {4, 5, 6}}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"uneven lengths", [][]int{{1}, {0, 2, 4, 6, 8, 10}, {5}}, []int{0, 1, 2, 4, 5, 6, 8, 10}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			lists := make([]*ListNode, len(tc.lists))
			for i, l := range tc.lists {
				lists[i] = fromSlice(l)
			}
			got := toSlice(MergeKLists(lists))
			if !slices.Equal(got, tc.want) {
				t.Errorf("MergeKLists(%v) = %v, want %v", tc.lists, got, tc.want)
			}
		})
	}
}

func TestMergeKListsManyLists(t *testing.T) {
	const k, n = 500, 20
	lists := make([]*ListNode, k)
	var want []int
	for i := range k {
		vals := make([]int, n)
		for j := range n {
			vals[j] = j*k + i
		}
		lists[i] = fromSlice(vals)
		want = append(want, vals...)
	}
	slices.Sort(want)
	got := toSlice(MergeKLists(lists))
	if !slices.Equal(got, want) {
		t.Fatalf("MergeKLists with k=%d lists of %d nodes: got %d values, want %d sorted values", k, n, len(got), len(want))
	}
}
