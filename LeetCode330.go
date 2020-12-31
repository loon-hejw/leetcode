package main

import "fmt"

func minPatches(nums []int, n int) int {
	p := 0
	for i , x := 0 , 1 ; x <= n ; {
		if i < len(nums) && nums[i] <= x {
			x = x + nums[i]
			i ++
		} else {
			x = x * 2
			p ++
		}
	}
	return p
}

func main() {

	fmt.Println(minPatches([]int{1,5,10}, 5))
}
