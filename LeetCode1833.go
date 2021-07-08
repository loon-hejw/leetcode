package main

import (
	"fmt"
	"sort"
)

type name struct {
	
}

func maxIceCream(costs []int, coins int) (count int) {

	var b = false
	fmt.Println( !b )
	sort.Ints(costs)
	for _ , v := range costs {
		if coins < v {
			return
		}
		coins = coins - v
		count ++
	}
	return
}

func main() {
	fmt.Println(maxIceCream([]int{1,3,2,4,1} , 7))
}
