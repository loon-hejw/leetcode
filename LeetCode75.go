package main

func sortColors(nums []int)  {

	left , right := 0 , len(nums) - 1
	for i := 0 ; i < len(nums) && i <= right ; {
		if nums[i] == 0 && i >= left {
			nums[i] , nums[left] = nums[left] , nums[i]
			left ++
		} else if nums[i] == 2 && i <= right {
			nums[i] , nums[right] = nums[right] , nums[i]
			right --
		} else {
			i ++
		}
	}
}

func main() {
	
}
