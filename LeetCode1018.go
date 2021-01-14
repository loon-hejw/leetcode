package main

import "fmt"

func prefixesDivBy5(A []int) []bool {

	ans := make([]bool , len(A))
	sum := 0
	for k , v := range A {
		sum = ( sum << 1 | v ) % 5
		ans[k] = sum == 0
	}
	return ans
}

func main() {
	fmt.Println(prefixesDivBy5([]int{1,1,1,1}))
}
