package prime_test

import (
	"fmt"
	"github.com/afshin/prime"
)

func ExampleIsPrime() {
	var candidate uint64 = 5
	if prime.IsPrime(candidate) {
		fmt.Println("Five is prime.")
	} else {
		fmt.Println("Five is not prime.")
	}
	// Output: Five is prime.
}
