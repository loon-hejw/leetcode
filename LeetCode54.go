package main

import "fmt"

func spiralOrder(matrix [][]int) []int {

	y1 , y2 := 0 , len(matrix)
	x1 , x2 := 0 , len(matrix[0])

	n := len(matrix) * len(matrix[0])
	ans := []int{}
	for {
		// 第一次循环 上层循环
		for i := x1 ; i < x2 ; i ++ {
			ans = append(ans , matrix[y1][i])
		}

		x2 --
		// 第二次循环 最右循环
		for i := y1 + 1 ; i < y2 ; i ++ {
			ans = append(ans , matrix[i][x2])
		}

		y2 --
		// 第三次循环 下层循环
		for i := x2 - 1 ; i >= x1 ; i -- {
			ans = append(ans , matrix[y2][i])
		}

		y1 ++
		// 第四次循环 最左循环
		for i := y2 - 1; i >= y1 ; i -- {
			ans = append(ans , matrix[i][x1])
		}
		x1 ++

		if len(ans) >= n {
			break
		}
	}
	return ans[:n]
}

func main() {

	fmt.Println(spiralOrder([][]int{{1,2,3,4} , {5,6,7,8} , {9,10,11,12}}))
}