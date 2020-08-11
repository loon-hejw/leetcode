package main

import "fmt"

func uniquePaths(m int, n int) int {

	dp := make([][]int,n)
	// 初始化
	for i := 0 ; i < n ; i++{
		for j := 0 ; j < m ; j ++ {
			if i == 0 || j == 0 {
				dp[i] = append(dp[i],1)
			} else {
				dp[i] = append(dp[i],dp[i - 1][j] + dp[i][j - 1])
			}
		}
	}
	return dp[n - 1][m - 1]
}
func main() {
	fmt.Println(uniquePaths(3,2))
}
