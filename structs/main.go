package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	contact
}

type contact struct {
	email string
	zip   int
}

func main() {
	// Basic sintax to create structs
	jhon := person{"Jhon", "Laguna", contact{"jhonlag@gmail.com", 11001}}
	// With Properties names
	diana := person{
		firstName: "Diana",
		lastName:  "Martinez",
		contact: contact{
			email: "diana@gmail.com",
			zip:   11001,
		},
	}

	jhon.print()
	diana.print()

	jhon.updateName("Edwar")
	jhon.print()
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
