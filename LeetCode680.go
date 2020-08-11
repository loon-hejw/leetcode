package main

import "fmt"

func validPalindrome(s string) bool {
	left , right , lsum , rsum := 0 , len(s) - 1 , 0 , 0
	lleft , rright := 0 , len(s) - 1
	for left < right {
		for lleft < right && lsum <= 1 && s[lleft] != s[right] { lsum ++ ; lleft ++ }
		for left < rright && rsum <= 1 && s[left] != s[rright] { rsum ++ ; rright -- }
		left ++ ; right -- ; lleft ++ ; rright --
		if rsum > 1 && lsum > 1 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("abca"))
}
