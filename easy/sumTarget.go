package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(sumTarget([]int{1, 2, 3, 4, 2, 5, 1}, 9))  // true
	fmt.Println(sumTarget([]int{1, 2, 3, 4, 2, 5, 1}, 4))  // true
	fmt.Println(sumTarget([]int{1, 2, 3, 4, 2, 5, 1}, 3))  // true
	fmt.Println(sumTarget([]int{1, 2, 3, 4, 2, 5, 1}, 10)) // false
	fmt.Println(sumTarget([]int{1, 2, 3, 4, 2, 5, 1}, 1))  // false
}

func sumTarget(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	sort.Ints(nums)
	for left < right {
		switch {
		case nums[left]+nums[right] == target:
			return true
		case nums[left]+nums[right] < target:
			left++
		case nums[left]+nums[right] > target:
			right--
		}
	}
	return false
}
