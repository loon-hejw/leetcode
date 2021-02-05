package main

import "fmt"

func pivotIndex(nums []int) int {

	total := 0
	for _ , v := range nums {
		total += v
	}
	sum := 0
	for k , v := range nums {
		if ( v + sum  << 1 == total ) {
			return k
		}
		sum += v
	}
	return -1
}
func main() {
	fmt.Println(pivotIndex([]int{-1,-1,-1,-1,-1,0}))
}
