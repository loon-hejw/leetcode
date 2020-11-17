package main

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	max := func(a ,b int) int {
		if a > b {
			return a
		}
		return b
	}
	var dir4 = [][2]int{ { 1, 1 } , { 1 , -1} , { -1 , -1 } , { -1 , 1 }}
	ans := make([][]int , 1 , R * C)
	ans[0] = []int{ r0 , c0}
	maxDist := max(r0 , R-1-r0 ) + max(c0 , C-1-c0)
	row , col := r0 , c0
	for dist := 1 ; dist <= maxDist ; dist ++ {
		row --
		for i , dir := range dir4 {
			for i % 2 == 0 && row != r0 || i % 2 == 1 && col != c0 {
				if 0 <= row && row < R && 0 <= col && col < C {
					ans = append(ans , []int{row , col})
				}
				row += dir[0]
				col += dir[1]
			}
		}
	}
 	return ans
}

func main() {
	
}
