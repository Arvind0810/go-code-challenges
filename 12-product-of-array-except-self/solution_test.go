package productexceptself

import (
	"slices"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{name: "classic example", nums: []int{1, 2, 3, 4}, want: []int{24, 12, 8, 6}},
		{name: "single zero", nums: []int{-1, 1, 0, -3, 3}, want: []int{0, 0, 9, 0, 0}},
		{name: "two zeros", nums: []int{0, 4, 0}, want: []int{0, 0, 0}},
		{name: "two elements", nums: []int{5, 7}, want: []int{7, 5}},
		{name: "single element", nums: []int{9}, want: []int{1}},
		{name: "empty", nums: []int{}, want: []int{}},
		{name: "negatives", nums: []int{-2, -3, 4}, want: []int{-12, -8, 6}},
		{name: "all ones", nums: []int{1, 1, 1, 1}, want: []int{1, 1, 1, 1}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := slices.Clone(tc.nums)
			got := ProductExceptSelf(input)
			if len(got) != len(tc.want) || !slices.Equal(got, tc.want) {
				t.Errorf("ProductExceptSelf(%v) = %v, want %v", tc.nums, got, tc.want)
			}
			if !slices.Equal(input, tc.nums) {
				t.Errorf("ProductExceptSelf mutated its input: got %v, original %v", input, tc.nums)
			}
		})
	}
}
