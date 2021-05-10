package main

import "fmt"

func decode(encoded []int, first int) []int {
	ans := []int{first}
	for k , v := range encoded {
		ans = append(ans , v ^ ans[k])
	}
	return ans
}

func main() {
	fmt.Println(decode([]int{1,2,3} , 1))
}
