package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

var j int

func add() {
	time.Sleep(50 * time.Millisecond)
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		j++
		mutex.Unlock()
	}
}
func main() {
	for i := 0; i < 10; i++ {
		go add()
	}
	time.Sleep(5 * time.Second)
	fmt.Println(j)
}
