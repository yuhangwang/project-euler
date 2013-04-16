// Package euler provides convenience functions for solving Project
// Euler problems.
package euler

import "math"

// Factor
func Factor(n int) (factors map[int]int) {
	factors = make(map[int]int)
	for divisor := 2; divisor <= n; divisor++ {
		for n%divisor == 0 {
			n /= divisor
			factors[divisor] += 1
		}
		if n <= 1 {
			break
		}
	}
	return factors
}

// IsPalindrome returns true if the string is palindromic.
func IsPalindrome(s string) bool {
	runes := []rune(s)
	length := len(runes)
	for index, rune := range runes {
		if rune != runes[length-index-1] {
			return false
		}
	}
	return true
}

// IsPrime
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	if n < 9 {
		return true
	}
	if n%3 == 0 {
		return false
	}

	for x := 5; x <= int(math.Sqrt(float64(n))); x += 6 {
		if n%x == 0 || n%(x+2) == 0 {
			return false
		}
	}

	return true
}
