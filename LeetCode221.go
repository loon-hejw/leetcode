package main

import "fmt"

func maximalSquare(matrix [][]byte) int {
	dp := make([][]int , len(matrix))
	for i := 0 ; i < len(matrix) ; i ++ {
		dp[i] = make([]int,len(matrix[i]))
	}
	maxSquare := 0
	min := func(a ,b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func( a , b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0 ; i < len(matrix) ; i ++ {
		for j := 0 ; j < len(matrix[0]) ; j ++ {
			if matrix[i][j] == '1' {
				if j == 0 || i == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = min (min(dp[i - 1][j] , dp[i][j - 1] ), dp[i - 1][j - 1]) + 1
				}
				maxSquare  = max(maxSquare , dp[i][j])
			}
		}
	}
	return maxSquare * maxSquare
}
func main() {
	arr := [][]byte{ {'1','0','1','0','0'} , {'1','0','1','1','1'} , {'1','1','1','1','1'} , {'1','0','0','1','0'} }
	fmt.Println(maximalSquare(arr))
}