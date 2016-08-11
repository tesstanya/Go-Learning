// initialvars
package main

import (
	"fmt"
)

var i, j int = 1, 2

func main() {
	// variable automatically takes the type of the value assigned to it
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
