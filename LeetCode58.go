package main

import "fmt"

func lengthOfLastWord(s string) int {
	sum := 0
	str := []byte{}
	str = append(str , s...)
	for str[len(str) - 1] == ' ' {
		str = str[:len(str) - 1]
	}

	for _ ,v := range str {
		sum ++
		if v == ' ' {
			sum = 0
		}
	}
	return sum
}
func main() {
	fmt.Println(lengthOfLastWord("Hello World    "))
}
