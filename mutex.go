package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex // Tambahin Mutex buat mengunci akses ke counter

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()   // Kunci akses ke counter
		counter++      
		mutex.Unlock() // Buka kunci biar goroutine lain bisa masuk
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go increment(&wg)
	go increment(&wg)
	go increment(&wg)

	wg.Wait()
	fmt.Println("Counter akhir:", counter) 
}
