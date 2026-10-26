package fizzbuzz

import (
	"strconv"
)

// FizzBuzz returns the FizzBuzz sequence for 1..n.
// Multiples of 3 become "Fizz", multiples of 5 become "Buzz", multiples of
// both become "FizzBuzz", and all other numbers are written in decimal.
// It returns an empty slice when n < 1.
func FizzBuzz(n int) []string {
	if n < 1 {
		return []string{}
	}
	// TODO: implement
	var output []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			output = append(output, "FizzBuzz")
		} else if i%3 == 0 {
			output = append(output, "Fizz")
		} else if i%5 == 0 {
			output = append(output, "Buzz")
		} else {
			output = append(output, strconv.Itoa(i))
		}
	}
	return output
}
