package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{name: "empty", s: "", want: true},
		{name: "single pair", s: "()", want: true},
		{name: "sequential pairs", s: "()[]{}", want: true},
		{name: "nested", s: "{[()]}", want: true},
		{name: "deeply nested and sequential", s: "(([]){})[{}]", want: true},
		{name: "wrong type", s: "(]", want: false},
		{name: "interleaved", s: "([)]", want: false},
		{name: "single opener", s: "(", want: false},
		{name: "single closer", s: "]", want: false},
		{name: "unclosed openers", s: "(((", want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := IsValid(tc.s)
			if got != tc.want {
				t.Errorf("IsValid(%q) = %v, want %v", tc.s, got, tc.want)
			}
		})
	}
}
