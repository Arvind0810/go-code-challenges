package lis

import "testing"

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty", nums: []int{}, want: 0},
		{name: "nil", nums: nil, want: 0},
		{name: "single", nums: []int{7}, want: 1},
		{name: "classic", nums: []int{10, 9, 2, 5, 3, 7, 101, 18}, want: 4},
		{name: "with dip", nums: []int{0, 1, 0, 3, 2, 3}, want: 4},
		{name: "all equal", nums: []int{7, 7, 7, 7, 7, 7, 7}, want: 1},
		{name: "strictly increasing", nums: []int{1, 2, 3, 4, 5}, want: 5},
		{name: "strictly decreasing", nums: []int{5, 4, 3, 2, 1}, want: 1},
		{name: "negatives", nums: []int{-3, -1, -2, 0, -5, 4}, want: 4},
		{name: "duplicates do not count", nums: []int{1, 3, 3, 3, 5}, want: 3},
		{name: "late long run", nums: []int{9, 10, 1, 2, 3, 4}, want: 4},
		{name: "zigzag", nums: []int{1, 5, 2, 6, 3, 7, 4, 8}, want: 5},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LengthOfLIS(tc.nums); got != tc.want {
				t.Errorf("LengthOfLIS(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}

func TestLengthOfLISLargeInput(t *testing.T) {
	const n = 20000
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i
	}
	if got := LengthOfLIS(nums); got != n {
		t.Fatalf("LengthOfLIS(0..%d) = %d, want %d", n-1, got, n)
	}

	for i := range nums {
		nums[i] = (i%2)*n + i/2
	}
	if got, want := LengthOfLIS(nums), n/2+1; got != want {
		t.Fatalf("LengthOfLIS(interleaved) = %d, want %d", got, want)
	}
}
