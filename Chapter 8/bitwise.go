// bitwise
package main

import (
	"fmt"
)

func main() {

	var a uint = 60 // in bits 0011 1100
	var b uint = 13 // in bits 0000 1101
	var c uint = 0

	c = a & b /* becomes 0000 1100
	zeros for both stay as 0
	ones are copied only when both have 1
	or else any other ones turn into 0s */
	fmt.Printf("Line 1 - Value of c is %d\n\n", c) // 0000 1100 = 12

	c = a | b /* becomes 0011 1101
	zeros stay the same and ones get copied
	since one should be true for either */
	fmt.Printf("Line 2 - Value of c is %d\n\n", c)

	c = a ^ b /* becomes 0011 0001
	matching cases for 0^0 and 1^1 becomes 0
	0^1 or 1^0 becomes 1 */
	fmt.Printf("Line 3 - Value of c is %d\n\n", c) // 0011 0001 = 49)

	c = a << 2 /* becomes 1111 0000 when each digit is
	shifted to the left two times */
	fmt.Printf("Line 4 - Value of c is %d\n\n", c) // 1111 0000 = 240

	c = a >> 2 /* becomes 0000 1111 when each digit is
	shifted to the right two times */
	fmt.Printf("Line 5 - Value of c is %d\n\n", c) // 0000 1111 = 15
}
