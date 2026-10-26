package reversestring

// Reverse returns s with its Unicode code points (runes) in reverse order.
// Multi-byte UTF-8 characters must remain intact.
func Reverse(s string) string {
	// TODO: implement
	rn := []rune(s)
	nrn := []rune{}
	for i := len(rn) - 1; i >= 0; i-- {
		nrn = append(nrn, rn[i])
	}
	return string(nrn)
}
