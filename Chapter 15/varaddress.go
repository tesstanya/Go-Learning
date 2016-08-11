// varaddress

/* This program prints the
address of a variable
aka its memory location */

package main

import (
	"fmt"
)

func main() {
	var a int = 10

	fmt.Printf("Address of variable 'a': %x\n", &a)
}
