package main

import "fmt"

func maxProfit(prices []int) int {
	maxprices := 0
	minprices := prices[0]
	max := func(a , b int) int {
		if a > b {
			return a
		}
		return b
	}
	for i := 1 ; i < len(prices) ; i ++ {
		if prices[i] <  minprices {
			minprices = prices[i]
		} else {
			maxprices = max(prices[i] - minprices,maxprices)
		}
	}
	return maxprices
}
func main() {
	fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
