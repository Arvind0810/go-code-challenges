package validpalindrome

import (
	"regexp"
	"strings"
)

// IsPalindrome reports whether s reads the same forward and backward after
// lowercasing and discarding every rune that is not a letter or digit.
func IsPalindrome(s string) bool {
	// TODO: implement
	sn := SanitizeString(s)
	rn := []rune(sn)
	IsPalindrome := true

	for i := 0; i < (len(rn) / 2); i++ {
		if rn[i] != rn[len(rn)-(i+1)] {
			IsPalindrome = false
			break
		}
	}
	return IsPalindrome
}

func SanitizeString(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	result := re.ReplaceAllString(s, "")
	result = strings.ToLower(result)
	return result
}
