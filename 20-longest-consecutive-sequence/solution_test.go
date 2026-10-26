package longestconsecutive

import (
	"slices"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "classic example", nums: []int{100, 4, 200, 1, 3, 2}, want: 4},
		{name: "with duplicate zero", nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, want: 9},
		{name: "duplicates", nums: []int{1, 0, 1, 2}, want: 3},
		{name: "empty", nums: []int{}, want: 0},
		{name: "single element", nums: []int{42}, want: 1},
		{name: "all duplicates", nums: []int{7, 7, 7}, want: 1},
		{name: "negatives crossing zero", nums: []int{-2, -1, 0, 1, 5}, want: 4},
		{name: "no consecutive", nums: []int{10, 30, 20}, want: 1},
		{name: "two runs pick longer", nums: []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, want: 7},
		{name: "large values", nums: []int{1_000_000_000, -1_000_000_000, 999_999_999}, want: 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LongestConsecutive(slices.Clone(tc.nums)); got != tc.want {
				t.Errorf("LongestConsecutive(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
