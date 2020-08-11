package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	atoiType , sum ,left  := 1 , 0 , 0
	isNumber := func (b byte) bool {
		if b >= '0' && b <= '9' {
			return true
		}
		return false
	}
	// 去除空格
	for left < len(str) && str[left] == ' ' { left ++ }

	if left < len(str) && str[left] == '-' {
		atoiType = -1 ; left ++
	} else if left < len(str) && str[left] == '+' {
		left ++
	} else if left < len(str) && !isNumber(str[left]) {
		return 0
	}
	for left < len(str) && isNumber(str[left]) {
		sum = sum * 10 + int( str[left] - '0' )
		if sum * atoiType <= math.MinInt32 {
			return math.MinInt32
		}
		if sum * atoiType >= math.MaxInt32{
			return math.MaxInt32
		}
		left ++
	}
	return sum * atoiType
}
func main() {
	fmt.Println(myAtoi("-91283472332"))
}
