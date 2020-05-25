package main

import "fmt"

var a int

type custom int

var b custom

func main() {

	a = 42
	b = 43
	fmt.Println(a)
	fmt.Println(b)

	// Convert a to type custom, in Go cast not exists
	b = custom(a)
	fmt.Println(b)
}
