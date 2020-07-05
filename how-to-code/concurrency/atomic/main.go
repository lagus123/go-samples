package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var g sync.WaitGroup
	var counter int64

	const gs = 100
	g.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter\t:", atomic.LoadInt64(&counter))
			g.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	g.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
