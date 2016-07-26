// max
package main

import (
	"fmt"
)

func main() {
	// declare local variables
	var a int = 100
	var b int = 200
	var ret int

	// get the max of two numbers
	ret = max(a, b)

	fmt.Printf("Max value is : %d\n", ret)

}

func max(num1, num2 int) int {
	// declare local variable
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}

	return result
}
