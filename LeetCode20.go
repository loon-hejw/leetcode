package main

import (
	"fmt"
)

func plusOne(digits []int) []int {

	for v := len(digits) - 1 ; v >= 0 ; v -- {
		digits[v] ++ ;
		digits[v] = digits[v] % 10
		if( digits[v] != 0 ) {
			return digits
		}
	}

	plus := make([]int,len(digits) + 1)
	plus[0] = 1
	return plus
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 { return  0 }
	i := 0;
	for j := 1 ; j<len(nums) ; j ++  {
		if(nums[i] != nums[j]) {
			i = i + 1 ;
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func isValid(s string) bool {
	if len(s) == 0 { return true }
	if len(s) % 2 != 0 { return false }
	subscript := 0
	sub := make([]int,len(s))
	for i := 0 ; i < len(s) ; i++ {
		if s[i] == 40 {
			sub[subscript] = 41
			subscript ++
		} else if s[i] == 91 {
			sub[subscript] = 93
			subscript ++
		} else if s[i] == 123 {
			sub[subscript] = 125
			subscript ++
		} else {
			subscript --
			if subscript == -1 || int(s[i]) != sub[subscript] {
				return false
			}
			sub[subscript] = 0
		}
	}

	return subscript == 0 || (subscript != len(s) && sub[subscript - 1] == 0)
}

func main() {
	fmt.Println(plusOne([]int{9}))
}
