package main

import "fmt"

func maxProfit2(prices []int, fee int) int {

	lenght := len(prices)

	dp := make([][2]int , lenght)

	Max := func (a , b int) int {
		if a > b {
		 	return a
		}
		return b
	}
	dp[0][1] = -prices[0]
	for i := 1 ; i < lenght ; i ++ {
		dp[i][0] = Max(dp[i - 1][0] , dp[i - 1][1] + prices[i] - fee)
		dp[i][1] = Max(dp[i - 1][1] , dp[i - 1][0] -  prices[i])
	}
	return dp[lenght - 1][0]
}
// todo


func main() {
	fmt.Println(maxProfit2([]int{1,3,2,8,4,9} , 2))
}
