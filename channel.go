package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		ch <- "menambah data ke channel 1"
	}()
	go func() {
		ch <- "menambah data ke channel 2"
	}()
	go func() {
		ch <- "menambah data ke channel 3"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
