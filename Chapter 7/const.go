// const

// constants cannot be defined using the short hand method with :=

package main

import (
	"fmt"
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("Value of area : %d", area)
}
