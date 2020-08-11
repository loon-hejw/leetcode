package main

import "fmt"

func minPathSum(grid [][]int) int {
	Min := func( a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := grid
	for i := 0 ; i < len(grid) ; i ++ {
		for j := 0 ; j < len(grid[0]) ; j ++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if  i == 0 && j != 0  {
				dp[i][j] = grid[i][j] + grid[i][j - 1]
			} else if  i != 0 && j == 0 {
				dp[i][j] = grid[i][j] + grid[i - 1][j]
			} else {
				dp[i][j] = Min(grid[i - 1][j],grid[i][j - 1]) + grid[i][j]
			}
		}
	}
	return dp[len(grid) - 1][len(grid[0]) - 1]
}
func main() {
	grid := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	fmt.Println(minPathSum(grid))
}
