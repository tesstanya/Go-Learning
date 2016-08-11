// slice
package main

import "fmt"

func main() {
	var numbers = make([]int, 3, 5) //declare a slice of length 3 and cap 5

	printSlice(numbers)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
