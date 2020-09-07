package main

import (
	"fmt"
	"strconv"
)

func getPermutation(n int, k int) string {

	nums := []int{}
	for i := 1 ; i <= n ; i ++ {
		nums = append(nums,i)
	}

	fact := []int{}
	fact = append(fact,1)
	for i := 1 ; i < n ; i ++ {
		fact = append(fact,fact[i - 1] * i)
	}

	k = k - 1
	str := ""
	for i := n - 1 ; i >= 0 ; i -- {
		num := k / fact[i]
		str += strconv.Itoa(nums[num])
		nums = append(nums[:num], nums[num+1:]...)
		k = k % fact[i]
	}
	return str
}
func main() {
	fmt.Println(getPermutation(4,9))
}
