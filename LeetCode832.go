package main

import "fmt"

func flipAndInvertImage(A [][]int) [][]int {
	for _, row := range A {
		left, right := 0, len(row)-1
		for left < right {
			if row[left] == row[right] {
				row[left] ^= 1
				row[right] ^= 1
			}
			left++
			right--
		}
		if left == right {
			row[left] ^= 1
		}
	}
	return A
}

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1,1,0},{1,0,1},{0,0,0}}))
}
