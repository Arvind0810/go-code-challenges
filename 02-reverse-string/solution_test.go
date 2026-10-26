package reversestring

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{name: "empty", s: "", want: ""},
		{name: "single char", s: "a", want: "a"},
		{name: "two chars", s: "ab", want: "ba"},
		{name: "ascii word", s: "hello", want: "olleh"},
		{name: "palindrome", s: "racecar", want: "racecar"},
		{name: "with spaces and punctuation", s: "Go, rocks!", want: "!skcor ,oG"},
		{name: "accented latin", s: "héllo", want: "olléh"},
		{name: "cjk", s: "Hello, 世界", want: "界世 ,olleH"},
		{name: "emoji", s: "Hi 👋🌍", want: "🌍👋 iH"},
		{name: "only multibyte", s: "日本語", want: "語本日"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Reverse(tc.s)
			if got != tc.want {
				t.Errorf("Reverse(%q) = %q, want %q", tc.s, got, tc.want)
			}
			if !utf8.ValidString(got) {
				t.Errorf("Reverse(%q) returned invalid UTF-8: %q", tc.s, got)
			}
		})
	}
}
