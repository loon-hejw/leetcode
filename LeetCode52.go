package main

import "fmt"

func totalNQueens(n int) int {
	res :=0
	var dfs func(int,int,int,int,int)
	dfs = func(n,col,row,ld,rd int) {
		if row >= n {
			res ++
			return
		}
		// 将所有能放置queen的位置改为1，方便后续遍历
		bits := ^(col|ld|rd)&(1<<n-1)

		for bits >0 {
			pick := bits & -bits
			dfs(n,col|pick,row+1,(ld|pick)<<1,(rd|pick)>>1)
			bits &= bits-1
		}
	}
	dfs(n,0,0,0,0)
	return res
}
func main() {
	fmt.Println(totalNQueens(4))
}
