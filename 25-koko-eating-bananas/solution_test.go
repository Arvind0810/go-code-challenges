package koko

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	cases := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"leetcode example 1", []int{3, 6, 7, 11}, 8, 4},
		{"leetcode example 2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"leetcode example 3", []int{30, 11, 23, 4, 20}, 6, 23},
		{"single pile, plenty of time", []int{10}, 10, 1},
		{"single pile, one hour", []int{1_000_000_000}, 1, 1_000_000_000},
		{"single pile, two hours", []int{1_000_000_000}, 2, 500_000_000},
		{"h equals number of piles", []int{1, 1, 1, 1}, 4, 1},
		{"lots of slack", []int{5, 8, 2}, 100, 1},
		{"huge piles with big h", []int{805306368, 805306368, 805306368}, 1_000_000_000, 3},
		{"just enough hours", []int{312884470}, 312884469, 2},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinEatingSpeed(tc.piles, tc.h); got != tc.want {
				t.Errorf("MinEatingSpeed(%v, %d) = %d, want %d", tc.piles, tc.h, got, tc.want)
			}
		})
	}
}
