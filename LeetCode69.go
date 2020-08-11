package main

import "fmt"

func mySqrt(x int) int {

	if x == 1 || x == 2 {
		return 1
	}
	left , right , ans := 0 , x  , 0
	for left < right {
		mid  := left + ( right - left ) / 2
		guss := mid * mid
		if guss == x {
			return mid
		}
		if guss > x {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}

func main() {
	fmt.Println(mySqrt(11))
}
