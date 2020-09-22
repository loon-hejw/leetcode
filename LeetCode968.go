package main

/**
Apply a recusion function dfs.
Return 0 if it's a leaf.
Return 1 if it's a parent of a leaf, with a camera on this node.
Return 2 if it's coverd, without a camera on this node.

For each node,
if it has a child, which is leaf (node 0), then it needs camera.
if it has a child, which is the parent of a leaf (node 1), then it's covered.

If it needs camera, then res++ and we return 1.
If it's covered, we return 2.
Otherwise, we return 0.
*/
func minCameraCover(root *TreeNode) int {
	res := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 2
		}
		left , right := dfs(root.Left) , dfs(root.Right)
		if left == 0 || right == 0 {
			res ++
			return 1
		}
		if left == 1 || right == 1 {
			return 2
		}
		return 0
	}
	if dfs(root) < 1 {
		return res + 1
	}
	return res
}
func main() {
	
}
