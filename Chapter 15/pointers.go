// pointers
package main

import (
	"fmt"
)

func main() {
	var a int = 20 // actual variable declaration
	var ip *int    // pointer variable declaration

	ip = &a /* stores the address of variable a
	into the pointer variable ip */

	fmt.Printf("Address of variable 'a': %x\n", &a)

	// address stored in pointer variable
	fmt.Printf("Address stored in variable 'ip': %x\n", ip) /* same
	as the one from variable 'a' */

	// access the value using the pointer
	fmt.Printf("Value of *ip variable: %d\n", *ip) /* returns the value
	stored at that memory address which is the same as variable a
	which is 20 */
}
