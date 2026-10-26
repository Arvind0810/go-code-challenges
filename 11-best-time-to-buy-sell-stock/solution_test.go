package stockprofit

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{name: "classic example", prices: []int{7, 1, 5, 3, 6, 4}, want: 5},
		{name: "strictly decreasing", prices: []int{7, 6, 4, 3, 1}, want: 0},
		{name: "empty", prices: []int{}, want: 0},
		{name: "nil", prices: nil, want: 0},
		{name: "single day", prices: []int{5}, want: 0},
		{name: "two days rising", prices: []int{1, 2}, want: 1},
		{name: "later minimum without higher sell", prices: []int{2, 4, 1}, want: 2},
		{name: "minimum after first peak", prices: []int{3, 8, 1, 9}, want: 8},
		{name: "all equal", prices: []int{4, 4, 4, 4}, want: 0},
		{name: "zero prices", prices: []int{0, 0, 10, 0}, want: 10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxProfit(tc.prices); got != tc.want {
				t.Errorf("MaxProfit(%v) = %d, want %d", tc.prices, got, tc.want)
			}
		})
	}
}
