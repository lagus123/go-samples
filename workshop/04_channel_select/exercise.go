package main

import (
	"fmt"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go send(eve, odd, quit)

	receive(eve, odd, quit)

	fmt.Println("Exit")

}

func send(e, o, q chan int) {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}
	q <- 0
}

func receive(e, o, q chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("Even channel", v)
			// logic
		case v := <-o:
			fmt.Println("Odd channel", v)
			// logic
		case v := <-q:
			fmt.Println("Quit channel", v)
			return
		}
	}
}
