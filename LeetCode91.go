package main

import (
	"fmt"
)

func numDecodings(s string) int {
	for i := 0 ; i < len(s) ; i ++ {
		fmt.Println(s[i:])
	}
	return 0 ;
}

func helper2(s string , n int) int {

	if n == 0 {
		return 0
	}
	fmt.Println(s)
	for i := 0 ; i < n ; i++ {
		c := ""
		copy([]uint8(fmt.Sprint(s[i:])),c)
		helper2(c,len(s) - i)
	}
	return 0
}
func main() {
	helper2("226",3)
}
