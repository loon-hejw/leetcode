package main

import "fmt"

func PredictTheWinner(nums []int) bool {

	n := len(nums)
	dp := make([][]int,n)
	for i := 0 ; i < len(nums) ; i++ {
		dp[i] = make([]int,n)
		dp[i][i] = nums[i]
	}
	dp[0][0] = 0
	max := func(a , b int) int {
		if a > b { return a }
		return b
	}
	for i := n - 2 ; i >= 0 ; i -- {
		for j := i + 1 ; j < n ; j ++ {
			dp[i][j] = max(nums[i] - dp[i + 1][j], nums[j] - dp[i][j - 1])
		}
	}
	return dp[0][n-1] >= 0
}

func main() {
	fmt.Println(PredictTheWinner([]int{1,5,233,7}))
}
