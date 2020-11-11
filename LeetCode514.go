package main

import "fmt"

func findRotateSteps(ring string, key string) int {

	m := len(ring)
	dp := make([][]int, m)
	for i := 0 ; i < len(dp) ; i ++ {
		dp[i] = make([]int , len(key))
	}
	ringMap := make([][]int,26)
	for k , v := range ring {
		p := v - 'a'
		ringMap[p] = append(ringMap[p] , k)
	}

	abs := func(i int) int {
		if i > 0 {
			return i
		}
		return  -1 * i
	}

	min := func( i , j int) int {
		if i > j {
			return j
		}
		return i
	}

	var dfs func(i , j int) int
	dfs = func( i , j int) int {
		if j == len(key) {
			return 0
		}
		if dp[i][j] > 0 {
			return dp[i][j]
		}
		dp[i][j] = 1 << 31
		for _ , v := range ringMap[key[j] - 'a'] {
			diff := abs( i - v )
			diff = min(diff , m - diff)
			dp[i][j] = min(dp[i][j] , diff + dfs(v , j + 1))
		}
		return dp[i][j]
	}
	return dfs(0 , 0) + len(key)
}

func main() {
	fmt.Println(findRotateSteps("godding","gd"))
}
