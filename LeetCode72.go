package main

import "fmt"

func minDistance(word1 string, word2 string) int {

	m , n := len(word1) , len(word2)
	dp := make([][]int , m + 1)
	for i := range dp {
		dp[i] = make([]int , n + 1)
		dp[i][0] = i
	}

	for i := 0 ; i < n + 1 ; i ++ {
		dp[0][i] = i
	}

	min := func(a , b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 1 ; i < m + 1 ; i ++ {
		for j := 1 ; j < n + 1 ; j ++ {
			dp[i][j] = min(min(dp[i][j - 1] , dp[i - 1][j]) , dp[i - 1][j - 1]) + 1
			if word2[j - 1] == word1[i - 1] {
				dp[i][j] = min(dp[i][j] , dp[i - 1][j - 1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	fmt.Println(minDistance("mouuse" , "mouse"))
}