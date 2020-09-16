package main

func solveSudoku(board [][]byte)  {
	var line , column [9][9]bool
	var block [3][3][9]bool
	var spaces [][2]int

	for i , row := range board {
		for j , b := range row {
			if b == '.' {
				spaces = append(spaces,[2]int{i,j})
			} else {
				digit := b - '1'
				// 将 每一行的数字记录 是不是被用过
				line[i][digit] = true
				// 将 每一列的数字记录 是不是被用过
				column[j][digit] = true
				// 将 3*3中矩形中的数字记录 是不是被用过
				block[i/3][j/3][digit] = true
			}
		}
	}

	// 定义递归
	var dfs func(int) bool

	// 递归函数定义
	dfs = func(pos int) bool {

		// 需要遍历的结果都已经遍历过了 代表整个数独解决完成
		if pos == len(spaces) {
			return true
		}

		// 获取对应的数组的 原下标
		i , j := spaces[pos][0] , spaces[pos][1]

		// 从1 ～ 9 遍历
		for digit := byte(0) ; digit < 9 ; digit ++ {

			// 如果 横 纵 单元内都没有值 则进去将对应的状态进行变更
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit]  {

				// 状态码改变
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1' ;

				// 如果全部递归成功则直接返回
				if dfs(pos + 1) {
					return true
				}

				// 递归失败状态 变化
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}
func main() {
	
}
