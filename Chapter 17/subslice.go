// subslice
package main

import "fmt"

func main() {
	//create a slice
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)

	// Print the original slice
	fmt.Println("numbers ==", numbers)

	// print the sub slice starting from index 1(included) to index 4(excluded)
	fmt.Printf("numbers[1:4] ==", numbers[1:4])

	// missing lower bound implies 0
	fmt.Println("numbers[:3] ==", numbers[:3])

	// missing upper bound implies len(s)
	fmt.Println("numbers[4:] ==", numbers[4:]) // the last number is included

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	//print the subslice starting from index 0(included) to index 2(excluded)
	numbers2 := numbers[:2]
	printSlice(numbers2)

	//print the subslice starting from index 2(included) to index 5(excluded)
	numbers3 := numbers[2:5]
	printSlice(numbers3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
