package main

/*
* https://play.golang.org/p/ZK7S5WgO54z
* Lanzar dos goroutines aparte de la principal, cada una debe imprimir un mensaje
* Utilizando waitgroups asegurarse que cada una es ejecutada antes de que el programa finalice
 */

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// Print OS Variables
	fmt.Println("CPUS\t\t\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t\t\t", runtime.NumGoroutine())

	// Add for 1 Goroutine to finished
	wg.Add(1)

	go foo()

	fmt.Println("GOROUTINES\t\t\t", runtime.NumGoroutine())

	// Wait until Goroutine is done
	wg.Wait()
}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo:", i)
	}
	// Finish Goroutine
	wg.Done()
}
