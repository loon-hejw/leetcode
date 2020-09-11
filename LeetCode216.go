package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	res  := [][]int{}
	comt := []int{}
	var dfs func(star , end  int )
	dfs = func(star, end int ) {
		if end == 0 && len(comt) == k{
			t := make([]int,len(comt))
			copy(t,comt)
			res = append(res , t)
			return
		}
		if star > 9 || star > n - k {
			return
		}
		dfs( star + 1 , end )

		comt = append(comt , star)
		dfs( star + 1 , end - star)
		comt = comt[:len(comt) - 1]
	}

	dfs(1 ,  n )
	return res
}

func main() {
	fmt.Println(combinationSum3(2,6))
}
