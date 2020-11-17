package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people , func(i, j int) bool {
		a ,b  := people[i] , people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	ans := [][]int{}
	for _ , v := range people {
		index := v[1]
		ans = append(ans[:index] , append([][]int{ v }, ans[index:]...) ... )
	}
	return ans
}

func main() {
	
}
