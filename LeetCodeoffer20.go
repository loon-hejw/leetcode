package main

import "fmt"

func isNumber(s string) bool {

	if s == "" {
		return false
	}

	left , right := 0 , len(s) - 1
	for left < right && s[left] == ' ' { left ++ }
	for left < right && s[right] == ' ' { right -- }

	// 定义状态 当前数字是否合法  ， e/E 是否存在 ， 小数点是否存在
	isNumber,isEExist ,isDotExist := false ,false ,false
	// 循环遍历
	for i := left ; i <= right ; i ++ {
		c := s[i]
		if c >= '0' && c <='9' {
			isNumber = true
		} else if c == '-' || c == '+' {
			if i != left && s[i - 1] != 'e' && s[i - 1] != 'E' {
				return false
			}
		} else if c == 'E' || c == 'e' {
			if isEExist || !isNumber {
				return false
			}
			isEExist = true
			isNumber = false
		} else if c == '.' {
			if isEExist || isDotExist {
				return false
			}
			isDotExist = true
		} else {
			return false
		}
	}
	return isNumber
}

func main() {
	fmt.Println(isNumber(" -54.53061"))
}
