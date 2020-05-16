package main

import (
	"fmt"
	"net/http"
)

// Channels
// Deben tener un tipo, este tipo define la información que puede ser transportada por el canal
// Son la unica forma de comunicar Go Routines
// Los mensajes a un canal son bloqueantes

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

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkStatus(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {

		c <- link + " Website is Down"
		return
	}

	c <- link + " is Up!"
}
