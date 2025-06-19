package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func send() {
	for i := 0; i < 10; i++ {
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
