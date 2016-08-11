// typecast
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	// the result of square root is converted to float64
	var f float64 = math.Sqrt(float64(x*x + y*y))

	/*short hand form of casting the float value to uint
	same as var z uint = uint(f) */
	z := uint(f)

	fmt.Println(x, y, z)
}
