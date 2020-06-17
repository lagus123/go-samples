package main

import "fmt"

func main() {
	fmt.Println(sum(2, 3, 4, 5, 6, 7, 8, 9))

	xi := []int{4, 8, 12, 15, 19}

	// Unfurling or destructuring slice
	fmt.Println(sum(xi...))

	// Anonymous function
	func() {
		fmt.Println("Anonymous function")
	}()

	func(x int) {
		fmt.Println("Anonymous function wit params", x)
	}(42)

	// Func expression
	f := func() {
		fmt.Println("Func expression")
	}
	f()

	x := sample()
	fmt.Println(x())
	// Alternative sintax
	fmt.Println(sample()())

	// Callback
	s2 := even(sum, xi...)
	fmt.Println(s2)
}

// Variadic parameters: 0 or unlimited parameters
// Need to be final parameter
func sum(x ...int) int {
	sum := 0

	for _, v := range x {
		sum += v
	}

	return sum

}

// Function return a function
func sample() func() int {
	return func() int {
		return 451
	}
}

// callback function
func even(f func(x ...int) int, vi ...int) int {
	var values []int

	for _, v := range vi {
		if v%2 == 0 {
			values = append(values, v)
		}
	}

	return f(values...)
}
