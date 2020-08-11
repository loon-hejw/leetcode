package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	charMap1 := map[int]int{}
	charMap2 := map[int]int{}
	for i := 0 ; i < len(s) ; i ++ {
		if charMap1[int(s[i])] == 0 && charMap2[int(t[i])] == 0{
			charMap1[int(s[i])] = int(t[i])
			charMap2[int(t[i])] = int(s[i])
		} else if charMap1[int(s[i])] != int(t[i]) || charMap2[int(t[i])] != int(s[i]){
			return false
		}
	}
	return true
}
func main() {
	fmt.Println(isIsomorphic("ab","aa"))
}
