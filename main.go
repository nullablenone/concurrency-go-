package main

import (
	"fmt"
	"time"
)

func task1() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 1 - Iterasi", i)
		time.Sleep(time.Second)
	}
}

func task2() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Task 2 - Iterasi", i)
		time.Sleep(time.Second)
	}
}

func main() {
	go task1() // Jalanin task1 sebagai Goroutine
	go task2() // Jalanin task2 sebagai Goroutine

	time.Sleep(6 * time.Second) // Tunggu supaya Goroutine selesai jalan
}
