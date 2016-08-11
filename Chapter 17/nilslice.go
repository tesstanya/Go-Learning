// nilslice
package main

import "fmt"

func main() {
	var numbers []int

	printSlice(numbers)

	if numbers == nil {
		fmt.Println("Slice is nil")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
