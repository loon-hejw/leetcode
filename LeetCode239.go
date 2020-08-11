package main

import (
	"sort"
)

func maxSlidingWindow(nums []int, k int) []int {

	// 排除出可以直接返回的数组
	if len(nums) == 0 || k == 1 { return nums }

	// 判断数组长度
	if( len(nums) < k ) {
		sort.Ints(nums)
		return nums[len(nums) - 1 : len(nums)]
	}

	// 设置遍量
	var area []int
	var temp []int

	// 开始切片
	for i := 0 ; i < len(nums) ; i ++ {

		for len(temp) != 0 && temp[0] < i - k + 1 {
			temp = temp[1:]
		}

		for len(temp) != 0 && nums[temp[len(temp) - 1]] < nums[i] {
			temp = temp[:len(temp)-1]
		}

		temp = append(temp,i)

		if i >= k - 1 {
			area = append(area,nums[temp[0]])
		}

	}

	return area
}

func main() {
	maxSlidingWindow([]int{7,2,4},2)
}
