package main

import (
	"fmt"
	"runtime"
	"sync/atomic"

	"sync"
)

func main() {
	var wg sync.WaitGroup

	var counter int64 = 0

	const gr = 100

	wg.Add(gr)

	//var mu sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			// mu.Lock()
			// v := counter
			// //runtime.Gosched()
			// v++
			// counter = v
			// mu.Unlock()
			// wg.Done()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()
		fmt.Println("Goroutines\t", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	fmt.Println("Counter:", counter)
}
