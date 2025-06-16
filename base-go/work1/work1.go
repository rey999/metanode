package main

import "fmt"

func singleNumber(nums []int) int {
	// fmt.Println(nums)
	nmap := make(map[int]int)

	for _, v := range nums {
		if _, e := nmap[v]; e {
			nmap[v] = nmap[v] + 1
		} else {
			nmap[v] = 1
		}
	}
	fmt.Println(nmap)
	for k, v := range nmap {
		if v == 1 {
			return k
		}
	}
	return 0
}

func main() {
	nums := []int{1, 2, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println(singleNumber(nums))
}
