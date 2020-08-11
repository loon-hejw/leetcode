package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if board == nil || len(board) == 0 {
		return false
	}
	charword := word[0]
	for k := 0; k < len(board); k++ {
		for n := 0; n < len(board[k]); n++ {
			if charword == board[k][n] {
				if existDfs(board,word,0,k,n) == true {
					return true
				}
			}
		}
	}
	return false
}

func existDfs(board [][]byte, word string, start, c , l int) bool {
	if c < 0 || c >= len(board) || l < 0 || l >= len(board[c]) || word[start] != board[c][l] {
		return false
	}
	if len(word) - 1 == start {
		return true
	}
	tmp := board[c][l]
	board[c][l] = '@'
	bool := existDfs(board,word,start + 1 ,c + 1 , l ) ||  existDfs(board,word,start + 1 ,c - 1 , l ) || existDfs(board,word,start + 1 ,c, l + 1 ) || existDfs(board,word,start + 1 ,c, l - 1 )
	board[c][l] = tmp
	return bool
}
func main() {
	b := [][]byte{{'A','B'},{'C','D'}}
	fmt.Println(exist(b,"ABCD"))
}
