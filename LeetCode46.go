package main

func permute(nums []int) [][]int {
	list := [][]int{}
	return backtrack(nums,list,0,len(nums))
}

func backtrack (nums []int , list [][]int , start int , end int) [][]int {

	if start == end {
		temp := make([]int,len(nums))
		copy(temp , nums)
		list = append(list , temp)
		return list
	}

	for i := start ; i < end ; i ++ {

		nums[i] , nums[start] = nums[start] , nums[i]

		list = backtrack( nums, list, start + 1 , end )

		nums[i] , nums[start] = nums[start] , nums[i]
	}

	return list
}
func main() {
	permute([]int{1,2,3})
}
