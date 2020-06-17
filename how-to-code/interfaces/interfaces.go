package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type agent struct {
	person
	ltk bool
}

type human interface {
	speak()
}

type test int

func (s agent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
}

func (p person) speak() {
	fmt.Println("I am", p.firstName, p.lastName)
}

func bar(h human) {

	switch h.(type) {
	case person:
		fmt.Println("I am a person with name", h.(person).firstName) // Assertion
	case agent:
		fmt.Println("I am a agent with name", h.(agent).firstName)
	default:
		fmt.Println("I am a human with name", h)
	}
}

func main() {

	sa1 := agent{
		person: person{
			firstName: "Jhon",
			lastName:  "Laguna",
		},
		ltk: false,
	}

	fmt.Println(sa1)
	sa1.speak()

	bar(sa1)

	p1 := person{
		firstName: "John",
		lastName:  "Doe",
	}

	fmt.Println(p1)
	p1.speak()

	bar(p1)

	// Conversion
	// Convert type T with underline type S to type
	var x test = 42
	var y int
	y = int(x)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

}
