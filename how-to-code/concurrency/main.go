package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// Print OS Variables
	fmt.Println("OS\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t\t\t", runtime.GOARCH)
	fmt.Println("CPUS\t\t\t\t", runtime.NumCPU())
	fmt.Println("GOROUTINES\t\t\t", runtime.NumGoroutine())

	// Add for 1 Goroutine to finished
	wg.Add(1)

	go foo()
	bar()

	fmt.Println("GOROUTINES\t\t\t", runtime.NumGoroutine())

	// Wait until Goroutine is done
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	// Finish Goroutine
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
