package main

import (
	"fmt"
)

func xorOperation(n int, start int) int {

	ans := start
	for i := 1 ; i < n ; i ++ {
		ans = ans ^ ( 2 * i + start )
	}
	return ans
}

func main() {
	fmt.Println(xorOperation(10 , 5 ))
}
