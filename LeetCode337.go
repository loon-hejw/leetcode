package main

//noinspection ALL
func rob(root *TreeNode) int {
	return robMax(robDfs(root))
}
func robMax( a ,b int) int {
	if a > b {
		return a
	}
	return b
}
func robDfs(root *TreeNode) (int,int) {
	if root == nil {
		return 0 , 0
	}
	// 比较爷爷 + 孙子 和 两个父级
	rl , lr := robDfs(root.Left)
	rr , ll := robDfs(root.Right)
	return robMax(root.Val + lr + ll, rl + rr ) , rl + rr
}
func main() {

}
