package main

import "fmt"

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	return backsub(nums,[]int{} ,[][]int{} , 0)
}

func backsub (nums []int ,list []int,child [][]int, start int ) [][]int {
	temp := make([]int,len(list))
	copy(temp,list)
	child = append(child, temp)

	for i := start ; i < len(nums) ; i ++ {
		list = append(list,nums[i])
		child = backsub(nums,list,child,i + 1)
		list = list[:len(list) - 1]
	}
	return child
}
func main() {
	fmt.Println(subsets([]int{1,2,3}))
}
