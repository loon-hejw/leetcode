package main

import "fmt"

func largeGroupPositions(s string) [][]int {

	ans := [][]int{}
	index := 0
	for index < len(s) {
		start := index
		for index < len(s) && s[start] == s[index] {
			index ++
		}
		if index - start >= 3 {
			ans = append(ans , []int{start , index - 1})
		}
	}
	return ans
}

func main() {
	fmt.Println(largeGroupPositions("aaa"))
}
