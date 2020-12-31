package main

func maximalRectangle(matrix [][]byte) int {

	if len(matrix) == 0 {
		return 0
	}
	ans := 0
	m , n := len(matrix) , len(matrix[0])

	left := make([][]int, m)

	for i , row := range matrix {
		left[i] = make([]int , n)
		for j , v := range row {
			if v == '0' {
				continue
			}
			if j == 0 {
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j - 1] + 1
			}
		}
	}

	for i , row := range  matrix {
		for j , v := range row {
			if v == '0' {
				continue
			}

			width := left[i][j]
			area := width
			for k := i - 1 ; k >= 0 ; k -- {
				width = min(width , left[k][j])
				area = max(area ,(i - k + 1) * width )
			}
			ans = max(ans ,area)
		}
	}

	return ans
}

func main() {
	
}
