package numberofislands

import "testing"

func toGrid(rows []string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, row := range rows {
		grid[i] = []byte(row)
	}
	return grid
}

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid []string
		want int
	}{
		{name: "empty grid", grid: nil, want: 0},
		{name: "single water", grid: []string{"0"}, want: 0},
		{name: "single land", grid: []string{"1"}, want: 1},
		{
			name: "one big island",
			grid: []string{
				"11110",
				"11010",
				"11000",
				"00000",
			},
			want: 1,
		},
		{
			name: "three islands",
			grid: []string{
				"11000",
				"11000",
				"00100",
				"00011",
			},
			want: 3,
		},
		{
			name: "diagonals do not connect",
			grid: []string{
				"101",
				"010",
				"101",
			},
			want: 5,
		},
		{name: "single row", grid: []string{"1011011"}, want: 3},
		{name: "single column", grid: []string{"1", "1", "0", "1"}, want: 2},
		{
			name: "ring around lake",
			grid: []string{
				"11111",
				"10001",
				"10101",
				"10001",
				"11111",
			},
			want: 2,
		},
		{
			name: "snake shaped island",
			grid: []string{
				"11111",
				"00001",
				"11111",
				"10000",
				"11111",
			},
			want: 1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NumIslands(toGrid(tc.grid))
			if got != tc.want {
				t.Errorf("NumIslands(%q) = %d, want %d", tc.grid, got, tc.want)
			}
		})
	}
}
