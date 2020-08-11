package main

import (
	"fmt"
)

func d(a int) int {
	if ( a == 1 || a == 2) {
		return 1
	}
	return d(a - 1) + d(a - 2)
}

func dp2(a int , dp []int) {
	if a == 1 || a == 2 {
		dp[a] = 1
		return
	}
	if (dp[a] != 0 ) {
		return
	}
	dp2(a - 1 ,dp)
	dp2(a - 2 ,dp)
	dp[a] = dp[ a - 1 ] + dp[ a - 2 ]
	return
}

func main() {
	dp := make([]int,8)
	dp2(7,dp)
	fmt.Println(dp)
}
