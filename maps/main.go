package main

import "fmt"

func main() {

	// Map ---> All Keys are the same type, All Values are the same type
	// Map always use square braces not dot properties

	// Method 2
	// var colors map[string]string

	//Method 3
	// colors := make(map[string]string)

	// Method 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	colors["white"] = "#ffffff"
	delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + " " + hex)
	}
}
