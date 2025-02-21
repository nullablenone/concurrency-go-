package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "menambah data ke channel"
		fmt.Println("other go routines")
	}()

	fmt.Println(<-ch)
}
