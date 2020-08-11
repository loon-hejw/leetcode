package main

import "fmt"

func maxProduct(nums []int) int {

	getMax := func( a int , b int ) int {
		if a > b {
			return a
		}
		return b
	}

	getMin := func( a int , b int ) int {
		if a < b {
			return a
		}
		return b
	}
	min , max , perMax   := nums[0] , nums[0] , nums[0]
	for i := 1 ; i < len(nums) ; i ++ {
		mi , ma := min , max
		max = getMax(ma * nums[i] , getMax(nums[i] * mi , nums[i]))
		min = getMin(mi * nums[i] , getMin(nums[i] * ma , nums[i]))
		perMax = getMax(max , perMax)
	}
	return perMax
}

func main() {
	fmt.Println(maxProduct([]int{-4,-3,-2}))
}
