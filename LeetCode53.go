package main

import "fmt"

func maxSubArray(nums []int) int {
	getMax := func( a int , b int) int {
		if a > b {
			return a
		}
		return b
	}
	max , pre := nums[0] , 0
	for i := 0 ; i < len(nums) ; i ++ {
		pre = getMax(pre + nums[i] , nums[i])
		max = getMax(pre , max)
	}
	return max
}
func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}
