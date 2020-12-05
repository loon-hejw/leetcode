package main

import "fmt"

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1 , -1}
	}
	left , right := 0 , len(nums) - 1
	index := 0
	for left <= right {
		t := (left + right) >> 1
		if nums[t] == target {
			index = t
			break
		} else if nums[t] > target {
			right = t - 1
		} else {
			left  = t + 1
		}
	}

	i1 , i2 := index , index
	for i1 >= 0 && nums[i1] == target {
		i1 --
	}
	for i2 < len(nums) &&nums[i2] == target {
		i2 ++
	}

	if i1 == i2 && i1 == 0 {
		return []int{-1 , -1}
	}
	return []int{i1 + 1 , i2 - 1}
}

func main() {
	fmt.Println(searchRange([]int{5,7,7,8,8,10}, 8))
}
