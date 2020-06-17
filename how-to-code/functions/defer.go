package main

import "fmt"

func main() {
	// Runs at the end of current routine
	defer test()
	bar()
}

func test() {
	fmt.Println("Test")
}

func bar() {
	fmt.Println("Bar")
}
