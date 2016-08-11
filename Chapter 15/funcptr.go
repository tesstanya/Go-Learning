// funcptr
package main

import "fmt"

func main() {
	// local variable definition
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values.
	&a indicates pointer to a, address of variable a
	&b indicates pointer to b, address of variable b
	*/

	swap(&a, &b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int

	temp = *x // save the value at address x
	*x = *y   // save address of y into x
	*y = temp // save address of x in temp into y
}
