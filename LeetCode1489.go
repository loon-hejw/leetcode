package main

import (
	"math"
	"sort"
)

type unionFind struct {
	parent   []int
	size     []int
	setCount int
}

func newUnionFind(n int) *unionFind {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind{parent: parent, size: size, setCount: n}
}

func (uf *unionFind) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind) union(x, y int) bool {
	fx, fy := uf.find(x), uf.find(y)
	if fx == fy {
		return false
	}
	if uf.size[fx] < uf.size[fy] {
		fx, fy = fy, fx
	}
	uf.size[fx] += uf.size[fy]
	uf.parent[fy] = fx
	uf.setCount--
	return true
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {

	for i , e := range edges {
		edges[i] = append(e , i)
	}
	sort.Slice(edges , func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	calcMst := func(uf *unionFind , ignoreID int) int {
		mstValue := 0
		for i , e := range edges {
			if i != ignoreID && uf.union(e[0] , e[1]) {
				mstValue += e[2]
			}
		}
		if uf.setCount > 1 {
			return math.MaxInt64
		}
		return mstValue
	}

	mstValue := calcMst(newUnionFind(n) , -1)
	var keyEdges , pseudokeyEdges []int
	for i , e := range edges {
		if calcMst(newUnionFind(n) , i) > mstValue {
			keyEdges = append(keyEdges , e[3])
			continue
		}

		uf := newUnionFind(n)
		uf.union(e[0] , e[1])
		if e[2] + calcMst(uf , i ) == mstValue {
			pseudokeyEdges = append(pseudokeyEdges , e[3])
		}
	}
	return [][]int{keyEdges , pseudokeyEdges}
}

func main() {

}
