package main

import "fmt"

var child [][]int
func findSubsequences(nums []int) [][]int {
	child = [][]int{}
	findSubsequenceshelper([]int{},0 , nums)
	return child
}

func findSubsequenceshelper(list []int , index int , nums []int) {
	if len(list) > 1 {
		temp := make([]int, len(list))
		copy( temp , list)
		child = append(child , temp)
	}
	hashMap := map[int]bool{}
	for i := index ; i < len(nums) ; i ++ {
		if ok := hashMap[nums[i]] ; ok {
			continue
		}
		if len(list) == 0 || nums[i] >= list[len(list) - 1] {
			hashMap[nums[i]] = true
			list = append(list , nums[i])
			findSubsequenceshelper(list,i + 1 , nums)
			list = list[0:len(list) - 1]
		}
	}
}
func main() {
	fmt.Println(findSubsequences([]int{4,6,7,7}))
}
