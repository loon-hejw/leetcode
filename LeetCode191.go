package main

import "fmt"

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num = num & (num - 1 )
		count ++
	}
	return count
}
func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
