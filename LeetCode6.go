package main

import "fmt"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	min := numRows
	if min > len(s) {
		min = len(s)
	}
	ans := make([]string ,min )
	curRow := 0
	goingDown := false
	for i := range s {
		ans[curRow] += string(s[i])
		if curRow == 0 || curRow == numRows - 1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow += 1
		} else {
			curRow += -1
		}
	}

	s1 := ""
	for i := range ans {
		s1 += ans[i]
	}
	return s1
}

func main() {
	fmt.Println(convert("PAYPALISHIRING",5))
}
