package twosum

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{name: "basic", nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{name: "cannot reuse element", nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{name: "duplicate values", nums: []int{3, 3}, target: 6, want: []int{0, 1}},
		{name: "negatives", nums: []int{-1, -2, -3, -4, -5}, target: -8, want: []int{2, 4}},
		{name: "mixed signs to zero", nums: []int{-3, 4, 3, 90}, target: 0, want: []int{0, 2}},
		{name: "zeros", nums: []int{0, 4, 3, 0}, target: 0, want: []int{0, 3}},
		{name: "pair at the end", nums: []int{1, 2, 3, 4, 5, 6}, target: 11, want: []int{4, 5}},
		{name: "no solution", nums: []int{1, 2, 3}, target: 100, want: nil},
		{name: "single element", nums: []int{5}, target: 10, want: nil},
		{name: "empty", nums: []int{}, target: 0, want: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TwoSum(tc.nums, tc.target)
			if tc.want == nil {
				if len(got) != 0 {
					t.Errorf("TwoSum(%v, %d) = %v, want nil", tc.nums, tc.target, got)
				}
				return
			}
			normalised := slices.Clone(got)
			slices.Sort(normalised)
			if !slices.Equal(normalised, tc.want) {
				t.Errorf("TwoSum(%v, %d) = %v, want %v", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}
