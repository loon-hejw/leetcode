package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	comt := []int{}
	res  := [][]int{}
	sort.Ints(candidates)

	var freq [][2]int
	for _ , num := range candidates {
		if freq == nil || num != freq[len(freq) - 1][0] {
			freq = append(freq,[2]int{num,1})
 		} else {
			freq[len(freq) - 1][1] ++
		}
	}

	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}

	var dfs func(target , index int)
	dfs = func(target, index int) {
		if target < 0 || index == len(freq) {
			return
		}
		if target == 0 {
			t := make([]int,len(comt))
			copy(t,comt)
			res = append(res,t)
			return
		}
		dfs(target , index + 1)

		most := min(target/freq[index][0] , freq[index][1])
		for i := 1 ; i <= most ; i++ {
			comt = append(comt,freq[index][0])
			dfs(target - i * freq[index][0] , index + 1)
		}
		comt = comt[:len(comt)-most]
	}
	dfs(target,0)
	return res
}


func main() {
	fmt.Println(combinationSum2([]int{10,1,2,7,6,1,5},8))
}
