package main

import "fmt"

func singleNumber(nums []int) int {
	a , b := 0 , 0
	for _  , v := range nums {
		b = ( v ^ b ) &^ a
		a = ( v ^ a ) &^ b
	}
	return b
}

func main() {
	fmt.Println( ^1 )
	//fmt.Println( singleNumber([]int{2,2,3,2}) )
}
