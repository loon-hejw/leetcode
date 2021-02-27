package main

import "fmt"

func transpose(matrix [][]int) [][]int {

	n , m := len(matrix) , len(matrix[0])
	ans := make([][]int , m)
	for i := 0 ; i < m ; i ++ {
		ans[i] = make([]int,n)
	}

	for i := 0 ; i < n ; i ++ {
		for j := 0 ; j < m ; j ++ {
			ans[j][i] = matrix[i][j]
		}
	}

	return ans
}

func main() {
	fmt.Println(transpose([][]int{{1,2,3},{4,5,6}}))
}
