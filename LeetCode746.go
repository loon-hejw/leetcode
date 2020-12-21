package main

import "fmt"

func minCostClimbingStairs(cost []int) int {

	Min := func(a ,b int) int {
		if a > b {
			return b
		}
		return a
	}

	p , q := 0 , 0
	for i := 2 ; i <= len(cost) ; i++ {
		p , q = q , Min(q + cost[i - 1] , p + cost[i - 2])
	}
	return q
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10 , 15 , 20}))
}
