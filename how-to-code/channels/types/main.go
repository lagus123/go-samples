package main

import "fmt"

func main() {
	c := make(chan int)
	cso := make(chan<- int) // Send only channel
	cro := make(<-chan int) // Receive only channel

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cso)
	fmt.Printf("%T\n", cro)
}
