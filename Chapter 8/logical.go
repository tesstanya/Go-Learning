// logical
package main

import (
	"fmt"
)

func main() {
	var a bool = true
	var b bool = false
	if a && b {
		fmt.Printf("Line 1 - Condition is true\n")
	}
	if a || b {
		fmt.Printf("Line 2 - Condition is true\n")
	}
	// changing the value of a and b
	a = false
	b = true
	if a && b {
		fmt.Printf("Line 3 - Condition is true\n")
	} else {
		fmt.Printf("Line 3 - Condition is not true\n")
	}
	if !(a && b) {
		fmt.Printf("Line 4 - Condition is true\n")
	}
}
