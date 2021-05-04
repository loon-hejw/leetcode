package main

func rangeSumBST(root *TreeNode, low int, high int) int {

	ans := 0
	var dfs func(root *TreeNode)
	dfs = func (root *TreeNode) {

		if root == nil {
			return
		}

		if root.Val >= low && root.Val <= high {
			ans += root.Val
		}

		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)
	return ans
}

func main() {
	
}
