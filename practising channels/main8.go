package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int)

	go populate(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func populate(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			for v := range c {
				fmt.Println(v)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
