package main

import "fmt"

func integerBreak(n int) int {
	dp := make([]int,n + 1)
	max := func ( a , b int ) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 2 ; i <= n ; i++ {
		count := 0
		for j := 1 ; j < i ; j ++ {
			count =  max(max( j * ( i - j ) , j * dp[i - j]),count)
		}
		dp[i] = count
	}
	return dp[n]
}
func main() {
	fmt.Println(integerBreak(10))
}