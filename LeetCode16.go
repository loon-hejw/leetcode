package main

import (
	"fmt"
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {

	sort.Ints(nums)

	roop := nums[0] + nums[1] + nums[len(nums) - 1]
	for i := 0 ; i < len(nums) - 2 ; i ++ {
		left , right := i + 1 , len(nums) - 1

		for left < right {
			num := nums[i] + nums[left] + nums[right]

			if num == target {
				return num
			} else if num < target {
				for left < right && nums[left] == nums[left + 1] { left ++ }
				left ++
			} else {
				for left < right && nums[right] == nums[right - 1] { right -- }
				right --
			}

			if abs(roop - target) > abs(num - target) {
				roop = num
			}
		}
	}

	return roop
}
func main() {
	fmt.Println(threeSumClosest([]int{-1,2,1,-4},1))
}
