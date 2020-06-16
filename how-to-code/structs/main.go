package main

import "fmt"

func main() {

	// Anonymous Struct
	p1 := struct {
		first string
		last  string
	}{
		first: "Jhon",
		last:  "Laguna",
	}

	fmt.Println(p1)
}
