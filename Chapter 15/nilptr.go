// nilptr
package main

import (
	"fmt"
)

func main() {
	var ptr *int

	/* since the pointer is not assigned a value,
	it gets assigned the default int value of 0 */
	fmt.Printf("The value of ptr is : %x\n", ptr)
}
