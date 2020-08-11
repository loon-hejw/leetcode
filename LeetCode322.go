package main

import (
	"fmt"
)

func coinChange(coins []int, amount int) int {
	dp := make([]int,amount + 1)
	dp[0] = 0
	for i := 1 ; i < len(dp) ; i ++ {
		dp[i] = amount + 1
	}

	for i := 1 ; i <= amount ; i ++{
		for j := 0 ; j < len(coins) ; j ++ {
			if i > coins[j] {
				continue
			}
			dp[i] = getMin(dp[i],1 + dp[ i - coins[j]])
		}
	}
	if dp[amount] == amount + 1 {
		return -1
	}
	return dp[amount]
}

func getMin(a int , b int ) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(coinChange([]int{2},3))
}
