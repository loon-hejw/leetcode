package main

import "fmt"

func findCircleNum(M [][]int) int {
	if M == nil || len(M) == 0 {
		return 0
	}
	count := 0
	visited := make([]int,len(M))
	for i := 0 ; i < len(visited) ; i ++ {
		if visited[i] == 0 {
			Dfs(M,visited,i)
			count ++
		}
	}
	return count
}
func Dfs(M [][]int,visited []int , r  int) {
	for i := 0 ; i < len(M) ; i ++ {
		if M[r][i] == 1 && visited[i] == 0 {
			visited[i] = 1
			Dfs(M,visited,i)
		}
	}
}
func main() {
	m := [][]int{{1,0,0,1},{0,1,1,0},{0,1,1,1},{1,0,1,1}}
	fmt.Println(findCircleNum(m))
}