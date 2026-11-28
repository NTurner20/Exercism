// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

func IsDivisibleByFour(number int) bool {
	return number%4 == 0
}

func IsDivisibleBy100(number int) bool {
	return number%100 == 0
}

func IsDivisibleBy400(number int) bool {
	return number%400 == 0
}

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	// if divisible by 4
	if IsDivisibleByFour(year) {
		// if divisble by 100 and not 400 return false
		if IsDivisibleBy100(year) && !IsDivisibleBy400(year) {
			return false
		}
		return true
	} else {
		return false
	}
}
