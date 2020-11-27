package main

import (
	"fmt"
	"sort"
)

func maximumGap(nums []int) int {

	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := 0
	for i := 1 ; i < len(nums) ; i ++ {
		t := nums[i] - nums[i - 1]
		if t > max {
			max = t
		}
	}
	return max
}

func main() {
	fmt.Println(maximumGap([]int{1,2,4}))
}
