package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var j int64

func add(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		atomic.AddInt64(&j, 1)
	}
}
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(&wg)
	}
	wg.Wait()
	fmt.Println(j)
}
