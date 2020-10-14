package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	ans , pre := math.MaxInt32 , -1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != -1 && node.Val - pre < ans {
			ans = node.Val - pre
		}
		pre = node.Val
		dfs(node.Right)
	}
	dfs(root)
	return ans
}
func main() {
	
}
