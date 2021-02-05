package main

import "fmt"

type waterUnion struct {
	parent, size []int
}

func newWaterUnion(n int) *waterUnion {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &waterUnion{parent, size}
}

func (e *waterUnion) waterUnionFind(x int) int {
	if e.parent[x] != x {
		e.parent[x] = e.waterUnionFind(e.parent[x])
	}
	return e.parent[x]
}

func (e *waterUnion) waterUnionCreat(x, y int) {
	fy, fx := e.waterUnionFind(y), e.waterUnionFind(x)
	if fy == fx {
		return
	}
	if e.size[fy] > e.size[fx] {
		fy, fx = fx, fy
	}
	e.size[fx] += e.size[fy]
	e.parent[fy] = fx
	return
}

func (e *waterUnion) isUnionIn(x, y int) bool {
	return e.waterUnionFind(y) == e.waterUnionFind(x)
}

type pair struct {
	x , y int
}

var dirs = []pair{{ -1 , 0 } , { 1 , 0 }, { 0 , -1 },{ 0 , 1 }}

func swimInWater(grid [][]int) int {

	n := len(grid)
	pos := make([]pair , n * n)

	for i , row := range grid {
		for j , h := range row {
			pos[h] = pair{ i , j}
		}
	}

	uf := newWaterUnion(n * n )
	for th := 0 ;; th ++ {
		p := pos[th]
		for _ , d := range dirs {
			if x , y := p.x + d.x  , p.y + d.y ; 0 <= x && x < n && y >= 0 && y < n && grid[x][y] <= th {
				uf.waterUnionCreat(x * n + y , p.x * n + p.y )
			}
		}

		if uf.isUnionIn(0 , n * n - 1) {
			return th
		}
	}
}

func main() {
	fmt.Println(swimInWater([][]int{ {0,1,2,3,4} , {24,23,22,21,5} , {12,13,14,15,16} , {11,17,18,19,20 } , {10,9,8,7,6} }))
}
