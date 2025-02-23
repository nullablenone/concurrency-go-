package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "ini dari goroutine 1"
	}()
	go func() {
		time.Sleep(time.Second)
		ch2 <- "ini dari goroutine 2"
	}()

	select {
	case pesan1 := <-ch1:
		fmt.Println(pesan1)
	case pesan2 := <-ch2:
		fmt.Println(pesan2)
	default:
		fmt.Println("gak tau")
	}

}
