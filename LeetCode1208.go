package main

import "fmt"

func equalSubstring(s string, t string, maxCost int) int {

	quenu := []int{}
	for k , v := range s {
		quenu = append( quenu , abs(int(v) - int(t[k])))
	}
	maxLen := 0
	sum , start := 0 , 0
	for k , v := range quenu {
		sum += v
		if sum > maxCost {
			sum -= quenu[start]
			start = k
		}
		maxLen = max(maxLen ,k - start + 1)
	}
	return maxLen
}


func abs (a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func max (a , b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(equalSubstring("abcd","bcdf",3))
}
