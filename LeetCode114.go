package main

func flatten(root *TreeNode)  {
	result := flattenhelp(root,[]*TreeNode{})
	for _ , v := range result {
		root = v
		root.Left = nil
		root = root.Right
	}
}

func flattenhelp( root *TreeNode, result []*TreeNode) []*TreeNode {
	if root != nil {
		result = append(result , root)
		flattenhelp(root.Left , result)
		flattenhelp(root.Right , result)
	}
	return result
}

func main() {
	
}
