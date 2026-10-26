package mergeintervals

import (
	"reflect"
	"slices"
	"testing"
)

func cloneIntervals(intervals [][]int) [][]int {
	out := make([][]int, len(intervals))
	for i, interval := range intervals {
		out[i] = slices.Clone(interval)
	}
	return out
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{name: "classic example", intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, want: [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{name: "touching", intervals: [][]int{{1, 4}, {4, 5}}, want: [][]int{{1, 5}}},
		{name: "unsorted input", intervals: [][]int{{1, 4}, {0, 4}}, want: [][]int{{0, 4}}},
		{name: "empty", intervals: [][]int{}, want: nil},
		{name: "single interval", intervals: [][]int{{5, 7}}, want: [][]int{{5, 7}}},
		{name: "contained interval", intervals: [][]int{{1, 10}, {2, 3}, {4, 5}}, want: [][]int{{1, 10}}},
		{name: "no overlaps reversed", intervals: [][]int{{7, 8}, {4, 5}, {1, 2}}, want: [][]int{{1, 2}, {4, 5}, {7, 8}}},
		{name: "chain merge", intervals: [][]int{{2, 3}, {4, 6}, {1, 2}, {3, 4}}, want: [][]int{{1, 6}}},
		{name: "zero length intervals", intervals: [][]int{{0, 0}, {0, 0}, {1, 1}}, want: [][]int{{0, 0}, {1, 1}}},
		{name: "duplicates", intervals: [][]int{{2, 5}, {2, 5}, {6, 9}}, want: [][]int{{2, 5}, {6, 9}}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Merge(cloneIntervals(tc.intervals))
			if len(got) == 0 && len(tc.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Merge(%v) = %v, want %v", tc.intervals, got, tc.want)
			}
		})
	}
}
