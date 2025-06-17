package main

import "fmt"

func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	digits = append([]int{1}, digits...)
	return digits
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
}
