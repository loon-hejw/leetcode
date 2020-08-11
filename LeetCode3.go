package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	strArr := [26]int{}
	lenght , rk := 0 , -1
	max := func( a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 0 ; i < len(s) ; i ++ {
		if i != 0 {
			strArr[s[i - 1] - 'a'] = 0
		}
		for rk + 1 < len(s) && strArr[s[rk + 1] - 'a'] == 0 {
			strArr[s[rk + 1] - 'a'] ++
			rk ++
		}
		lenght = max(lenght , rk - i + 1)
	}
	return lenght
}
func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
