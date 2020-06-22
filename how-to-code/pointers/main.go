package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // Get the mem address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // Type *int (pointer to int)

	b := &a
	fmt.Println(*b) // Get the value stored in the mem address

	x := 10
	foo(&x)
}

// Pass by Reference full sintax
func foo(x *int) {
	fmt.Println(x)
	fmt.Println(*x)

	*x = 100

	fmt.Println(x)
	fmt.Println(*x)
}