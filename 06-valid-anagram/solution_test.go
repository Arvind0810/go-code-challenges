package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "both empty", s: "", t: "", want: true},
		{name: "single same", s: "a", t: "a", want: true},
		{name: "single different", s: "a", t: "b", want: false},
		{name: "classic", s: "anagram", t: "nagaram", want: true},
		{name: "different letters", s: "rat", t: "car", want: false},
		{name: "different lengths", s: "ab", t: "abc", want: false},
		{name: "same letters different counts", s: "aab", t: "abb", want: false},
		{name: "identical", s: "listen", t: "listen", want: true},
		{name: "all alphabet", s: "abcdefghijklmnopqrstuvwxyz", t: "zyxwvutsrqponmlkjihgfedcba", want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsAnagram(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("IsAnagram(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}

func TestIsAnagramUnicode(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{name: "both empty", s: "", t: "", want: true},
		{name: "ascii anagram", s: "listen", t: "silent", want: true},
		{name: "accented", s: "été", t: "éét", want: true},
		{name: "accent vs plain", s: "été", t: "ete", want: false},
		{name: "case sensitive", s: "Ab", t: "ab", want: false},
		{name: "cjk", s: "日本語", t: "語日本", want: true},
		{name: "emoji", s: "🍕🍔🍕", t: "🍔🍕🍕", want: true},
		{name: "emoji counts differ", s: "🍕🍔🍔", t: "🍔🍕🍕", want: false},
		{name: "spaces and punctuation count", s: "a b!", t: "!b a", want: true},
		{name: "same byte length different runes", s: "éa", t: "aè", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsAnagramUnicode(tc.s, tc.t)
			if got != tc.want {
				t.Errorf("IsAnagramUnicode(%q, %q) = %v, want %v", tc.s, tc.t, got, tc.want)
			}
		})
	}
}
