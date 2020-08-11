package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	hold1 , hold2 := math.MinInt64 , math.MinInt64
	release1 , release2 := 0 ,0
	max := func( a , b int) int {
		if a > b {
			return a
		}
		return b
	}
	for _ , v := range prices {
		release2 = max(release2 , hold2 + v )
		hold2    = max(hold2    , release1 - v)
		release1 = max(release1 , hold1 + v)
		hold1    = max(hold1    , -v)
	}
	return release2
}
func main() {
	fmt.Println(maxProfit([]int{3,3,5,0,0,3,1,4}))
}
