// swapint

/* This program demonstrates how
the values of the actual variables do not
get changed when they are passed by value
as the parameters of a function */

package main

import (
	"fmt"
)

func main() {
	//local variable definition
	var a int = 100
	var b int = 200

	fmt.Printf("Before swap, value of a : %d\n", a)
	fmt.Printf("Before swap, value of b : %d\n", b)

	//calling swap function to switch values
	swap(a, b)

	fmt.Printf("After swap, value of a : %d\n", a)
	fmt.Printf("After swap, value of b : %d\n", b)
}

func swap(x, y int) int {
	var temp int

	temp = x
	x = y
	y = temp

	return temp
}
