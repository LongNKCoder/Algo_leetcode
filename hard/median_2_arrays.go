package main

import (
	"fmt"
	"sort"
)

//Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.
//
//The overall run time complexity should be O(log (m+n)).
//
//
//
//Example 1:
//
//Input: nums1 = [1,3], nums2 = [2]
//Output: 2.00000
//Explanation: merged array = [1,2,3] and median is 2.
//Example 2:
//
//Input: nums1 = [1,2], nums2 = [3,4]
//Output: 2.50000
//Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.

func main() {
	fmt.Println(solution([]int{1, 2}, []int{3, 4}))
	fmt.Println(solution([]int{1, 2, 3, 4}, []int{5, 6}))
	fmt.Println(solution([]int{1, 3}, []int{2, 7}))
	fmt.Println(solution([]int{}, []int{3}))
}

func solution(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	switch {
	case len(nums) == 1:
		return float64(nums[0])
	case len(nums) == 0:
		return 0
	case len(nums)%2 == 0:
		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2
	default:
		return float64(nums[len(nums)/2])
	}
}
