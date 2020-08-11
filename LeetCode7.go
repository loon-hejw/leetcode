package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	res := 0
	for x != 0 {
		res = res * 10 + x % 10
		x   = x / 10
	}
	if res > math.MaxInt32 || res < math.MinInt32 {
		return 0
	}
	return res
}
func helper (x int) int {
	sum := 0
	for  x > 0 {
		sum = (sum + (x % 10)) * 10
		x = x / 10
	}
	sum = sum / 10
	if sum >= 2147483647 {
		sum = 0
	}
	return sum
}

func main() {
	fmt.Println(reverse(120))
}
