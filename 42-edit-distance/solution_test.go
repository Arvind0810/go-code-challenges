package editdistance

import (
	"strings"
	"testing"
)

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name         string
		word1, word2 string
		want         int
	}{
		{name: "both empty", word1: "", word2: "", want: 0},
		{name: "insert all", word1: "", word2: "abc", want: 3},
		{name: "delete all", word1: "abc", word2: "", want: 3},
		{name: "identical", word1: "golang", word2: "golang", want: 0},
		{name: "horse to ros", word1: "horse", word2: "ros", want: 3},
		{name: "intention to execution", word1: "intention", word2: "execution", want: 5},
		{name: "single substitution", word1: "cat", word2: "cut", want: 1},
		{name: "single insertion", word1: "cat", word2: "cart", want: 1},
		{name: "kitten to sitting", word1: "kitten", word2: "sitting", want: 3},
		{name: "completely different", word1: "abc", word2: "xyz", want: 3},
		{name: "reverse", word1: "abcd", word2: "dcba", want: 4},
		{name: "unicode runes", word1: "héllo", word2: "hello", want: 1},
		{name: "case sensitive", word1: "Go", word2: "go", want: 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := MinDistance(tc.word1, tc.word2); got != tc.want {
				t.Errorf("MinDistance(%q, %q) = %d, want %d", tc.word1, tc.word2, got, tc.want)
			}
		})
	}
}

func TestMinDistanceSymmetric(t *testing.T) {
	pairs := [][2]string{{"horse", "ros"}, {"sunday", "saturday"}, {"", "x"}, {"distance", "editing"}}
	for _, p := range pairs {
		a, b := MinDistance(p[0], p[1]), MinDistance(p[1], p[0])
		if a != b {
			t.Errorf("MinDistance is not symmetric: (%q,%q)=%d but (%q,%q)=%d", p[0], p[1], a, p[1], p[0], b)
		}
	}
}

func TestMinDistanceLarge(t *testing.T) {
	word1 := strings.Repeat("ab", 250)
	word2 := strings.Repeat("ba", 250)
	if got, want := MinDistance(word1, word2), 2; got != want {
		t.Errorf("MinDistance(\"abab...\", \"baba...\") = %d, want %d", got, want)
	}
}
