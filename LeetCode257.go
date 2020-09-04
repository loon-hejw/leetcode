package main

import (
	"fmt"
	"strconv"
)

var res []string
func binaryTreePaths(root *TreeNode) []string {
	res = []string{}

	if root == nil {
		return res
	}
	Dfs(root,"")

	return res
}

func Dfs (root *TreeNode,s string) {
	if root == nil {
		return
	}
	T := s + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		res = append( res , T )
		return
	}
	T = T + "->"
	Dfs (root.Left , T)
	Dfs (root.Right , T)
}
func main() {
	fmt.Println("222" + string(1+'0'))
}
