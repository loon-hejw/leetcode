package main

func islandPerimeter(grid [][]int) int {
	max := 0
	for i := 0 ; i < len(grid) ; i ++ {
		for j := 0 ; j < len(grid[0]) ; j ++ {
			if grid[i][j] == 1 {
				max = max + 4
				if i > 0 && grid[i - 1][j] == 1 {
					max = max - 2
				}
				if j > 0 && grid[i][j - 1] == 1 {
					max = max - 2
				}
			}
		}
	}
	return max
}
func main() {
	
}
