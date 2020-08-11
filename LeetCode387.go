package main

import "fmt"

func firstUniqChar(s string) int {
	strMap := [26]int{}
	for _ , v := range s {
		strMap[int(v) - 'a'] ++
	}
	for k , v := range s {
		if strMap[int(v) - 'a'] == 1 {
			return k
		}
	}
	return -1
}
func main() {
	fmt.Println(firstUniqChar("leetcode"))
}
