// concatenate

/* The Join method is used for
concatenating two strings and
displaying as one string*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	//create and initialize a string array
	greetings := []string{"Hello", "world!"}

	//concatenate the string with a space in between
	fmt.Println(strings.Join(greetings, " "))
}
