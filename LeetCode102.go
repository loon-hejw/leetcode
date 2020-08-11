package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}

	tree := []*TreeNode{root}

	for i := 0 ; len(tree) > 0 ; i ++ {
		res = append(res, []int{})
		temp := []*TreeNode{}
		for j := 0 ; j < len(tree) ; j ++ {
			node := tree[j]
			res[i] = append(res[i], node.Val)

			if node.Left != nil {
				temp = append(temp, node.Left)
			}

			if node.Right != nil {
				temp = append(temp ,node.Right)
			}
		}
		tree = temp
	}

	return res
}

func main() {

}
