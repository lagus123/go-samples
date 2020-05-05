package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// Basic sintax to create structs
	jhon := person{"Jhon", "Laguna"}
	// With Properties names
	diana := person{firstName: "Diana", lastName: "Martinez"}

	fmt.Println(jhon.firstName, diana.firstName)
}
