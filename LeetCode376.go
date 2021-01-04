package main

func wiggleMaxLength(nums []int) int {

	if len(nums) < 2 {
		return len(nums)
	}

	ans := 1
	prevDiff := nums[1] - nums[0]
	if prevDiff != 0  {
		ans = 2
	}
	for i := 2 ; i < len(nums) ; i ++ {
		diff := nums[i] - nums[i - 1]
		if diff > 0 && prevDiff <= 0 || diff < 0 && prevDiff >= 0 {
			ans ++
			prevDiff = diff
		}
	}
	return ans
}

func main() {
	
}
