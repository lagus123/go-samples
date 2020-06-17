package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

type agent struct {
	person
	ltk bool
}

// func (r receiver) identifier(parameters) (return(s)) { code }
func (s agent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
}

func main() {

	sa1 := agent{
		person: person{
			firstName: "Jhon",
			lastName:  "Laguna",
		},
		ltk:    false,
	}

	fmt.Println(sa1)
	sa1.speak()
}