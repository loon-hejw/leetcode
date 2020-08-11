package main

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {

	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	x , y , xi , yi := len(matrix[0]) - 1 , len(matrix)  - 1, 0 , 0

	if matrix[0][0] > target || matrix[y][x] < target {
		return false
	}

	for yi <= y {
		mid_y := yi + ( y - yi ) / 2
		if matrix[mid_y][0] == target {
			return true
		} else if matrix[mid_y][0] > target {
			y = mid_y - 1
		} else {
			yi = mid_y + 1
		}
	}

	for xi <= x {
		mid_x := xi + ( x - xi ) / 2
		if matrix[y][mid_x] == target {
			return true
		} else if matrix[y][mid_x] > target {
			x = mid_x - 1
		} else {
			xi = mid_x + 1
		}
	}
	return false ;
}

func main() {

	matrix := [][]int{{1,   3,  5,  7},{10, 11, 16, 20},{23, 30, 34, 50}}
	fmt.Println(searchMatrix(matrix,13))
}
