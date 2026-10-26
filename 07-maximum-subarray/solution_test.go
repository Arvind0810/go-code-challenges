package maxsubarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "classic", nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "single positive", nums: []int{1}, want: 1},
		{name: "single negative", nums: []int{-7}, want: -7},
		{name: "all negative", nums: []int{-3, -1, -2}, want: -1},
		{name: "all positive", nums: []int{5, 4, -1, 7, 8}, want: 23},
		{name: "zeros and negatives", nums: []int{-2, 0, -1}, want: 0},
		{name: "best at the end", nums: []int{-1, -2, 3, 4}, want: 7},
		{name: "reset after big drop", nums: []int{8, -19, 5, -4, 20}, want: 21},
		{name: "duplicates", nums: []int{2, 2, -1, 2, 2}, want: 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxSubArray(tc.nums)
			if got != tc.want {
				t.Errorf("MaxSubArray(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
