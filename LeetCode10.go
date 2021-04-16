package main

import "fmt"

func isMatch(s string, p string) bool {
	m , n := len(s) , len(p)
	dp := make([][]bool  , m + 1)
	for k := range dp {
		dp[k] = make([]bool , n + 1)
	}
	dp[0][0] = true
	for i := 0 ; i <= m ; i ++ {
		for j := 1 ; j <= n ; j ++ {
			if p[j - 1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j - 2]
				if boolMath( i , j - 1 , s, p) {
					dp[i][j] = dp[i][j] || dp[i - 1][j]
				}
			} else if boolMath( i , j , s, p) {
				dp[i][j] = dp[i][j] || dp[i - 1][j - 1]
			}
		}
	}
	return dp[m][n]
}

func boolMath (i , j  int , s , p string) bool {
	if i == 0 {
		return false
	}
	if p[j - 1] == '.' {
		return true
	}
	return  s[i - 1] == p[j - 1]
}

func main() {
	fmt.Println(isMatch("aabbca" , "a.*a"))
}
