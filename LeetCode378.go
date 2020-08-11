package main

import "fmt"

func kthSmallest(matrix [][]int, k int) int {
	length := len(matrix)
	left , right := matrix[0][0] , matrix[length - 1][length - 1]
	for left < right {
		mid := left + (right - left) / 2
		bCheck := check(matrix,mid, k , length) ;
		if bCheck {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check( matrix [][]int, mid int , k int , length int) bool {
	sum := 0
	vertical ,cross := length - 1 , 0
	for vertical >= 0 && cross < length {
		if matrix[vertical][cross] <= mid {
			sum = sum + vertical + 1
			cross ++
		} else {
			vertical --
		}
	}
	return sum >= k
}

func main() {
	t := [][]int{{1,5,9},{10,11,13},{12,13,15}};
	fmt.Println(kthSmallest(t,7))
}
