package main

import "fmt"

func countBinarySubstrings(s string) int {
	cur , pre , res := 1 , 0 , 0
	min := func (a ,b int) int {
		if a > b {
			return b
		}
		return a
	}
	for i := 1 ; i < len(s) ; i ++ {
		if s[i] == s[i - 1] {
			cur ++
		} else {
			res = res + min(cur,pre)
			pre = cur
			cur = 1
		}
	}
	return res + min(cur,pre)
}


func main() {
	fmt.Println(countBinarySubstrings("00110011"))
}
