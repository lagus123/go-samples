package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)
	receive(c)

	fmt.Println("Exit")
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
