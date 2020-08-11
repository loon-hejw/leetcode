package main

import "fmt"

func minSubArrayLen(s int, nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	start , end , sum , ans , temp := 0,0,0,s,0

	for end < len(nums) {
		sum = sum + nums[end]
		temp = temp + nums[end]
		for sum >= s {
			ans = mathMin(ans , end - start + 1)
			sum = sum - nums[start]
			start ++
		}
		end ++
	}

	if temp >= s {
		return ans
	}
	return 0 ;
}

func mathMin( a int ,b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {

	fmt.Println(minSubArrayLen(7,[]int{2,3,1,2,4,3}))

}
