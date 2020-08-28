package main

import "fmt"

func judgeCircle(moves string) bool {
	cross , vertical  := 0 , 0
	for _ , v := range moves {
		if v == 'L' {
			cross ++
		} else if v == 'R' {
			cross --
		} else if v == 'U' {
			vertical ++
		} else if v == 'D' {
			vertical --
		}
	}
	return cross == 0 && vertical == 0
}
func main() {
	fmt.Println(judgeCircle("UD"));
}
