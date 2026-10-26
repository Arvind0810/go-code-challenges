package romantoint

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{name: "single symbol", s: "I", want: 1},
		{name: "repeated symbol", s: "III", want: 3},
		{name: "subtractive IV", s: "IV", want: 4},
		{name: "subtractive IX", s: "IX", want: 9},
		{name: "additive mix", s: "LVIII", want: 58},
		{name: "subtractive XL and XC", s: "XCIX", want: 99},
		{name: "subtractive CD", s: "CDXLIV", want: 444},
		{name: "all subtractive pairs", s: "MCMXCIV", want: 1994},
		{name: "largest value", s: "MMMCMXCIX", want: 3999},
		{name: "longest numeral", s: "MMMDCCCLXXXVIII", want: 3888},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := RomanToInt(tc.s)
			if got != tc.want {
				t.Errorf("RomanToInt(%q) = %d, want %d", tc.s, got, tc.want)
			}
		})
	}
}
