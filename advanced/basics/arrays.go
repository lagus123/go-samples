package main

import "fmt"

func main() {
	// Zero based index
	// Is better use slices, arrays as used as building blocks for that
	var x [5]int
	fmt.Println(x)

	x[3] = 10
	fmt.Println(x)
	fmt.Println(len(x))
}
