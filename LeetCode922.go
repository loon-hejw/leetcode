package main

import "fmt"

func sortArrayByParityII(A []int) []int {
	j := 1
	for i := 0 ; i < len(A) ; i = i + 2 {
		if A[i] & 1 == 1 {
			for A[j] & 1 == 1 {
				j = j + 2
			}
			A[i] , A[j] = A[j] , A[i]
		}
	}
	return A
}

func main() {
	fmt.Println(sortArrayByParityII([]int{4,2,5,7}))
}
