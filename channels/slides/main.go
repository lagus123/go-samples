package main

import (
	"fmt"
	"net/http"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	for _, link := range links {
		go checkLink(link)
	}

	// time.Sleep(time.Second * 3)
	// fmt.Println("done")

	// ==== CHANNELS =======

	// c := make(chan string)

	// for _, link := range links {
	// 	go checkLink(link, c)
	// }

	// fmt.Println(<-c) // Blocking code

	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// Dont pause main goroutine (the idea is main goroutine is always away)
	// If added in child goroutine we have to wait n seconds before the get call is executed, is not the best approach
	// We use goruotine anonymous (function literal)
	// Pause goroutines by n seconds before run execution

	// l is changed on the main goroutine, the bug is beacuse we are using a variable used by another goroutine ins this case the main goroutine

	// for l := range c {
	// 	go func() {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(l, c)
	// 	}()
	// }

	// so the solution is pass link value in function literal, so the scope of l is only on the function literal.
	// for l := range c {
	// 	go func(link string) {
	// 		time.Sleep(5 * time.Second)
	// 		checkLink(link, c)
	// 	}(l)
	// }

}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "is Down!")
		return
	}

	fmt.Println(link, "is Up!")
}

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		c <- link + " Is down!"
// 		return
// 	}

// 	c <- link + " Is up!"

// }

// func checkLink(link string, c chan string) {
// 	_, err := http.Get(link)

// 	if err != nil {
// 		fmt.Println(link, " Is Down")
// 		c <- link
// 		return
// 	}

// 	fmt.Println(link, " Is up!")
// 	c <- link

// }
