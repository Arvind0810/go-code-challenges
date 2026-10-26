package coinchange

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{name: "classic example", coins: []int{1, 2, 5}, amount: 11, want: 3},
		{name: "impossible amount", coins: []int{2}, amount: 3, want: -1},
		{name: "zero amount", coins: []int{1}, amount: 0, want: 0},
		{name: "zero amount impossible coins", coins: []int{7}, amount: 0, want: 0},
		{name: "single coin exact", coins: []int{1}, amount: 1, want: 1},
		{name: "greedy fails", coins: []int{1, 3, 4}, amount: 6, want: 2},
		{name: "unsorted coins", coins: []int{5, 2, 1}, amount: 7, want: 2},
		{name: "coin larger than amount", coins: []int{10, 3}, amount: 9, want: 3},
		{name: "large amount", coins: []int{186, 419, 83, 408}, amount: 6249, want: 20},
		{name: "impossible with multiple coins", coins: []int{4, 6}, amount: 7, want: -1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := CoinChange(tc.coins, tc.amount)
			if got != tc.want {
				t.Errorf("CoinChange(%v, %d) = %d, want %d", tc.coins, tc.amount, got, tc.want)
			}
		})
	}
}
