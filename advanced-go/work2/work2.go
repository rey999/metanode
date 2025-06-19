package main

import (
	"fmt"
)

func m2(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	m2(arr)
	fmt.Println(arr)
}
