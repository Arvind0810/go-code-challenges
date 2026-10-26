package houserobber

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{name: "empty street", nums: nil, want: 0},
		{name: "single house", nums: []int{5}, want: 5},
		{name: "two houses pick larger", nums: []int{2, 7}, want: 7},
		{name: "classic four houses", nums: []int{1, 2, 3, 1}, want: 4},
		{name: "classic five houses", nums: []int{2, 7, 9, 3, 1}, want: 12},
		{name: "skip two in a row", nums: []int{2, 1, 1, 2}, want: 4},
		{name: "all zeros", nums: []int{0, 0, 0}, want: 0},
		{name: "greedy largest fails", nums: []int{5, 9, 5}, want: 10},
		{name: "increasing values", nums: []int{1, 2, 3, 4, 5, 6}, want: 12},
		{name: "longer street", nums: []int{183, 219, 57, 193, 94, 233, 202, 154, 65, 240, 97, 234, 100, 249, 186, 66, 90, 238, 168, 128}, want: 1954},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Rob(tc.nums)
			if got != tc.want {
				t.Errorf("Rob(%v) = %d, want %d", tc.nums, got, tc.want)
			}
		})
	}
}
