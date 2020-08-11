package main

import "fmt"

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	dp , m := make([]int,len(obstacleGrid[0])) , len(obstacleGrid[0])
	dp[0] = 1
	for _ , v := range obstacleGrid {
		for i := 0 ; i < m ; i ++ {
			if v[i] == 1 {
				dp[i] = 0
			} else if i > 0 {
				dp[i] = dp[i] + dp[i - 1]
			}
		}
	}
	return dp[m - 1]
}
func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{{0,0,0},{0,1,0},{0,0,0}}))
}
