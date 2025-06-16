package main

import (
	"fmt"
	"strconv"
)

func work2(num int) bool {
	str := strconv.Itoa(num)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
func main() {
	b := work2(123454321)
	fmt.Println(b)
}
