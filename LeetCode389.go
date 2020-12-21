package main

import "fmt"

func findTheDifference(s string, t string) byte {
	var b byte
	for i := range s {
		b ^= s[i] ^ t[i]
	}
	return b ^ t[len(t) - 1]
}

func main() {
	fmt.Println(findTheDifference("abcd", "abcde"))
}
