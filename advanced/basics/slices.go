package main

import "fmt"

func main() {
	// Values from same types
	x := []int{1, 3, 5, 9}

	for _, v := range x {
		fmt.Println(v)
	}

	// Slice values from slice
	// [pos1: pos2] pos1 start position pos2 final position, upto but not includind
	fmt.Println(x[0:3])

	// Append values to slice
	x = append(x, 4, 7)
	fmt.Println(x)
}
