package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {

	// 递归查找
	var find func (parent []int , f int) int
	find = func(parent []int , f int) int {
		if f != parent[f] {
			parent[f] = find(parent , parent[f])
		}
		return parent[f]
	}

	parent := make([]int,2001)
	for i := 0 ; i < len(parent) ; i ++ {
		parent[i] = i
	}

	for _ ,v := range edges {
		f , t := v[0] , v[1]
		if find(parent,f) == find(parent,t) {
			return v
		} else {
			parent[find(parent,f)] = find(parent,t)
		}
	}
	return []int{0,0}
}


func main() {
	fmt.Println(findRedundantConnection([][]int{{1,2},{1,3},{2,3}}))
}
