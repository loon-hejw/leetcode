package main

import "fmt"

func minimumOperations(leaves string) int {

	n := len(leaves)
	dp := make([][3]int,n)
	isY := func(b bool) int {
		if b { return 1 }
		return 0
	}
	min := func(a ,b int) int {
		if a < b { return a }
		return b
	}
	dp[0][0] = isY(leaves[0] == 'y')
	dp[0][1] , dp[0][2] , dp[1][2] = n , n, n
	for i := 1 ; i < n ; i ++ {
		dp[i][0] = dp[i - 1][0] + isY(leaves[i] == 'y')
		dp[i][1] = min (dp[i - 1][0] , dp[i - 1][1]) + isY(leaves[i] == 'r')
		if i >= 2 {
			dp[i][2] = min(dp[i - 1][1] , dp[i - 1][2]) + isY(leaves[i] == 'y')
		}
	}
	return dp[n - 1][2]
}

func main() {
	fmt.Println(minimumOperations("rrryyyrryyyrr"))
}
