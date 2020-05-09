package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())

}

func (englishBot) getGreeting() string {
	// Custom Logic
	return "Hello how are you!"
}

func (spanishBot) getGreeting() string {
	// Custom Logic
	return "Hola, como estas!"
}
