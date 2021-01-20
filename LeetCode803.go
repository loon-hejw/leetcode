package main

func hitBricks(grid [][]int, hits [][]int) []int {

	for _, hit := range hits {
		grid[hit[0]][hit[1]]--
	}

	remain := make(map[int]bool)
	for i := 0; i < len(grid[0]); i++ {
		if 1 == grid[0][i] {
			findRemain(grid, 0, i, remain)
		}
	}

	ans := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		grid[hits[i][0]][hits[i][1]]++
		if 1 != grid[hits[i][0]][hits[i][1]] {
			continue
		}
		c := len(remain)
		if hits[i][0] == 0 || (hits[i][0] > 0 && contains(remain, (hits[i][0]-1)*len(grid[0])+hits[i][1])) ||
			(hits[i][0] < len(grid)-1 && contains(remain, (hits[i][0]+1)*len(grid[0])+hits[i][1])) ||
			(hits[i][1] > 0 && contains(remain, hits[i][0]*len(grid[0])+hits[i][1]-1)) ||
			(hits[i][1] < len(grid[0])-1 && contains(remain, hits[i][0]*len(grid[0])+hits[i][1]+1)) {
			findRemain(grid, hits[i][0], hits[i][1], remain)
			ans[i] = len(remain) - c - 1
		}
	}
	return ans
}

func contains(table map[int]bool, k int) bool {
	_, ok := table[k]
	return ok
}

func findRemain(grid [][]int, row, col int, remain map[int]bool) {
	if 1 != grid[row][col] || contains(remain, row*len(grid[0])+col) {
		return
	}
	remain[row*len(grid[0])+col] = true
	if row > 0 {
		findRemain(grid, row-1, col, remain)
	}
	if row < len(grid)-1 {
		findRemain(grid, row+1, col, remain)
	}
	if col > 0 {
		findRemain(grid, row, col-1, remain)
	}
	if col < len(grid[0])-1 {
		findRemain(grid, row, col+1, remain)
	}
}


func main() {
	
}
