package main

import "fmt"

func fairCandySwap(A []int, B []int) []int {
	// A(sum) = sumA
	// B(sum) = sumB
	/**
	 sumA - x + y  = sumB + x - y
	 sumA - sunB = 2x - 2y

	 (sumA - sumB) >> 1 + y = x
	*/
	mapA := map[int]bool{}
	sum  := 0
	for _ , v := range A {
		sum += v
	}
	for _ , v := range B {
		sum -= v
	}
	for _ , v := range B {
		mapA[sum >> 1 + v] = true
	}
	for _ , v := range A {
		if mapA[v] {
			return []int{v , -1 *(sum >> 1 - v) }
		}
	}
	return []int{}
}

func main() {
	fmt.Println(fairCandySwap([]int{1,2,5} , []int{2,4}))
}
