package main

import "fmt"

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	m , n := len(haystack) , len(needle)
	kmp := make([]int , n)

	for i , j := 0 , 1 ; j < len(needle) ; j ++ {
		if i > 0 && needle[i] != needle[j] {
			i = kmp[i - 1]
		}
		if needle[i] == needle[j] {
			i ++
		}
		kmp[j] = i
	}

	for i , j := 0 , 0 ; i < m ; i ++ {

		for j > 0 && haystack[i] != needle[j] {
			j = kmp[j - 1]
		}

		if haystack[i] == needle[j] {
			j ++
		}

		if j == n {
			return i - n + 1
		}
	}
	return -1
}

func test() {
	fmt.Println("test")
}

func main() {
	fn := test
	fn()
	fmt.Println(strStr("mississipi","issipi"))
}