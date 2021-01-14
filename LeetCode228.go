package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {

	n := len(nums)
	ans := []string{}
	i := 1
	for n > 0 {
		star := i - 1
		for i < n && nums[i] == nums[i - 1] + 1 {
			i ++
		}
		end := i - 1
		if nums[star] != nums[end] {
			ans = append(ans , strconv.Itoa(nums[star])+ "->" +strconv.Itoa(nums[end]) )
		} else {
			ans = append(ans ,strconv.Itoa(nums[end]))
		}
		if i >= n {
			break
		}
		i ++
	}
	return ans
}
func main() {
	fmt.Println(summaryRanges([]int{}))
}
