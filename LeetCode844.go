package main

import "fmt"

func backspaceCompare(S string, T string) bool {

	var sList []byte
	var tList []byte
	for i := 0 ; i < len(S) ; i++ {
		if S[i] != '#' {
			sList = append(sList , S[i])
		} else if len(sList) > 0 {
			sList = sList[0:len(sList) - 1]
		}
	}
	for j := 0 ; j < len(T) ; j ++ {
		if T[j] != '#' {
			tList = append(tList , T[j])
		} else if len(tList) > 0 {
			tList = tList[0:len(tList) - 1]
		}
	}
	fmt.Println(string(sList))
	fmt.Println(string(tList))
	return string(sList) == string(tList)
}
func main() {
	fmt.Println(backspaceCompare("abc#","bac#"))
}
