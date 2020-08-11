package main

import "fmt"

func myPow(x float64, n int) float64 {

	if n <= 0 {
		return 1.0 / quickMul(x, -n)
	} else {
		return quickMul(x, n)
	}
}

func quickMul(x float64, n int) float64 {

	if n == 0 {
		return 1
	}
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}

func main() {
	fmt.Println(myPow(4,4))
}
