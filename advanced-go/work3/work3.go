package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func oddNum() {
	// defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			fmt.Println(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func evenNum() {
	// defer wg.Done()
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func main() {
	// wg.Add(2)
	go oddNum()
	go evenNum()
	// wg.Wait()
	time.Sleep(1 * time.Second)
}
