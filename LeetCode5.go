package main

import "fmt"

func longestPalindrome(s string) string {

	n := len(s)
	ans := ""
	dp := make([][]int , n)
	for i := range dp {
		dp[i] = make([]int , n)
	}
	for i := 0 ; i < n ; i ++ {
		for j := 0 ; j + i < n ; j ++ {
			k := i + j
			if i == 0 {
				dp[j][k] = 1
			} else if i == 1 {
				if s[j] == s[k] {
					dp[j][k] = 1
				}
			} else {
				if s[j] == s[k] {
					dp[j][k] = dp[j + 1][k - 1]
				}
			}

			if dp[j][k] > 0 && i + 1 > len(ans) {
				ans = s[j : j + i + 1]
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}
