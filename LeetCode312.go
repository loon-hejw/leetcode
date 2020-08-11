package main

import "fmt"

func maxCoins(nums []int) int {

	dp  := make( [][]int , len(nums)  + 2 )
	vel := make( []int , len(nums) + 2 )
	vel[0] = 1
	vel[len(vel) - 1]  = vel[0]

	for i := 0 ; i < len(nums) + 2 ; i ++ {
		dp[i] = make([]int,len(nums) + 2)
	}
	for i := 1 ; i <= len(nums) ; i ++ {
		vel[i] = nums[i - 1]
	}
	max := func( a int ,b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := len(nums) - 1 ; i >= 0 ; i -- {
		for j := i + 2 ; j <= len(nums) + 1 ; j ++ {
			for k := i + 1 ; k < j ; k ++ {
				maxNumber := vel[i] * vel[j] * vel[k]
				maxNumber = maxNumber + dp[i][k] + dp[k][j]
				dp[i][j] = max(maxNumber , dp[i][j])
			}
		}
	}
	return  dp[0][len(nums) + 1]
}
func main() {
	fmt.Println(maxCoins([]int{3,1,5,8}))
}