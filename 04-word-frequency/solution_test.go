package wordfrequency

import (
	"reflect"
	"testing"
)

func TestTopWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		n    int
		want []WordCount
	}{
		{name: "empty text", text: "", n: 3, want: []WordCount{}},
		{name: "only punctuation", text: "!!! ...", n: 5, want: []WordCount{}},
		{name: "n zero", text: "a b c", n: 0, want: []WordCount{}},
		{name: "n negative", text: "a b c", n: -1, want: []WordCount{}},
		{name: "single word", text: "hello", n: 1, want: []WordCount{{"hello", 1}}},
		{
			name: "tie broken alphabetically",
			text: "the cat and the hat",
			n:    2,
			want: []WordCount{{"the", 2}, {"and", 1}},
		},
		{
			name: "case insensitive with punctuation",
			text: "Go, go, GO! Gophers love Go.",
			n:    3,
			want: []WordCount{{"go", 4}, {"gophers", 1}, {"love", 1}},
		},
		{
			name: "n larger than distinct words",
			text: "b a b",
			n:    10,
			want: []WordCount{{"b", 2}, {"a", 1}},
		},
		{
			name: "apostrophes hyphens and digits split words",
			text: "Don't stop: well-known 42 is 42",
			n:    10,
			want: []WordCount{{"42", 2}, {"don", 1}, {"is", 1}, {"known", 1}, {"stop", 1}, {"t", 1}, {"well", 1}},
		},
		{
			name: "unicode words",
			text: "Café café CAFÉ naïve",
			n:    2,
			want: []WordCount{{"café", 3}, {"naïve", 1}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopWords(tc.text, tc.n)
			if len(tc.want) == 0 {
				if got == nil || len(got) != 0 {
					t.Errorf("TopWords(%q, %d) = %#v, want empty non-nil slice", tc.text, tc.n, got)
				}
				return
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("TopWords(%q, %d) = %v, want %v", tc.text, tc.n, got, tc.want)
			}
		})
	}
}
