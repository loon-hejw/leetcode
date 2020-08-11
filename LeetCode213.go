package main

import "fmt"

func rob2(nums []int) int {
	dp := make([]int,len(nums))
	Max := func( a int , b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2  {
		return Max(nums[0],nums[1])
	}else {
		dp[0] = nums[0]
		dp[1] = Max( nums[0] , nums[1] )
		for i := 2 ; i < len(nums) - 1 ; i ++ {
			dp[i] = Max(nums[i] + dp[i - 2],dp[i - 1])
		}
		dp[0] = dp[len(nums) - 2]
		dp[1] = nums[1]
		dp[2] = Max( nums[1] , nums[2] )
		for i := 3 ; i < len(nums) ; i ++ {
			dp[i] = Max(nums[i] + dp[i - 2],dp[i - 1])
		}
	}
	return Max(dp[0],dp[len(nums) - 1])
}
func main() {
	fmt.Println(rob2([]int{1,2,3,1}))
}
