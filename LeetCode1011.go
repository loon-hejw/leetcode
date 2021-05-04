package main

import "fmt"

func shipWithinDays(weights []int, D int) int {

	left, right := 0, 0
	for _, v := range weights {
		if left < v {
			left = v
		}
		right += v
	}

	for left < right {
		mid := (left + right) >> 1
		need, cur := 1, 0
		for _, v := range weights {
			if cur+v > mid {
				need++
				cur = 0
			}
			cur += v
		}

		if need <= D {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func main() {
	fmt.Println(shipWithinDays([]int{1,2,3,4,5,6,7,8,9,10} , 5))
}
