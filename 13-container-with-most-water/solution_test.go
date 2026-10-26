package containerwater

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{name: "classic example", height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, want: 49},
		{name: "two lines", height: []int{1, 1}, want: 1},
		{name: "outer lines win", height: []int{4, 3, 2, 1, 4}, want: 16},
		{name: "empty", height: []int{}, want: 0},
		{name: "single line", height: []int{5}, want: 0},
		{name: "inner pair wins", height: []int{1, 2, 4, 3}, want: 4},
		{name: "increasing", height: []int{1, 2, 3, 4, 5}, want: 6},
		{name: "contains zeros", height: []int{0, 0, 7, 0, 0}, want: 0},
		{name: "tall middle pair", height: []int{2, 3, 10, 5, 7, 8, 9}, want: 36},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MaxArea(tc.height); got != tc.want {
				t.Errorf("MaxArea(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}
