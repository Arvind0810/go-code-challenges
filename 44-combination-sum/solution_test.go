package combinationsum

import (
	"fmt"
	"slices"
	"testing"
)

func normalise(combos [][]int) [][]int {
	out := make([][]int, len(combos))
	for i, c := range combos {
		cp := slices.Clone(c)
		slices.Sort(cp)
		out[i] = cp
	}
	slices.SortFunc(out, slices.Compare)
	return out
}

func equalCombos(a, b [][]int) bool {
	return slices.EqualFunc(a, b, func(x, y []int) bool { return slices.Equal(x, y) })
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{name: "classic", candidates: []int{2, 3, 6, 7}, target: 7, want: [][]int{{2, 2, 3}, {7}}},
		{name: "multiple", candidates: []int{2, 3, 5}, target: 8, want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{name: "impossible", candidates: []int{2}, target: 1, want: [][]int{}},
		{name: "single repeated", candidates: []int{1}, target: 3, want: [][]int{{1, 1, 1}}},
		{name: "unsorted input", candidates: []int{7, 3, 2, 6}, target: 7, want: [][]int{{2, 2, 3}, {7}}},
		{name: "candidate larger than target", candidates: []int{10, 20}, target: 5, want: [][]int{}},
		{name: "exact single", candidates: []int{4, 9}, target: 9, want: [][]int{{9}}},
		{
			name:       "many",
			candidates: []int{2, 3, 4, 5},
			target:     10,
			want: [][]int{
				{2, 2, 2, 2, 2}, {2, 2, 2, 4}, {2, 2, 3, 3}, {2, 3, 5},
				{2, 4, 4}, {3, 3, 4}, {5, 5},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := slices.Clone(tc.candidates)
			got := normalise(CombinationSum(input, tc.target))
			want := normalise(tc.want)
			if !equalCombos(got, want) {
				t.Errorf("CombinationSum(%v, %d) = %v, want %v (order-insensitive)", tc.candidates, tc.target, got, want)
			}
		})
	}
}

func TestCombinationSumNoDuplicatesAndValidSums(t *testing.T) {
	candidates := []int{2, 3, 5, 7, 11}
	target := 30
	got := CombinationSum(candidates, target)
	if len(got) == 0 {
		t.Fatalf("CombinationSum(%v, %d) returned no combinations", candidates, target)
	}
	seen := map[string]bool{}
	for _, c := range normalise(got) {
		sum := 0
		for _, x := range c {
			if !slices.Contains(candidates, x) {
				t.Errorf("combination %v uses %d which is not a candidate", c, x)
			}
			sum += x
		}
		if sum != target {
			t.Errorf("combination %v sums to %d, want %d", c, sum, target)
		}
		key := fmt.Sprint(c)
		if seen[key] {
			t.Errorf("duplicate combination %v (permutations of the same multiset count once)", c)
		}
		seen[key] = true
	}
}

func TestCombinationSumResultsAreIndependent(t *testing.T) {
	got := CombinationSum([]int{1, 2}, 4)
	if len(got) != 3 {
		t.Fatalf("CombinationSum([1 2], 4) returned %d combinations %v, want 3", len(got), got)
	}
	for _, c := range got {
		sum := 0
		for _, x := range c {
			sum += x
		}
		if sum != 4 {
			t.Fatalf("combination %v sums to %d; did you append a slice that was later mutated? Copy the path before storing it", c, sum)
		}
	}
}
