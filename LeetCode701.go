package main

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	var dfs func(root *TreeNode) *TreeNode
	dfs = func(root *TreeNode) *TreeNode {
		if root == nil {
			return &TreeNode{
				Val:val,
				Left: nil,
				Right: nil,
			};
		}
		if root.Val < val {
			root.Right = dfs(root.Right)
		}
		if root.Val > val {
			root.Left = dfs(root.Left)
		}
		return root
	}
	return dfs(root)
}
func main() {
	
}
