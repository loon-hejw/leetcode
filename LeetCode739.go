package main

import "fmt"

func dailyTemperatures(T []int) []int {

	if len(T) == 1 { return []int{} }
	temp := make([]int,len(T))
	var stack []int
	for i := 0 ; i < len(T) ; i ++ {
		for len(stack) != 0 && T[stack[len(stack) - 1]] < T[i] {
			temp[stack[len(stack) - 1]] = i - stack[len(stack) - 1]
			stack = stack[0 : len(stack) - 1]
		}
		stack = append(stack, i)
	}
	return temp
}

func main() {
	fmt.Println(dailyTemperatures([]int{73,74,75,71,69,72,76,73}))
}
