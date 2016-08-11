// statictype2
package main

import (
	"fmt"
)

func main() {
	var x float64 = 20.0

	/* since y is not declared with an explicit type,
	you can use %T to determine its type */
	y := 42

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
}
