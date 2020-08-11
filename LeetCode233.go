package main

import (
	"fmt"
	"math"
)

func countDigitOne(n int) int {
	sum , i , num := 0 , 1 , n
	for num > 0 {
		if num % 10 == 0 {
			sum = sum + (num / 10) * i
		}
		if num % 10 == 1 {
			sum = sum + (num / 10) * i + (n % i) + 1
		}
		if num % 10 > 1 {
			sum = sum + int(math.Ceil( float64(num) / 10.0 )) * i
		}
		num = num / 10
		i = i * 10
	}
	return sum
}
func main() {
	fmt.Println(countDigitOne(13))
}
