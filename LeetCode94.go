package main

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var dfs func(tree *TreeNode)
	dfs = func(tree *TreeNode) {
		if tree.Left != nil {
			dfs(tree.Left)
		}
		res = append(res,tree.Val)
		if tree.Right != nil {
			dfs(tree.Right)
		}
	}
	if root != nil {
		dfs(root)
	}
	return res
}
func main() {
	
}
