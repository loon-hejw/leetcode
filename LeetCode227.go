package main

import "fmt"

func calculate2(s string) int {
	stack := []int{}
	preSign := '+'
	num := 0
	for i , ch := range s {

		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num * 10 + int(s[i] - '0')
		}
		if !isDigit && ch != ' ' || i == len(s) - 1 {
			switch preSign {
			case '+':
				stack = append(stack , num)
			case '-':
				stack = append(stack , -num)
			case '*':
				stack[len(stack) - 1] *= num
			case '/':
				stack[len(stack) - 1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _ , v := range stack {
		num += v
	}
	return num
}

func main() {
	fmt.Println(calculate2(" 3/2 "))
}
