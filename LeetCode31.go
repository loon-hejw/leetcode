package main

func nextPermutation(nums []int)  {

	n := len(nums) - 1
	index := n - 1
	for index >= 0 && nums[index] >= nums [index + 1] {
		index --
	}

	if index >= 0 {
		for n >= 0 && nums[n] <= nums[index] {
			n --
		}
		nums[index] , nums[n] = nums[n] , nums[index]
	}

	left  := index + 1
	right := len(nums) - 1
	for right > left {
		nums[right] , nums[left] = nums[left] , nums[right]
		right --
		left  ++
	}
}



func main() {
	nextPermutation([]int{3,2,1})
}
