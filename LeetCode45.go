package main

func jump(nums []int) int {
	jump ,  end , max  := 0 ,0 ,0
	for i := 0 ; i < len(nums) - 1 ; i ++ {
		if i + nums[i] > max {
			max = i + nums[i]
		}
		if i == end {
			end = max
			jump ++
		}
	}
	return jump
}
func main() {
	
}
