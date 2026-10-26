package mergesortedarrays

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{name: "classic", nums1: []int{1, 2, 3, 0, 0, 0}, m: 3, nums2: []int{2, 5, 6}, n: 3, want: []int{1, 2, 2, 3, 5, 6}},
		{name: "nums2 empty", nums1: []int{1}, m: 1, nums2: []int{}, n: 0, want: []int{1}},
		{name: "nums1 empty", nums1: []int{0}, m: 0, nums2: []int{1}, n: 1, want: []int{1}},
		{name: "nums2 all smaller", nums1: []int{4, 5, 6, 0, 0, 0}, m: 3, nums2: []int{1, 2, 3}, n: 3, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "nums2 all larger", nums1: []int{1, 2, 3, 0, 0}, m: 3, nums2: []int{7, 8}, n: 2, want: []int{1, 2, 3, 7, 8}},
		{name: "interleaved", nums1: []int{1, 3, 5, 0, 0, 0}, m: 3, nums2: []int{2, 4, 6}, n: 3, want: []int{1, 2, 3, 4, 5, 6}},
		{name: "duplicates", nums1: []int{2, 2, 0, 0}, m: 2, nums2: []int{2, 2}, n: 2, want: []int{2, 2, 2, 2}},
		{name: "negatives", nums1: []int{-5, -1, 0, 0, 0}, m: 2, nums2: []int{-3, -2, 4}, n: 3, want: []int{-5, -3, -2, -1, 4}},
		{name: "unequal lengths", nums1: []int{0, 10, 0, 0, 0, 0}, m: 2, nums2: []int{-1, 1, 5, 20}, n: 4, want: []int{-1, 0, 1, 5, 10, 20}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := slices.Clone(tc.nums1)
			got := slices.Clone(tc.nums1)
			Merge(got, tc.m, tc.nums2, tc.n)
			if !slices.Equal(got, tc.want) {
				t.Errorf("Merge(%v, %d, %v, %d) -> nums1 = %v, want %v", input, tc.m, tc.nums2, tc.n, got, tc.want)
			}
		})
	}
}
