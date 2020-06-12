package main

import "fmt"

func main() {

	m := map[string][]string{
		"Jhon":  []string{"Movie", "Culture"},
		"Diana": []string{"Dance", "Culture"},
	}

	m["Jess"] = []string{"Movies", "Medicine"}

	delete(m, "Diana")

	for k, v := range m {
		fmt.Println(k)
		for _, v2 := range v {
			fmt.Print(v2 + "\t")
		}
		fmt.Print("\n \n")
	}

}
