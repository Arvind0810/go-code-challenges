package wordbreak

import (
	"strings"
	"testing"
)

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name string
		s    string
		dict []string
		want bool
	}{
		{name: "leetcode", s: "leetcode", dict: []string{"leet", "code"}, want: true},
		{name: "reuse word", s: "applepenapple", dict: []string{"apple", "pen"}, want: true},
		{name: "cats and og", s: "catsandog", dict: []string{"cats", "dog", "sand", "and", "cat"}, want: false},
		{name: "single word", s: "go", dict: []string{"go"}, want: true},
		{name: "empty dict", s: "go", dict: nil, want: false},
		{name: "greedy trap", s: "aaaaaaab", dict: []string{"a", "aa", "aaa"}, want: false},
		{name: "needs shorter prefix", s: "goalspecial", dict: []string{"go", "goal", "goals", "special"}, want: true},
		{name: "prefix only", s: "carsx", dict: []string{"car", "cars"}, want: false},
		{name: "overlapping options", s: "pineapplepenapple", dict: []string{"apple", "pen", "applepen", "pine", "pineapple"}, want: true},
		{name: "dict word longer than s", s: "a", dict: []string{"aa"}, want: false},
		{name: "case sensitive", s: "GoLang", dict: []string{"go", "lang"}, want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := WordBreak(tc.s, tc.dict); got != tc.want {
				t.Errorf("WordBreak(%q, %q) = %v, want %v", tc.s, tc.dict, got, tc.want)
			}
		})
	}
}

func TestWordBreakExponentialTrap(t *testing.T) {
	s := strings.Repeat("a", 150) + "b"
	dict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	if WordBreak(s, dict) {
		t.Fatalf("WordBreak(\"aaa...ab\", a-words) = true, want false")
	}
	if !WordBreak(s[:150], dict) {
		t.Fatalf("WordBreak(\"aaa...a\", a-words) = false, want true")
	}
}
