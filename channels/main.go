package main

import (
	"fmt"
	"net/http"
	"time"
)

// Channels
// Deben tener un tipo, este tipo define la información que puede ser transportada por el canal
// Son la unica forma de comunicar Go Routines
// Los mensajes a un canal son bloqueantes
// No tratar de usar shared variables en routines

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://amazon.com",
		"http://golang.org",
	}

	c := make(chan string)

	for _, url := range links {
		go checkStatus(url, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "Website is Down!")
		c <- link
		return
	}

	fmt.Println(link, "is Up!")
	c <- link
}
