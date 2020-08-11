package main

func rotate(nums []int, k int)  {
	nums2 := make([]int,len(nums))
	for i , v := range nums {
		nums2[( i + k) % len(nums)] = v
	}
	for i ,v := range nums2 {
		nums[i] = v
	}
}

func reverse (nums[] int , start int , end int) {
	for start < end {
		nums[start] ,nums[end] =   nums[end] ,nums[start]
		start ++
		end --
	}
}
// 反转数组
func reverserotate(nums []int, k int)  {
	k = k % len(nums)
	reverse(nums , 0 , len(nums) - 1 )
	reverse(nums , 0 , k - 1 )
	reverse(nums ,  k ,len(nums) - 1)
}

func main() {
	rotate([]int{1,2,3,4,5,6,7},3)
}
