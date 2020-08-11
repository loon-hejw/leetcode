package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {

	// 字符串大小写
	s = strings.ToLower(s)

	left , right := 0 , len(s) - 1
	for left < right {
		for left < right && !isalnum(s[left]) { left ++ }
		for left < right && !isalnum(s[right]) { right -- }

		if(s[left] != s[right]) {
			return false
		}

		left ++
		right --
	}
	return true
}

func isalnum(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) { return false }
	lenght := [26]int{}
	for k := 0 ; k < len(s) ; k ++ {
		lenght[int(s[k])-97] ++ ;
		lenght[int(t[k])-97] -- ;
	}

	for _ , k := range lenght {
		if k != 0 { return false }
	}
	return true
}
func main() {
	Str := "0123456789AZaz"
	for v , k := range Str{
		fmt.Println(v , k)
	}

	isAnagram2("anagram", "nagaram")
}



