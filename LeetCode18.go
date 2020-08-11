package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	sort.Ints(nums)
	result := [][]int{}
	for i := 0 ; i < len(nums) - 3 ; i ++ {
		// 根判断重复
		for i > 0 && nums[i - 1] == nums[i]  { i ++ }
		for j := i + 1 ; j < len(nums) - 2 ; j ++ {
			for  j != i + 1 && j < len(nums) -2 && nums[j - 1] == nums[j] { j ++ }
			left , right := j + 1 , len(nums) - 1
			for left < right {
				temp := nums[i] + nums[j] + nums[left] + nums[right]
				if temp == target {
					t := []int{ nums[i] , nums[j] , nums[left] , nums[right] }
					result = append( result , t)
					for left < right && nums[right] == nums[right - 1] { right -- }
					right --
					for left < right && nums[left] == nums[left + 1] { left ++ }
					left ++
				} else if temp > target {
					for left < right && nums[right] == nums[right - 1] { right -- }
					right --
				} else {
					for left < right && nums[left] == nums[left + 1] { left ++ }
					left ++
				}
			}
		}
	}
	return result
}
func main() {
	t := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(t,0))
}
