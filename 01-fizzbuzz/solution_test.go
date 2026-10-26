package fizzbuzz

import (
	"slices"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{name: "zero", n: 0, want: []string{}},
		{name: "negative", n: -5, want: []string{}},
		{name: "one", n: 1, want: []string{"1"}},
		{name: "first fizz", n: 3, want: []string{"1", "2", "Fizz"}},
		{name: "first buzz", n: 5, want: []string{"1", "2", "Fizz", "4", "Buzz"}},
		{
			name: "first fizzbuzz",
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
		{
			name: "past second fizzbuzz",
			n:    31,
			want: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz", "16", "17", "Fizz", "19", "Buzz",
				"Fizz", "22", "23", "Fizz", "Buzz", "26", "Fizz", "28", "29", "FizzBuzz",
				"31",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FizzBuzz(tc.n)
			if len(tc.want) == 0 && got == nil {
				t.Errorf("FizzBuzz(%d) = nil, want empty non-nil slice", tc.n)
				return
			}
			if !slices.Equal(got, tc.want) {
				t.Errorf("FizzBuzz(%d) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func TestFizzBuzzLarge(t *testing.T) {
	got := FizzBuzz(10000)
	if len(got) != 10000 {
		t.Fatalf("len(FizzBuzz(10000)) = %d, want 10000", len(got))
	}
	checks := map[int]string{99: "Fizz", 100: "Buzz", 9990: "FizzBuzz", 9998: "9998"}
	for i, want := range checks {
		if got[i-1] != want {
			t.Errorf("FizzBuzz(10000)[%d] = %q, want %q", i-1, got[i-1], want)
		}
	}
}
