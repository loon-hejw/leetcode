package main

func merge(nums1 []int, m int, nums2 []int, n int)  {

	start , end , p := m - 1 ,n - 1 , m + n - 1

	for p > -1 {
		if start > -1 && end > -1 {
			if(nums1[start] > nums2[end]) {
				nums1[p] = nums1[start]
				start --
			} else {
				nums1[p] = nums2[end]
				end --
			}
		}else if start == -1 {
			nums1[p] = nums2[end]
			end --
		} else if end == -1 {
			nums1[p] = nums1[start]
			start --
		}
		p --
	}
}
func main() {
	
}
