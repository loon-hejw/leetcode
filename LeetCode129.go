package main

func sumNumbers(root *TreeNode) int {
	sum := 0
	var dfs func(root *TreeNode , t int)
	dfs = func(root *TreeNode , t int) {
		if root == nil {
			return
		}
		if root.Right == nil && root.Left == nil {
			sum = sum + t * 10 + root.Val
			return
		}
		dfs(root.Left  ,  t * 10 + root.Val)
		dfs(root.Right ,  t * 10 + root.Val)
	}
	dfs(root, 0)
	return sum
}
func main() {
	
}
