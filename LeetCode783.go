package main

import "math"

func minDiffInBST(root *TreeNode) int {

	ans , pre := math.MaxInt32 , 0
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pre != 0 && ans > root.Val - pre {
			ans = root.Val - pre
		}
		pre = root.Val
		dfs(root.Right)
	}
	dfs(root)
	return ans
}


func main() {
	
}
