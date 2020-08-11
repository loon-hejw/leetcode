package main

import "sort"

func intersect(nums1 []int, nums2 []int) []int {

	sort.Ints(nums1)
	sort.Ints(nums2)

	nums1_start , nums2_start := 0 , 0
	sum := []int{}
	for nums1_start < len(nums1) && nums2_start < len(nums2) {

		if nums1[nums1_start] == nums2[nums2_start] {
			sum = append(sum , nums2[nums2_start])
			nums1_start ++
			nums2_start ++
		} else if nums1[nums1_start] > nums2[nums2_start] {
			nums2_start ++
		} else {
			nums1_start ++
		}
	}
	return sum
}

func main() {
	
}
