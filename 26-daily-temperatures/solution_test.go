package dailytemps

import (
	"slices"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		name  string
		temps []int
		want  []int
	}{
		{"leetcode example", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"strictly increasing", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"strictly decreasing", []int{90, 80, 70, 60}, []int{0, 0, 0, 0}},
		{"single day", []int{50}, []int{0}},
		{"empty", []int{}, []int{}},
		{"all equal", []int{70, 70, 70}, []int{0, 0, 0}},
		{"equal is not warmer", []int{70, 70, 71}, []int{2, 1, 0}},
		{"valley then peak", []int{55, 38, 53, 81, 61, 93, 97, 32, 43, 78}, []int{3, 1, 1, 2, 1, 1, 0, 1, 1, 0}},
		{"warm day at the very end", []int{89, 62, 70, 58, 47, 47, 46, 76, 100}, []int{8, 1, 5, 4, 3, 2, 1, 1, 0}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			input := slices.Clone(tc.temps)
			got := DailyTemperatures(input)
			if len(got) != len(tc.want) {
				t.Fatalf("DailyTemperatures(%v) returned %d values, want %d: got %v", tc.temps, len(got), len(tc.want), got)
			}
			if !slices.Equal(got, tc.want) {
				t.Errorf("DailyTemperatures(%v) = %v, want %v", tc.temps, got, tc.want)
			}
			if !slices.Equal(input, tc.temps) {
				t.Errorf("input was modified: now %v, originally %v", input, tc.temps)
			}
		})
	}
}
