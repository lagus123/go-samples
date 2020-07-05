package main

/*
* https://play.golang.org/p/pjK0PJHOrSr
* Lanzar dos goroutines aparte de la principal, cada una debe imprimir un mensaje
* Utilizando waitgroups asegurarse que cada una es ejecutada antes de que el programa finalice
 */

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go func() {
		fmt.Println("Hello from thing 1")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing 2")
		wg.Done()
	}()

	wg.Wait()
}
