package main

import "fmt"

func calculate(s string) int {

	q := []int{1}
	n := len(s)
	sign := 1
	ans := 0
	for i := 0 ; i < n ; {
		switch s[i] {
		case ' ':
			i++
		case '+':
			sign = q[len(q) - 1]
			i++
		case '-':
			sign = -q[len(q) - 1]
			i++
		case '(':
			q = append(q , sign)
			i++
		case ')':
			q = q[:len(q) - 1]
			i++
		default:
			num := 0
			for ; i < n && '0' <= s[i] && s[i] <= '9' ; i ++ {
				num = num * 10 + int(s[i]-'0')
			}
			ans += sign * num
		}
	}
	return ans
}

func main() {
	fmt.Println(calculate("1 + 1"))
}