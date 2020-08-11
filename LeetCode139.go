package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {

	dictionary := make(map[string]bool)
	for _ , v := range wordDict {
		dictionary[v] = true
	}

	check := make([]bool,len(s) + 1)
	check[0] = true
	for i := 1 ; i <= len(s) ; i ++ {
		for j := 0 ; j < i ; j ++ {
			if check[j] && dictionary[s[j:i]] {
				check[i] = true
				break
			}
		}
	}

	return check[len(s)]
}
func main() {
	fmt.Println(wordBreak("leetcode",[]string{"leet","code"}))
}
