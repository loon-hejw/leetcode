package main

import "sort"

func maximumProduct(nums []int) int {
	sort.Ints(nums)
	a := nums[0] * nums[1] * nums[len(nums) - 1]
	b :=  nums[len(nums) - 3] * nums[len(nums) - 2] * nums[len(nums) - 1]
	if a > b {
		return a
	}
	return b
}

func main() {
	
}
