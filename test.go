package main

import (
	"fmt"
	"sort"
)

var unionSet []int

func unionSetFind(x int) int {
	if unionSet[x] == x {
		return x
	}
	unionSet[x] = unionSetFind(unionSet[x])
	return unionSet[x]
}

func unionSetinit(x, y int) {
	unionX, unionY := unionSetFind(x), unionSetFind(y)
	if unionX == unionY {
		return
	}
	unionSet[unionY] = unionX
}

func smallestStringWithSwaps2(s string, pairs [][]int) string {

	unionSet = make([]int, len(s))
	// 将每一个节点都设置成父节点
	for i := 0; i < len(s); i++ {
		unionSet[i] = i
	}
	for _, v := range pairs {
		unionSetinit(v[0], v[1])
	}

	unionGp := map[int][]int{}
	for i, v := range unionSet {
		unionGp[unionSetFind(v)] = append(unionGp[unionSetFind(v)], i)
	}

	newS := []byte(s)
	for _, v := range unionGp {
		t := make([]int, len(v))
		copy(t, v)
		sort.Slice(v, func(i, j int) bool {
			return newS[v[i]] < newS[v[j]]
		})

		for i := 0; i < len(v); i++ {
			newS[t[i]] = s[v[i]]
		}
	}
	return string(newS)
}

func main() {
	fmt.Println(smallestStringWithSwaps2("dcab", [][]int{{0, 3}, {1, 2}, {0, 2}}))
}
