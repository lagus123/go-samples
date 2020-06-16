package main

import "fmt"

func main() {
	foo(2, 3, 4, 5, 6, 7, 8, 9)

	xi := []int{4, 8, 12}

	// Unfurling or destructuring slice
	foo(xi...)
}

// Variadic parameters: 0 or unlimited parameters
// Need to be final parameter
func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}

	fmt.Println(sum)

}
