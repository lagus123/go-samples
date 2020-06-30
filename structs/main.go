package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Contact
}

type Contact struct {
	email string
	zip   int
}

type ByAge []Person

func (a ByAge) Len() int { return len(a)}
func (a ByAge) Swap(i, j int) {
	a[1], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func main() {
	// Basic sintax to create structs
	jhon := Person{"Jhon", "Laguna", 30, Contact{"jhonlag@gmail.com", 11001}}
	// With Properties names
	diana := Person{
		FirstName: "Diana",
		LastName:  "Martinez",
		Age:       27,
		Contact: Contact{
			email: "diana@gmail.com",
			zip:   11001,
		},
	}

	jhon.print()
	jhon.updateName("Edwar")
	jhon.print()

	people := []Person{
		jhon,
		diana,
	}

	fmt.Printf("\n%+v", people)
	// custom sort
	sort.Sort(ByAge(people))
	fmt.Printf("\n%+v", people)
}

func (p *Person) updateName(name string) {
	p.FirstName = name
}

func (p Person) print() {
	fmt.Printf("%+v\n", p)
}
