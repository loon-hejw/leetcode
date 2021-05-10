package main

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {

	max := func(a , b int) int {
		if a > b {
			return a
		}
		return b
	}

	sort.Ints(nums)
	sum := make([]int , nums[len(nums) - 1] + 1)
	for _ , v := range nums {
		sum[v] += v
	}

	f , s := sum[0] , max(sum[0], sum[1])
	for i := 2 ; i < len(sum) ; i ++ {
		f , s = s , max(f + sum[i] , s)
	}
	return s
}

func main() {
	fmt.Println(deleteAndEarn([]int{2,2,3,3,3,4}))
}
