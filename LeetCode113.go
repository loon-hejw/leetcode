package main

func pathSum(root *TreeNode, sum int) [][]int {
	var path []int
	var result [][]int
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		sum = sum - root.Val
		path = append(path , root.Val)
		defer func() {
			sum = sum + root.Val
			path = path[:len(path) - 1]
		}()
		if sum == 0 && root.Left == nil && root.Right == nil {
			temp := make([]int , len(path))
			copy(temp , path)
			result = append(result ,temp)
			return
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return result
}

func main() {
	
}
