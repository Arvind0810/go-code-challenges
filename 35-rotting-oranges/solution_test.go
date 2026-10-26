package rottingoranges

import "testing"

func copyGrid(grid [][]int) [][]int {
	out := make([][]int, len(grid))
	for i, row := range grid {
		out[i] = append([]int(nil), row...)
	}
	return out
}

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{name: "classic four minutes", grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, want: 4},
		{name: "unreachable fresh orange", grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, want: -1},
		{name: "no fresh oranges", grid: [][]int{{0, 2}}, want: 0},
		{name: "all empty", grid: [][]int{{0, 0}, {0, 0}}, want: 0},
		{name: "fresh with no rotten", grid: [][]int{{1}}, want: -1},
		{name: "multiple sources meet in middle", grid: [][]int{{2, 1, 1, 1, 2}}, want: 2},
		{name: "multiple sources in corners", grid: [][]int{{2, 1, 1}, {1, 1, 1}, {1, 1, 2}}, want: 2},
		{name: "diagonal does not spread", grid: [][]int{{2, 0}, {0, 1}}, want: -1},
		{name: "single column chain", grid: [][]int{{2}, {1}, {1}, {1}}, want: 3},
		{name: "already all rotten", grid: [][]int{{2, 2}, {2, 2}}, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := OrangesRotting(copyGrid(tc.grid))
			if got != tc.want {
				t.Errorf("OrangesRotting(%v) = %d, want %d", tc.grid, got, tc.want)
			}
		})
	}
}
