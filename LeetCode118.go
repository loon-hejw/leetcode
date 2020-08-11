package main

func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	result := make([][]int,numRows)
	result[0] = []int{1}
	for i := 1 ; i < numRows ; i ++ {
		for j := 0 ; j <= i ; j ++ {
			if j == 0 || j == i {
				result[i] = append(result[i],1)
			} else {
				result[i] = append(result[i],result[i - 1][j] + result[i - 1][j - 1])
			}
		}
	}
	return result
}
func main() {
	generate(6)
}
