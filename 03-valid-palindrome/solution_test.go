package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty", s: "", want: true},
		{name: "only spaces", s: "   ", want: true},
		{name: "only punctuation", s: ".,!?", want: true},
		{name: "single char", s: "a", want: true},
		{name: "classic sentence", s: "A man, a plan, a canal: Panama", want: true},
		{name: "not palindrome", s: "race a car", want: false},
		{name: "letter and digit mismatch", s: "0P", want: false},
		{name: "digits palindrome", s: "12321", want: true},
		{name: "question sentence", s: "Was it a car or a cat I saw?", want: true},
		{name: "unicode letters", s: "Été", want: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsPalindrome(tc.s)
			if got != tc.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
