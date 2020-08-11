package main

import "fmt"

func findMagicIndex(nums []int) int {
	for k , v := range nums {
		if k == v {
			return v
		}
	}
	return -1
}

func main() {
	temp := []int{0, 2, 3, 4, 5}
	fmt.Println(findMagicIndex(temp))
}