package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	comt := []int{}
	res  := [][]int{}
	sort.Ints(candidates)
	var dfs func(target , index int)
	dfs = func(target, index int) {
		if target < 0 || index == len(candidates) {
			return
		}
		if target == 0 {
			t := make([]int,len(comt))
			copy(t,comt)
			res = append(res,t)
			return
		}
		dfs(target , index + 1)
		if target - candidates[index] >= 0 {
			comt = append(comt,candidates[index])
			dfs(target - candidates[index] , index)
			comt = comt[:len(comt)-1]
		}
	}
	dfs(target,0)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2,3,6,7},7))
}
