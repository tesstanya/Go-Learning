// nestedforloop

/* This program loops from 2 to 100 and
finds all the prime numbers between them */

package main

import (
	"fmt"
)

func main() {
	/* local variable definition */
	var i, j int

	// outer loop that iterates through the values
	for i = 2; i < 100; i++ {

		// inner loop that looks for factors
		/* the number from the outer loop
		is divided by every number less than it
		in the inner loop to find a factor if present */
		for j = 2; j <= (i / j); j++ {
			if i%j == 0 { //if two numbers are evenly divided...
				break //if factor found, not prime
			}
		}

		if j > (i / j) {
			fmt.Printf("%d is prime\n", i)
		}
	}
}
