package main

import (
	"fmt"
	"strings"
)

func isSubsequence(s string, t string) bool {
	index := -1
	for i := 0 ; i < len(s) ; i ++  {
		index = strings.IndexByte(t, s[i])
		if index == -1 || t == ""{
			return false
		}
		t = t[index + 1:]
	}
	return true
}
func main() {
	fmt.Println(isSubsequence("axc","ahbgdc"))
}