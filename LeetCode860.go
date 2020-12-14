package main

import "fmt"

func lemonadeChange(bills []int) bool {
	five  ,ten := 0 , 0
	for _ , v := range bills {
		if v == 5 {
			five ++
		} else if v == 10 {
			five -- ; ten ++
		} else if ten > 0 {
			ten -- ; five --
		} else {
			five = five - 3
		}
		if five < 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(lemonadeChange([]int{5,5,10}))
}
