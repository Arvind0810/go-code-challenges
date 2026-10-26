package redundantconnection

import (
	"slices"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  []int
	}{
		{name: "triangle", edges: [][]int{{1, 2}, {1, 3}, {2, 3}}, want: []int{2, 3}},
		{name: "cycle of four with tail", edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, want: []int{1, 4}},
		{name: "two nodes doubled edge", edges: [][]int{{1, 2}, {2, 1}}, want: []int{2, 1}},
		{name: "redundant edge is first in cycle order", edges: [][]int{{2, 3}, {1, 2}, {1, 3}}, want: []int{1, 3}},
		{name: "cycle closes in the middle", edges: [][]int{{1, 2}, {2, 3}, {3, 1}, {3, 4}, {4, 5}}, want: []int{3, 1}},
		{name: "unordered labels", edges: [][]int{{3, 4}, {1, 2}, {2, 4}, {3, 5}, {2, 5}}, want: []int{2, 5}},
		{name: "long cycle through all nodes", edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 1}}, want: []int{6, 1}},
		{name: "star plus one extra edge", edges: [][]int{{1, 2}, {1, 3}, {1, 4}, {1, 5}, {4, 5}}, want: []int{4, 5}},
		{name: "two merged components", edges: [][]int{{1, 2}, {3, 4}, {5, 6}, {2, 3}, {4, 5}, {6, 1}}, want: []int{6, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FindRedundantConnection(tc.edges)
			if !slices.Equal(got, tc.want) {
				t.Errorf("FindRedundantConnection(%v) = %v, want %v", tc.edges, got, tc.want)
			}
		})
	}
}
