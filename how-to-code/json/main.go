package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type Result []struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Age   int    `json:"age"`
}

func main() {

	p1 := Person{
		"Jhon",
		"Laguna",
		30,
	}

	p2 := Person{
		"Pedro",
		"Perez",
		27,
	}

	people := []Person{p1, p2}
	fmt.Println(people)

	b, err := json.Marshal(people)
	check(err)
	fmt.Println(string(b))

	// Using Result from Marshall
	var result []Person
	json.Unmarshal(b, &result)
	fmt.Printf("%+v\n", result)

	// Using Slice of Bytes
	s := []byte(`[{"first":"Jhon","last":"Laguna","age":30},{"first":"Pedro","last":"Perez","age":27}]`)
	var response Result
	err = json.Unmarshal(s, &response)
	check(err)
	fmt.Printf("%+v", response)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
