package main

func search2(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	// 对比33题 增加这两行代码即可解决
	for left < right && nums[left] == nums[left + 1] { left ++ }
	for left < right && nums[right] == nums[right - 1] { right -- }

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			return true
		}

		// 判断是否在前半部分查找
		if (nums[left] <= target && target <= nums[mid]) || (nums[mid] <= nums[right] && (target < nums[mid] || target > nums[right]))  {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
func main() {
	
}
