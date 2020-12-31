package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {

	n := len(intervals)
	if n == 0 {
		return n
	}
	sort.Slice(intervals , func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans , right := 1 , intervals[0][1]
	for _ , v := range intervals[1:] {
		if v[0] >= right {
			ans ++
			right = v[1]
		}
	}
	return n - ans
}

func main() {
	
}
