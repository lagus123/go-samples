package main

import (
	"fmt"
)

var y = 30

func main() {
	// Println returns number of bytes written and error (which is ignored)
	n, _ := fmt.Println("Hello, playground", 42, true)
	// Prints number of bytes written
	fmt.Println(n)

	// Print with formatting
	fmt.Printf("%#x\n%b\n%+v \n", y, y, y)

	// Prints to string with formatting
	s := fmt.Sprintf("%#x\n%b\n%+v", y, y, y)

	fmt.Println(s)
}
