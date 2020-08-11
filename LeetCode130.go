package main

import "fmt"

func solve(board [][]byte)  {

	for i := 0 ; i < len(board) ; i ++ {
		for j := 0 ; j < len(board[i]) ; j ++ {
			if (i == 0 || j == 0 || i == len(board) - 1 || j == len(board[i]) - 1) && (board[i][j] == 'O') {
				dfs2(&board, i , j )
			}
		}
	}

	for i := 0 ; i < len(board) ; i ++ {
		for j := 0 ; j < len(board[i]) ; j ++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'A' {
				board[i][j] = 'O'
			}
		}
	}
	fmt.Println(board,'A')
}
func dfs2 (board *[][]byte , x , y int)  {
	if  x < 0 || y < 0 || len((*board)) <= x || len((*board)[x]) <= y {
		return
	}
	if (*board)[x][y] == 'A' || (*board)[x][y] == 'X' {
		return
	}
	(*board)[x][y] = 'A'
	dfs2(board , x + 1 , y )
	dfs2(board , x - 1,  y )
	dfs2(board , x , y + 1 )
	dfs2(board , x , y - 1 )
}
func main() {
	//solve([][]byte{{'X','X','X'},{'X','O','X'},{'X','X','X'}})

	//fmt.Println(1011 &^ 1101)
	fmt.Printf("%b",1011 &^ 1101)
}
