package wordfrequency

import (
	"sort"
	"strings"
	"unicode"
)

// WordCount pairs a lowercased word with the number of times it occurs.
type WordCount struct {
	Word  string
	Count int
}

// TopWords returns the n most frequent words in text, counted
// case-insensitively, where a word is a maximal run of letters or digits.
// Results are ordered by count descending, then word ascending. It returns an
// empty slice when n <= 0 or text contains no words.
func TopWords(text string, n int) []WordCount {
	// TODO: implement

	// return if n is less then or equel to 0
	if n <= 0 {
		return []WordCount{}
	}

	// Create a map for count store
	counts := make(map[string]int)
	// Filter the the string for not letter and not digit values
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	// convert the words into lower case and add word counts through a loop
	for _, word := range words {
		word = strings.ToLower(word)
		counts[word]++
	}

	// Create a new slice of word count
	result := make([]WordCount, 0, len(counts))
	// Append the word counts
	for word, count := range counts {
		result = append(result, WordCount{
			Word:  word,
			Count: count,
		})
	}
	// sort the slices
	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}
		return result[i].Word < result[j].Word
	})

	if n > len(result) {
		n = len(result)
	}

	return result[:n]
}
