package main

import (
	"fmt"
	"strings"
)

func longestSubstring(s string, k int) int {

	if s == "" {
		return 0
	}

	cnt := [26]int{}
	for _ , ch := range s {
		cnt[ch - 'a'] ++
	}

	var split byte
	for i , c := range cnt {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}

	if split == 0 {
		return len(s)
	}

	ans := 0
	for _ , substr := range strings.Split(s , string(split)) {
		ans = max(ans , longestSubstring(substr , k))
	}
	return ans
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestSubstring("aaabb" , 3))
}
