package main

import (
	"fmt"
	"math"
)

func commonChars(A []string) []string {

	minFreq := [26]int{}
	for i := range minFreq {
		minFreq[i] = math.MaxInt32
	}
	for _ , word := range A {
		freq := [26]int{}
		for _ , b := range  word {
			freq[b - 'a'] ++
		}
		for i , f := range freq {
			if f < minFreq[i] {
				minFreq[i] = f
			}
		}
	}
	ans := []string{}
	for i := 0 ; i < 26 ; i ++ {
		for j := 0 ; j < minFreq[i] ; j ++ {
			ans = append(ans , string('a' + i))
		}
	}
	return ans
}

func main() {
	fmt.Println(commonChars([]string{"bella","label","roller"}))
}
