package main

import "fmt"

func convertToTitle(columnNumber int) string {

	if columnNumber <= 0 {
		return ""
	}

	var s []byte
	for columnNumber > 0 {
		columnNumber --
		s = append(s ,  byte(65 + columnNumber % 26))
		columnNumber /= 26
	}

	for i , j := 0 , len(s) - 1; j > i ; i ++  {
		s[i] , s[j - i] = s[j - i] , s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(convertToTitle(701))
}
