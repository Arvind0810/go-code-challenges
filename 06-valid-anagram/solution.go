package validanagram

import "sort"

// IsAnagram reports whether t is an anagram of s.
// Both strings contain only lowercase English letters 'a'-'z'.
func IsAnagram(s, t string) bool {
	// TODO: implement
	if len(s) != len(t) {
		return false
	}
	s1 := []byte(s)
	t1 := []byte(t)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	for i := 0; i < len(s1); i++ {
		if s1[i] != t1[i] {
			return false
		}
	}
	return true
}

// IsAnagramUnicode reports whether t is an anagram of s, comparing Unicode
// code points case-sensitively with no normalisation.
func IsAnagramUnicode(s, t string) bool {
	// TODO: implement

	if len(s) != len(t) {
		return false
	}
	s1 := []rune(s)
	t1 := []rune(t)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})
	sort.Slice(t1, func(i, j int) bool {
		return t1[i] < t1[j]
	})
	for i := 0; i < len(s1); i++ {
		if s1[i] != t1[i] {
			return false
		}
	}
	return true
}
