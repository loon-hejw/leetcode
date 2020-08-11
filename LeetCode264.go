package main

import (
	"fmt"
	"math"
)

func nthUglyNumber(n int) int {
	nums := make([]int,n)

	index2 , index3 , index5 := 0 ,0 ,0
	nums[0] = 1

	for i := 1 ; i < len(nums); i ++ {
		nums[i] = maxThree(nums[index2] * 2,nums[index3] * 3,nums[index5] * 5)
		if nums[i] == nums[index2] * 2 {
			index2 ++
		}
		if nums[i] == nums[index3] * 3 {
			index3 ++
		}
		if nums[i] == nums[index5] * 5 {
			index5 ++
		}
	}
	return nums[n - 1]
}

func maxThree(index2 int,index3 int ,index5 int ) int {
	return int(math.Min(float64(index2),math.Min(float64(index3),float64(index5))))
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
