package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return Abs(height(root.Left) - height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left),height(root.Right)) + 1
}

func max(a , b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs (a  int) int {
	if a > 0 {
		return a
	}
	return -1 * a
}

func main() {
	
}
