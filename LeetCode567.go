package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {

	q1 , q2 := [26]int{} , [26]int{}
	n := len(s1)
	for k , v := range s1 {
		q1[v - 'a'] ++
		q2[s2[k] - 'a'] ++
	}

	if q1 == q2 {
		return true
	}

	for i := n ; i < len(s2) ; i ++ {

		q2[s2[i] - 'a'] ++
		q2[s2[i - n] - 'a'] --

		if q1 == q2 {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(checkInclusion("ab" , "eidbaooo"))
}
