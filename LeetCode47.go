package main

import (
	"fmt"
	"sort"
)

func permute(nums []int) [][]int {
	sort.Ints(nums)
	used := make([]bool,len(nums))
	res := [][]int{}
	list := []int{}
	return backtrack(nums,res,list,used)
}

func backtrack (nums []int , res [][]int ,list []int, user []bool) [][]int {

	if len(list) == len(nums) {
		temp := make([]int,len(list))
		copy(temp , nums)
		res = append(res , list)
		return res
	}

	for i := 0 ; i < len(nums) ; i ++ {

		if  user[i] || (i > 0  && nums[i - 1] == nums[i] && !user[i - 1]) {
			continue
		}
		user[i] = true
		list = append(list, nums[i])
		res = backtrack( nums, res ,list ,user )
		user[i] = false

		tmp := make([]int,len(list) - 1)
		copy(tmp,list[:len(list) - 1])
		list = tmp
	}

	return res
}
func main() {
	fmt.Println(permute([]int{2,2,1,1}))
}
