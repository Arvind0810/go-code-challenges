package wordfrequency

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
	return nil
}
