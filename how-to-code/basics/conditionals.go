package main

import "fmt"

func main() {

	if true {
		fmt.Println("True")
	}

	// Not operator
	if !false {
		fmt.Println("True")
	}

	// Switch on boolean value
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
	default:
		fmt.Println("This is default value")
	}

	// switch on value (Multiple values)
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}

}
