package main

import "fmt"

func isInterleave(s1 string, s2 string, s3 string) bool {

	m , n , k := len(s1) , len(s2) , len(s3)
	if m + n != k {
		return false
	}
	dp := make([][]bool,m + 1)
	for i := 0 ; i <= m ; i ++ {
		dp[i] = make([]bool,n + 1)
	}
	dp[0][0] = true
	for i := 0 ; i <= m ; i ++ {
		for j := 0 ; j <= n ; j ++ {
			p := i + j - 1
			if i > 0 {
				dp[i][j] = dp[i][j] || ( dp[i - 1][j] && s1[i - 1] == s3[p])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || ( dp[i][j - 1] && s2[j - 1] == s3[p])
			}
		}
	}
	return dp[m][n]
}
func main() {
	fmt.Println(isInterleave("aabcc","dbbca","aadbbcbcac"))
}
