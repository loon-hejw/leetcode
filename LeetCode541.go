package main

import "fmt"

func reverseStr(s string, k int) string {
	var s1 []byte
	s1 = append(s1,s...)
	for i := 0 ; i < len(s1) ; i = i + 2 * k {
		left  := i
		right := left + k - 1
		if  left + k > len(s1) {
			right = len(s1) - 1
		}
		for left < right {
			s1[left] , s1[right] = s1[right] , s1[left]
			left ++ ; right --
		}
	}
	return string(s1)
}
func main() {
	fmt.Println(reverseStr("abcdefg",2))
}
