package nqueens

import (
	"fmt"
	"slices"
	"strings"
	"testing"
)

func normaliseBoards(boards [][]string) []string {
	out := make([]string, len(boards))
	for i, b := range boards {
		out[i] = strings.Join(b, "/")
	}
	slices.Sort(out)
	return out
}

func validateBoard(t *testing.T, n int, board []string) {
	t.Helper()
	if len(board) != n {
		t.Fatalf("board has %d rows, want %d: %q", len(board), n, board)
	}
	queenCol := make([]int, n)
	for r, row := range board {
		if len(row) != n {
			t.Fatalf("row %d has length %d, want %d: %q", r, len(row), n, row)
		}
		if strings.Count(row, "Q") != 1 || strings.Count(row, ".") != n-1 {
			t.Fatalf("row %d must contain exactly one 'Q' and %d '.': %q", r, n-1, row)
		}
		queenCol[r] = strings.IndexByte(row, 'Q')
	}
	for r1 := range n {
		for r2 := r1 + 1; r2 < n; r2++ {
			c1, c2 := queenCol[r1], queenCol[r2]
			if c1 == c2 || r2-r1 == c2-c1 || r2-r1 == c1-c2 {
				t.Fatalf("queens at (%d,%d) and (%d,%d) attack each other in %q", r1, c1, r2, c2, board)
			}
		}
	}
}

func TestSolveNQueensExact(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{name: "n=1", n: 1, want: [][]string{{"Q"}}},
		{name: "n=2 has none", n: 2, want: [][]string{}},
		{name: "n=3 has none", n: 3, want: [][]string{}},
		{
			name: "n=4",
			n:    4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := normaliseBoards(SolveNQueens(tc.n))
			want := normaliseBoards(tc.want)
			if !slices.Equal(got, want) {
				t.Errorf("SolveNQueens(%d) = %q, want %q (order-insensitive)", tc.n, got, want)
			}
		})
	}
}

func TestSolveNQueensValidAndDistinct(t *testing.T) {
	counts := map[int]int{5: 10, 6: 4, 7: 40, 8: 92}
	for n, want := range counts {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			boards := SolveNQueens(n)
			if len(boards) != want {
				t.Fatalf("SolveNQueens(%d) returned %d boards, want %d", n, len(boards), want)
			}
			for _, b := range boards {
				validateBoard(t, n, b)
			}
			keys := normaliseBoards(boards)
			if len(slices.Compact(keys)) != want {
				t.Fatalf("SolveNQueens(%d) returned duplicate boards", n)
			}
		})
	}
}

func TestTotalNQueens(t *testing.T) {
	want := []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352, 724, 2680}
	for n, w := range want {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			if got := TotalNQueens(n); got != w {
				t.Errorf("TotalNQueens(%d) = %d, want %d", n, got, w)
			}
		})
	}
}

func TestTotalMatchesSolve(t *testing.T) {
	for n := 1; n <= 8; n++ {
		if a, b := TotalNQueens(n), len(SolveNQueens(n)); a != b {
			t.Errorf("n=%d: TotalNQueens = %d but len(SolveNQueens) = %d", n, a, b)
		}
	}
}
