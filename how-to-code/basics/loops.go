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

	//loops

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after for loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop.

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
