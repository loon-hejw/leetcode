package main

import (
	"fmt"
)

/**
* Definition for a Node.
* type Node struct {
*     Val int
*     Children []*Node
* }
*/

type Node struct {
	Val int
	Children []*Node
}

var res []int

func preorder(root *Node) []int {

	res = []int{}

	helper(root)

	return res
}

func helper (child *Node) {

	if child == nil {
		return
	}

	res = append(res , child.Val)

	for _ , children := range child.Children {
		helper(children)
	}
}
var levelres [][]int
func main() {
	fmt.Println( len(levelres))
}
