// namedreturn

/* This program demonstrates a function that returns more
than one value with named return variables.
This should only be used in short functions as it will be
hard to read and follow in longer functions */

package main

import (
	"fmt"
)

/* there are two values returned here and they are the named
variables x and y */
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // called a naked return since it has no value after it
}

func main() {
	fmt.Println(split(17))
}
