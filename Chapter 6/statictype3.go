// statictype3
package main

import (
	"fmt"
)

func main() {
	var a, b, c = 3, 4, "foo" //values are assigned respectively

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("c is of type %T\n", c)
}
