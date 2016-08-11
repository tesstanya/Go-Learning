// zerovalues

/* This program prints the zero values of variables
that are declared without an initial value assigned */

package main

import (
	"fmt"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string

	// %v is the default value, %q is a single-quote character
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
