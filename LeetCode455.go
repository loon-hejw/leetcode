package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	child := 0
	for i , j := 0 ,  0 ; i < len(g) && j < len(s) ; j ++ {
		if g[i] <= s[j] {
			i ++
			child ++
		}
	}
	return child
}
func main() {
	
}
