package main

import (
	"fmt"
	"sort"
)

var bit = [1e4 + 1]int{}

func init() {
	for i := 1 ; i <= 1e4 ; i ++ {
		bit[i] = bit[i >> 1] + i&1
	}
}

func sortByBits(arr []int) []int {
	sort.Slice(arr , func(i, j int) bool {
		x , y := arr[i] , arr[j]
		cx , cy := bit[x] , bit[y]
		return cx < cy || cx ==  cy && x < y
	})
	return arr
}

func main() {
	fmt.Println(sortByBits([]int{0,1,2,3,4,5,6,7,8}))
}
