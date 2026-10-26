package threesum

import (
	"reflect"
	"slices"
	"testing"
)

func normalize(triplets [][]int) [][]int {
	out := make([][]int, 0, len(triplets))
	for _, triplet := range triplets {
		sorted := slices.Clone(triplet)
		slices.Sort(sorted)
		out = append(out, sorted)
	}
	slices.SortFunc(out, slices.Compare)
	return out
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{name: "classic example", nums: []int{-1, 0, 1, 2, -1, -4}, want: [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{name: "no solution", nums: []int{0, 1, 1}, want: nil},
		{name: "all zeros", nums: []int{0, 0, 0, 0}, want: [][]int{{0, 0, 0}}},
		{name: "empty", nums: []int{}, want: nil},
		{name: "fewer than three", nums: []int{-1, 1}, want: nil},
		{name: "exactly three", nums: []int{-3, 1, 2}, want: [][]int{{-3, 1, 2}}},
		{name: "many duplicates", nums: []int{-2, 0, 0, 2, 2, -2}, want: [][]int{{-2, 0, 2}}},
		{
			name: "multiple triplets",
			nums: []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6},
			want: [][]int{
				{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2},
				{-2, -2, 4}, {-2, 0, 2},
			},
		},
		{name: "all positive", nums: []int{1, 2, 3, 4}, want: nil},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normalize(ThreeSum(slices.Clone(tc.nums)))
			want := normalize(tc.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("ThreeSum(%v) = %v, want %v", tc.nums, got, want)
			}
		})
	}
}
