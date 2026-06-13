package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(2)
	go kind()
	go jeph()

	fmt.Println("CPU\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Wait()
}

func kind() {
	for i := 0; i < 5; i++ {
		fmt.Printf("my name is kind and i have %v dogs\n", i)
	}
	defer wg.Done()
}

func jeph() {
	for i := 0; i < 5; i++ {
		fmt.Printf("my name is jeph and i have %v cats\n", i)
	}
	defer wg.Done()
}
