package main

import "fmt"

func Rotate(matrix [][]int)  {

	if (len(matrix) < 1) {
		return
	}
	length := len(matrix)
	left , right := 0 , length - 1
	for left < right {
		matrix[left] , matrix[right] = matrix[right] , matrix[left]
		left ++
		right --
	}

	for i := 0 ; i < length - 1; i ++ {
		for j := i + 1 ; j < length ; j ++ {
			matrix[i][j] , matrix[j][i] = matrix[j][i] , matrix[i][j]
		}
	}

	PrintlnRoate(matrix)
}

/** 打印参数*/
func PrintlnRoate(matrix [][]int) {
	for i := range matrix{
		fmt.Println(matrix[i])
	}
}

func main() {
	Rotate([][]int{{ 5, 1, 9,11} ,{2, 4, 8,10} , {13, 3, 6, 7} , {15,14,12,16} })
}
