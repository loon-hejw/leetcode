package main

import (
	"fmt"
)
func maxScoreSightseeingPair(A []int) int {
	if len(A) <= 1 { return 0 }
	if len(A) == 2 { return A[0] + A[1] - 1}
	ans,max :=  0 ,A[0] + 0
	for i := 1 ; i < len(A)  ; i++ {
		if ans < max + A[i] - i {
			ans = max + A[i] - i
		}
		if max < A[i] + i {
			max = A[i] + i
		}
	}
	return ans
}
func main() {
	fmt.Println(maxScoreSightseeingPair([]int{2,2,2}))
}
