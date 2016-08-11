// locglobvar

/* This program has the same variable name
for local and global variable, and demonstrates
that local variable takes priority */

package main

import (
	"fmt"
)

//global variable declaration
var g int = 20

func main() {
	//local variable declaration
	var g int = 10

	fmt.Printf("value of g = %d\n", g)
}
