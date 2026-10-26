package topkfrequent

import (
	"slices"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{name: "classic example", nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{name: "single element", nums: []int{1}, k: 1, want: []int{1}},
		{name: "negatives", nums: []int{4, -1, -1, 4, 4, 7}, k: 1, want: []int{4}},
		{name: "k equals distinct count", nums: []int{5, 3, 5, 9}, k: 3, want: []int{3, 5, 9}},
		{name: "all same", nums: []int{2, 2, 2, 2}, k: 1, want: []int{2}},
		{name: "unsorted input", nums: []int{3, 0, 1, 0, 3, 3, 1, 3, 0}, k: 2, want: []int{0, 3}},
		{name: "negative most frequent", nums: []int{-5, -5, -5, 6, 6, 7, 8}, k: 2, want: []int{-5, 6}},
		{name: "zero counted", nums: []int{0, 0, 1}, k: 1, want: []int{0}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := slices.Clone(TopKFrequent(slices.Clone(tc.nums), tc.k))
			slices.Sort(got)
			want := slices.Clone(tc.want)
			slices.Sort(want)
			if len(got) != len(want) || !slices.Equal(got, want) {
				t.Errorf("TopKFrequent(%v, %d) = %v, want %v", tc.nums, tc.k, got, want)
			}
		})
	}
}
