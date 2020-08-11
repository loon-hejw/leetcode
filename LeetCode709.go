package main

import "fmt"

func toLowerCase(str string) string {
	result := make([]byte,len(str))
	for k := range str {
		if str[k] >= 'A' && str[k] < 'Z' {
			result[k] = str[k] + 32
		} else{
			result[k] = str[k]
		}
	}
	return string(result[0:])
}
func main() {
	fmt.Print(toLowerCase("ALLO"))
}
