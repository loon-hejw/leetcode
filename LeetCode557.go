package main

import "fmt"

func reverseWords(s string) string {
	var s1 []byte
	s1 = append(s1 , s...)
	left , right := 0 , 0
	for i := 0 ; i < len(s1) ; i ++ {
		if s1[i] == ' ' {
			right = i - 1
			for left < right {
				s1[left] , s1[right] = s1[right] , s1[left]
				left ++ ; right --
			}
			left = i + 1
		}
	}
	if left < len(s1) {
		right = len(s1) - 1
		for left < right {
			s1[left] , s1[right] = s1[right] , s1[left]
			left ++ ; right --
		}
	}
	return string(s1)
}
func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}
