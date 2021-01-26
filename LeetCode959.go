package main

import "fmt"

type union struct {
	parent, size []int
	setCount     int
}

func unionfind(n int) *union {
	parent := make([]int, n)
	size := make([]int, n)

	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &union{parent: parent, size: size, setCount: n}
}

func (u *union) unionfind(x int) int {

	if u.parent[x] != x {
		u.parent[x] = u.unionfind(u.parent[x])
	}
	return u.parent[x]
}

func (u *union) createunion(x, y int) {
	fx, fy := u.unionfind(x), u.unionfind(y)
	if fx == fy {
		return
	}

	if u.size[fx] < u.size[fy] {
		fx, fy = fy, fx
	}
	u.size[fx] += u.size[fy]
	u.parent[fy] = fx
	u.setCount--
}

func regionsBySlashes(grid []string) int {

	n := len(grid)
	u := unionfind(4 * n * n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			idx := i*n + j
			if i < n-1 {
				bottom := idx + n
				u.createunion(idx*4+2, bottom*4)
			}
			if j < n-1 {
				right := idx + 1
				u.createunion(idx*4+1, right*4+3)
			}
			if grid[i][j] == '/' {
				u.createunion(idx*4, idx*4+3)
				u.createunion(idx*4+1, idx*4+2)
			} else if grid[i][j] == '\\' {
				u.createunion(idx*4, idx*4+1)
				u.createunion(idx*4+2, idx*4+3)
			} else {
				u.createunion(idx*4, idx*4+1)
				u.createunion(idx*4+1, idx*4+2)
				u.createunion(idx*4+2, idx*4+3)
			}
		}
	}
	return u.setCount
}

func main() {
	fmt.Println(regionsBySlashes([]string{"\\/","/\\"}))
}
