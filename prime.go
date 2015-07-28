// Package prime contains a helper function to check whether
// an integer is prime. In the future, it may contain other
// functions, but the API will remain backward-compatible.
package prime

import "math"

// floor is a private helper function to simplify some of
// the type conversions required in using math.Floor.
func floor(num uint64) uint64 {
	return uint64(math.Floor(float64(num)))
}

// IsPrime checks if candidate is a prime number.
func IsPrime(candidate uint64) bool {
	// 0 and 1 are not primes.
	if candidate < 2 {
		return false
	}
	// 2 is the only even number that is prime.
	if candidate == 2 {
		return true
	}
	// All other even numbers are not prime.
	if candidate%2 == 0 {
		return false
	}
	// LCV stands for loop control variable.
	var lcv uint64
	// For each divisor checked, close the gap from both sides.
	for lcv = 3; floor(candidate/lcv)+2 < lcv; lcv = lcv + 2 {
		if candidate%lcv == 0 {
			return false
		}
	}
	return true
}
