package main

import "fmt"

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0 ;
	}
	x , y , sum := len(grid) , len(grid[0]) , 0

	for i := 0 ; i < x ; i ++ {
		for j := 0 ; j < y ; j ++ {
			if grid[i][j] == '1' {
				sum = sum + 1
				DFS( grid,i,j)
			}
		}
	}
	return sum ;
}
func DFS (grid [][]byte,xi int, yi int) {
	if xi < 0 || yi < 0 || xi >= len(grid) || yi >= len(grid[0]) || grid[xi][yi] == '0'{
		return ;
	}

	grid[xi][yi] = '0';

	DFS(grid, xi , yi - 1)
	DFS(grid, xi , yi + 1)
	DFS(grid, xi - 1, yi )
	DFS(grid, xi + 1, yi )
}
func main() {
	fmt.Println(numIslands([][]byte{{'1','1','1','1','0'},{'1','1','0','1','0'},{'1','1','0','0','0'},{'0','0','0','0','0'}}))
}
