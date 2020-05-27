package main

import "fmt"

// Bit shifting and iota
const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	fmt.Printf("%b\n", kb)
	fmt.Printf("%b\n", mb)
	fmt.Printf("%b\n", gb)
}
