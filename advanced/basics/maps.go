package main

import "fmt"

func main() {

	m := map[string]int{
		"John":  30,
		"Diana": 27,
	}

	fmt.Println(m)
	fmt.Println(m["John"])
	// If the values is not present a map return Zero Value
	fmt.Println(m["James"])

	// For that we need use a trick call "comma ok" idiom
	if _, ok := m["John"]; ok {
		fmt.Println(m["John"])
	}

	// Print key - values
	for k, v := range m {
		fmt.Println(k, v)
	}

}
