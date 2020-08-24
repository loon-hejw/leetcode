package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	temp := s + s
	return strings.Index( temp[1 : len(temp) - 1] , s) != -1
}

func main() {
	fmt.Println(repeatedSubstringPattern("aba"))
}
