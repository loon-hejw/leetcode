package main

func findMode(root *TreeNode) []int {
	max ,count ,result := 0 , 0 , []int{}
	var res *TreeNode
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if res == nil {
			count = 1
		} else if root.Val == res.Val {
			count ++
		} else {
			count = 1
		}
		if count > max {
			max = count
			result = []int{}
			result = append(result,root.Val)
		}
		if count == max && result[len(result) - 1] != root.Val{
			result = append(result,root.Val)
		}
		res = root
		dfs(root.Right)
	}
	dfs(root)
	return result
}
func main() {
	
}
