package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	removeElement(nums, val)
	fmt.Println(nums)
}

func removeElement(nums []int, val int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
