package main

/*
* https://play.golang.org/p/PtwO78Z7n3b
* Crear un programa para incrementar el valor de un contador
* Lanzar +10 goroutines
* Cada goroutine debe leer el valor del contador y guardarlo en una nueva variable
* Esperar 1 segundo
* Incrementar el valor de la nueva variable
* Escribir el valor de la nueva variable en el contador
* Usar waitgroups para esperar la finalización de cada goroutine
* Esto crea un race condition, probarlo usando el flag -race al correr el prograa
 */

import (
	"fmt"
	"runtime"
	"sync"
)

var waitg sync.WaitGroup

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	const gs = 100
	counter := 0
	waitg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			waitg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	waitg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
