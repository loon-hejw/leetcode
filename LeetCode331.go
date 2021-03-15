package main

import "fmt"

func isValidSerialization(preorder string) bool {

	c := 1
	n := len(preorder)
	for i := 0 ; i < n ; {
		if c == 0 {
			return false
		}
		if preorder[i] == ',' {
			i ++
		} else if preorder[i] == '#' {
			c --
			i ++
		} else {
			for i < n  && preorder[i] != ',' {
				i ++
			}
			c ++
		}
	}
	return c == 0
}

func main() {
	fmt.Println(isValidSerialization("9,#,#,1"))
}