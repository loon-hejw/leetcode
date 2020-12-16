package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {

	pMap  := map[byte]string{}
	sMap  := map[string]byte{}
	words := strings.Split(s , " ")
	index := 0
	if len(words) != len(pattern) {
		return false
	}
	for _ , v := range words {
		if pMap[pattern[index]] == "" && sMap[v] == 0 {
			pMap[pattern[index]] = v
			sMap[v] = pattern[index]
			index ++
		} else if pMap[pattern[index]] != v || sMap[v] != pattern[index] {
			return false
		} else {
			index ++
		}
	}
	return true
}

func main() {
	fmt.Println(wordPattern("abba" ,"dog BBB BBB dog"))
}
