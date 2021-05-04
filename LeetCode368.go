package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {

	sort.Ints(nums)
	n := len(nums)
	dp := make([]int , n)

	for i := range dp {
		dp[i] = 1
	}

	maxsize , maxval := 1 , 1
	for i := 1 ; i < n ; i ++ {
		for j , v := range nums[:i] {
			if nums[i] % v == 0 && dp[j] + 1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxsize {
			maxsize , maxval = dp[i] , nums[i]
 		}
	}

	if maxsize == 1 {
		return []int{nums[0]}
	}

	res := []int{}
	for i := n - 1 ; i >= 0 && maxsize > 0 ; i -- {
		if dp[i] == maxsize && maxval % nums[i] == 0 {
			res = append( res , nums[i])
			maxval = nums[i]
			maxsize --
		}
	}

	return res
}

func main() {
	fmt.Println(largestDivisibleSubset([]int{1,2,4,8}))
}
