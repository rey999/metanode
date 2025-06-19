package main

import (
	"fmt"
	"time"
)

var ch = make(chan int, 100)

func send() {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
}

func read() {
	for val := range ch {
		fmt.Println(val)
	}
}
func main() {
	go send()
	go read()
	time.Sleep(1 * time.Second)
}
