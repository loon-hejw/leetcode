package main

import (
	"fmt"
)

var bArray = []int{
	2,3,5,7,11,13,17,19,
	23,29,31,37,41,43,47,
	53,59,61,67,71,73,79,
	83,89,97,101,
}

func groupAnagrams(strs []string) [][]string {
	product2Index := map[int]int{}
	index := 0
	rets  := [][]string{}
	for _ , v := range strs {
		t := 1
		for i := 0 ; i < len(v) ; i ++ {
			t = t * bArray[int(v[i] - 'a')]
		}
		if k, exist := product2Index[t]; exist {
			rets[k] = append(rets[k], v)
		} else {
			rets = append(rets, []string{v})
			product2Index[t] = index
			index++
		}
	}
	return rets
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat","tea","tan","ate","nat","bat"}))
}
