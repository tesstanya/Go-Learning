// swap
package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Tanya", "Tess")
	fmt.Println(a, b) //Space is added automatically between each string
}
