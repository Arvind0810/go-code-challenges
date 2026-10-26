package rotatematrix

import (
	"reflect"
	"slices"
	"testing"
)

func cloneMatrix(matrix [][]int) [][]int {
	out := make([][]int, len(matrix))
	for i, row := range matrix {
		out[i] = slices.Clone(row)
	}
	return out
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   [][]int
	}{
		{
			name:   "3x3",
			matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want:   [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name:   "4x4",
			matrix: [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want:   [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{name: "1x1", matrix: [][]int{{1}}, want: [][]int{{1}}},
		{name: "empty", matrix: [][]int{}, want: [][]int{}},
		{name: "2x2", matrix: [][]int{{1, 2}, {3, 4}}, want: [][]int{{3, 1}, {4, 2}}},
		{name: "negatives", matrix: [][]int{{-1, -2}, {-3, -4}}, want: [][]int{{-3, -1}, {-4, -2}}},
		{
			name:   "5x5",
			matrix: [][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
			want:   [][]int{{21, 16, 11, 6, 1}, {22, 17, 12, 7, 2}, {23, 18, 13, 8, 3}, {24, 19, 14, 9, 4}, {25, 20, 15, 10, 5}},
		},
		{name: "duplicates", matrix: [][]int{{0, 0, 1}, {0, 0, 1}, {0, 0, 1}}, want: [][]int{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cloneMatrix(tc.matrix)
			Rotate(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Rotate(%v) = %v, want %v", tc.matrix, got, tc.want)
			}
		})
	}
}
