package main

import (
	"fmt"
)

func removeDuplicateLetters(s string) string {

	numB := [26]int{}
	for i := range s {
		numB[s[i] - 'a'] ++
	}
	ansB  := []byte{}
	boolB := [26]bool{}
	for i := range s {
		if !boolB[s[i] - 'a'] {
			for len(ansB) > 0 && s[i] < ansB[len(ansB) - 1] {
				t := ansB[len(ansB) - 1] - 'a'
				if numB[t] == 0 {
					break
				}
				ansB = ansB[:len(ansB) - 1]
				boolB[t] = false
			}
			ansB = append(ansB , s[i])
			boolB[s[i] - 'a'] = true
		}
		numB[s[i] - 'a'] --
	}
	return string(ansB)
}
func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
}
