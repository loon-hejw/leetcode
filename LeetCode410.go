package main

import (
	"fmt"
	"math"
)

func splitArray(nums []int, m int) int {
	dp := make([][]int,len(nums) + 1)
	for i := 0 ; i < len(nums) + 1 ; i ++ {
		dp[i] = make([]int,m + 1)
		for j := 0 ; j < m + 1 ; j ++ {
			dp[i][j] = math.MaxInt64
		}
	}
	dp[0][0] = 0
	sub := make([]int,len(nums) + 1)
	for i := 0 ; i < len(nums) ; i ++ {
		sub[i + 1] = sub[i] + nums[i]
	}
	min := func(a int , b int) int {
		if a > b {
			return b
		}
		return a
	}
	max := func(a int , b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1 ; i < len(dp) ; i ++ {
		for j := 1 ; j <= min(i,m) ; j ++ {
			for k := 0 ; k < i ; k ++ {
				dp[i][j] = min( dp[i][j] , max( dp[k][j - 1] , sub[i] - sub[k]))
			}
		}
	}
	return dp[len(nums)][m]
}
func main() {
	fmt.Println(splitArray([]int{7,2,5,10,8},2));
}
