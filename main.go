package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("my first goroutine")
	go func() {
		fmt.Println("my other goroutine")
	}()
	go func() {
		fmt.Println("my other goroutine")
	}()
	time.Sleep(time.Second)
}
