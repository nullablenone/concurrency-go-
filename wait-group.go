package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(name string, wg *sync.WaitGroup) {
	defer wg.Done() // Kurangi counter setelah selesai

	for i := 1; i <= 3; i++ {
		fmt.Println("Hello", name, i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup // Bikin WaitGroup

	wg.Add(2) // Kita bakal punya 2 goroutine

	go sayHello("Alice", &wg)
	go sayHello("Bob", &wg)

	wg.Wait() // Tunggu semua goroutine selesai
	fmt.Println("Semua goroutine selesai!")
}
