package rotatedsearch

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found in right half", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"missing value", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"single element missing", []int{1}, 0, -1},
		{"single element found", []int{1}, 1, 0},
		{"empty slice", []int{}, 5, -1},
		{"not rotated", []int{1, 2, 3, 4, 5, 6}, 5, 4},
		{"rotated by one, target at end", []int{6, 1, 2, 3, 4, 5}, 5, 5},
		{"target is pivot head", []int{6, 7, 8, 1, 2, 3}, 6, 0},
		{"target is minimum", []int{30, 40, 50, 10, 20}, 10, 3},
		{"two elements rotated", []int{3, 1}, 1, 1},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Search(tc.nums, tc.target); got != tc.want {
				t.Errorf("Search(%v, %d) = %d, want %d", tc.nums, tc.target, got, tc.want)
			}
		})
	}
}

func TestSearchEveryIndex(t *testing.T) {
	base := []int{-8, -3, 0, 2, 5, 9, 14, 21}
	for shift := range base {
		rotated := append(append([]int{}, base[shift:]...), base[:shift]...)
		for i, v := range rotated {
			if got := Search(rotated, v); got != i {
				t.Fatalf("Search(%v, %d) = %d, want %d (rotation by %d)", rotated, v, got, i, shift)
			}
		}
		if got := Search(rotated, 1); got != -1 {
			t.Fatalf("Search(%v, 1) = %d, want -1 (rotation by %d)", rotated, got, shift)
		}
	}
}
