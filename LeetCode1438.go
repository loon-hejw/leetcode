package main

import "fmt"

func longestSubarray(nums []int, limit int) int {

	maxQ , minQ := []int{} , []int{}
	left , ans := 0 , 0
	for k , v := range nums {
		for len(minQ) > 0 && minQ[len(minQ) - 1] > v {
			minQ = minQ[:len(minQ) - 1]
		}
		minQ = append(minQ , v)
		for len(maxQ) > 0 && maxQ[len(maxQ) - 1] < v {
			maxQ = maxQ[:len(maxQ) - 1]
		}
		maxQ = append(maxQ , v)

		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0] - minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left ++
		}
		ans = max(ans ,  k - left  + 1)
	}
	return ans
}


func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestSubarray([]int{10,1,2,4,7,2} , 5))
}
