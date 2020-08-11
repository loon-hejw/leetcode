package main

import "fmt"

func reverseOnlyLetters(S string) string {
	var s1 []byte
	s1 = append(s1 , S... )
	left , right := 0 , len(s1) - 1
	isLetters := func (b byte) bool {
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z'){
			return true
		}
		return false
	}
	for left < right {
		for left < right && !isLetters(s1[left]) { left ++;}
		for left < right && !isLetters(s1[right]) { right --;}
		s1[left] , s1[right] = s1[right] , s1[left]
		left ++ ; right --
	}
	return string(s1)
}
func main() {
	fmt.Println(reverseOnlyLetters("a-bC-dEf-ghIj"))
}
