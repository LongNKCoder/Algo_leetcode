package main

//Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique element appears only once. The relative order of the elements should be kept the same. Then return the number of unique elements in nums.
//
//Consider the number of unique elements of nums to be k, to get accepted, you need to do the following things:
//
//Change the array nums such that the first k elements of nums contain the unique elements in the order they were present in nums initially. The remaining elements of nums are not important as well as the size of nums.
//Return k.

func main() {
	nums := []int{1, 1, 2}
	removeDuplicates(nums)
}

func removeDuplicates(nums []int) int {
	leftPointer := 1
	for rightPointer := 1; rightPointer < len(nums); rightPointer++ {
		if nums[rightPointer] != nums[rightPointer-1] {
			nums[leftPointer] = nums[rightPointer]
			leftPointer++
		}
	}
	return leftPointer
}
