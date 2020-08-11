package main

import "fmt"

func trap(height []int) int {
	if len(height) < 3 { return 0 }

	// 初始变量
	ans,l,r := 0 , 0 , len(height) - 1

	// 寻找边界
	for l < r && height[l] <= height[l + 1] { l ++ }
	for l < r && height[r] <= height[r - 1] { r -- }

	left , right := 0,0
	for l < r {
		left = height[l]
		right = height[r]

		if left <= right {
			l = l + 1
			for l < r && left >= height[l] {
				ans = ans + (left - height[l])
				l = l + 1
			}
		} else {
			r = r - 1
			for l < r && right >= height[r] {
				ans = ans + (right - height[r])
				r = r - 1
			}
		}
 	}
	return ans
}
func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
}
