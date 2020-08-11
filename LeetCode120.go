package main

func minimumTotal(triangle [][]int) int {
	min := func(a int, b int) int {
		if a > b {
			return b
		}
		return a
	}

	dp := make([]int , len(triangle) + 1)
	for i := len(triangle) - 1 ; i >= 0 ; i -- {
		for j := 0 ; j < i ; j ++ {
			dp[j] = min(dp[j] , dp[j + 1]) + triangle[i][j]
		}
	}
	return dp[0]
}

func main() {
	
}
