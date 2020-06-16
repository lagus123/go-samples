package main

import (
	"fmt"
)

// Custom types
type hotdog int

var x hotdog

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
