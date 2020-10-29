package main

import "fmt"

func numJewelsInStones(J string, S string) int {
	sum := 0
	stoneMap := map[int32]bool{}
	for _ , v := range J {
		stoneMap[v] = true
	}
	for _ ,v := range S {
		if  stoneMap[v] {
			sum = sum + 1
		}
	}
	return sum
}
func main() {
	fmt.Println(numJewelsInStones("aA","aAAbbbb"))
}
