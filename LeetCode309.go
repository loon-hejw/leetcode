package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	max := func (a int , b int) int {
		if a > b {
			return a
		}
		return b
	}
	dp := make([][]int,len(prices))
	dp[0] = append(dp[0],-prices[0])
	dp[0] = append(dp[0],0)
	dp[0] = append(dp[0],0)
	for i := 1 ; i < len(prices) ; i ++ {
		dp[i] = append(dp[i],max(dp[i - 1][0],dp[i - 1][2] - prices[i]));
		dp[i] = append(dp[i],dp[i - 1][0] + prices[i]);
		dp[i] = append(dp[i],max(dp[i - 1][1],dp[i - 1][2]))
	}
	return max(dp[len(prices) - 1][1],dp[len(prices) - 1][2]);
}

func proMax( a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(maxProfit([]int{1,2,3,0,2}))
}
