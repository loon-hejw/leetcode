package main

import "fmt"

func numSubarraysWithSum(nums []int, goal int) int {

	var (
	 	cnt = map[int]int{}
	 	res = 0
	 	sum = 0
	)
	for _ , num := range nums {
		cnt[sum] ++
		sum = sum + num
		res = res + cnt[sum - goal]
	}
	return res
}

func main() {
	fmt.Println(numSubarraysWithSum([]int{1,0,1,0,1} , 2))
}
