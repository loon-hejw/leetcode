package main

import "fmt"

func longestValidParentheses(s string) int {
	max := 0
	stack := []int{}
	stack = append(stack, -1)
	for i := 0 ; i < len(s) ; i ++ {
		if s[i] == '(' {
			stack = append( stack , i )
		} else {
			stack = stack[:len(stack) - 1]
			if len(stack) == 0 {
				stack = append(stack , i)
			} else {
				max = getMax(max , i - stack[len(stack) - 1]);
				stack = stack[:len(stack)]
			}
		}
	}
	return  max
}
func getMax ( a int , b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(longestValidParentheses(")()())"))
}