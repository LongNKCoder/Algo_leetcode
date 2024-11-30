package main

import "fmt"

func getConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	j := 0
	for i := 0; i < len(ans); i++ {
		if j > len(nums)-1 {
			j = 0
		}
		ans[i] = nums[j]
		j++
	}
	return ans
}

func main() {
	nums := []int{1, 2, 1}
	fmt.Println(getConcatenation(nums))
}
