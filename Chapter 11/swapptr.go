// swapptr

/* This program demonstrates how
the values of the actual variables
get changed when they are passed by reference
as the parameters of a function using pointers */

package main

import (
	"fmt"
)

func main() {
	// local variable definition
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	/* calling a function to swap the values.
	&a indicates the address of variable a and
	&b indicates the address of variable b */

	swap(&a, &b) //passing the address of the variables

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}

func swap(x *int, y *int) {
	var temp int
	temp = *x //saves the value at address x
	*x = *y   //saves the value at address y to x
	*y = temp //saves the value of temp into y
}
