package main

import (
	"fmt"
	"strings"
)

var neighbors = [6][]int{{1, 3}, {0, 2, 4}, {1, 5}, {0, 4}, {1, 3, 5}, {2, 4}}

type PuzzleSet struct {
	Map map[string]bool
	Set []string
}

func newPuzzleSet() PuzzleSet {
	return PuzzleSet{
		Set: []string{},
		Map: map[string]bool{},
	}
}

func slidingPuzzle(board [][]int) int {
	const target = "123450"
	s := make([]byte, 0, 6)
	for _, r := range board {
		for _, v := range r {
			s = append(s, '0'+byte(v))
		}
	}

	if string(s) == target {
		return 0
	}

	get := func(status string) (ret []string) {
		s := []byte(status)
		x := strings.Index(status, "0")
		for _, y := range neighbors[x] {
			s[x], s[y] = s[y], s[x]
			ret = append(ret, string(s))
			s[x], s[y] = s[y], s[x]
		}
		return
	}

	count := 1
	end := newPuzzleSet()
	start := newPuzzleSet()
	end.Set = append( end.Set , target)
	start.Set =  append(start.Set , string(s))
	end.Map[target] = true
	start.Map[string(s)] = true

	isRepeat := map[string]bool{}

	for len(end.Set) > 0 && len(start.Set) > 0 {

		if len(start.Set) > len(end.Set) {
			temp := start
			start = end
			end = temp
		}

		temp := newPuzzleSet()
		for _ , s := range start.Set {
			for _ , next := range get(s) {
				if end.Map[next] {
					return count
				}
				if isRepeat[next] || temp.Map[next]{
					continue
				}
				temp.Set = append(temp.Set , next)
				temp.Map[next] = true
				isRepeat[next] = true
			}
		}
		count ++
		start = temp
	}
	return -1
}


func main() {
	fmt.Println(slidingPuzzle([][]int{{4 ,1,2} , {5,0 , 3}}))
}
