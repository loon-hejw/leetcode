package main

import "fmt"

func solveNQueens(n int) [][]string {
	board := [][]string{}
	temp := []string{}
	for i := 0 ; i < n ; i ++ {
		temp = append(temp, ".") ;
	}
	for i := 0 ; i < n ; i ++ {
		board = append(board, temp) ;
	}

	return dfs(board,0,[][]string{})
}

func dfs(board [][]string , index int ,res [][]string) [][] string {

	// 格式化输出
	if index == len(board) {
		temp := []string{}
		for i := 0 ; i < len(board)  ; i ++{
			s := ""
			for j := 0 ; j < len(board[i]) ; j ++ {
				s = s + board[i][j]
			}
			temp = append(temp ,s )
		}
		res = append(res,temp)
		return res
	}

	for j := 0 ; j < len(board) ; j ++ {
		if validate(board,j,index) {
			temp := make([]string,len(board[j]))
			copy(temp,board[j])
			temp[index] = "Q"
			board[j]= temp
			res = dfs(board,index + 1 , res)
			temp[index] = "."
			board[j]= temp
		}
	}

	return res
}

func validate(board [][]string, x int ,y int) bool {

	for i := 0 ; i < len(board) ; i ++ {
		for j := 0 ; j < y ; j ++ {
			if board[i][j] == "Q" && (x + j == y + i || x + y == i + j || x == i) {
				return false
			}
		}
	}
	return true
}
func main() {
	fmt.Println(solveNQueens(4))
}
