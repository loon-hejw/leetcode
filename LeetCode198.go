package main

import "fmt"

func rob(nums []int) int {

	dp := make([]int,len(nums))
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[1]
	} else {
		Max := func( a int , b int) int {
			if a > b {
				return a
			}
			return b
		}
		dp[0] = nums[0]
		dp[1] = Max( nums[0] , nums[1] )
		for i := 2 ; i < len(nums) ; i ++ {
			dp[i] = Max(nums[i] + dp[i - 2],dp[i - 1])
		}
	}
	return dp[len(nums) - 1]
}
func main() {
	fmt.Println(rob([]int{1,2,3,1}))
}
