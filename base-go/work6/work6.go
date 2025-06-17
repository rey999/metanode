package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var slow, fast int

	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}

	return slow + 1
}

func main() {
	nums := []int{1, 1, 2}
	length := removeDuplicates(nums)
	fmt.Println(length, nums[:length])
}
