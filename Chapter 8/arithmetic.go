// arithmetic
package main

import (
	"fmt"
)

func main() {
	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("Line 1 - value of c is %d\n", c)

	c = a - b
	fmt.Printf("Line 2 - value of c is %d\n", c)

	c = a * b
	fmt.Printf("Line 3 - value of c is %d\n", c)

	c = a / b
	fmt.Printf("Line 4 - value of c is %d\n", c)

	c = a % b
	fmt.Printf("Line 5 - value of c is %d\n", c)

	a++
	fmt.Printf("Line 6 - value of a is %d\n", a)

	a--
	fmt.Printf("Line 7 - value of a is %d\n", a)
}
