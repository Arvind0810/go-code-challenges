package longestsubstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "classic example", s: "abcabcbb", want: 3},
		{name: "all same", s: "bbbbb", want: 1},
		{name: "substring not subsequence", s: "pwwkew", want: 3},
		{name: "empty", s: "", want: 0},
		{name: "single char", s: "x", want: 1},
		{name: "space only", s: " ", want: 1},
		{name: "all unique", s: "abcdef", want: 6},
		{name: "left must not move back", s: "abba", want: 2},
		{name: "repeat far back", s: "dvdf", want: 3},
		{name: "unicode runes", s: "héllo wörld", want: 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := LengthOfLongestSubstring(tc.s); got != tc.want {
				t.Errorf("LengthOfLongestSubstring(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
