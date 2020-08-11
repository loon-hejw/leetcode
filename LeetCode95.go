package main

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return rateTrees(1, n)
}
func rateTrees(start int , end int ) []*TreeNode {
	TreeNodes := []*TreeNode{}
	if start > end {
		TreeNodes = append(TreeNodes, nil)
		return TreeNodes
	}
	for i := start ; i <= end ; i ++ {
		LeftTree  := rateTrees(start , i - 1)
		RightTree := rateTrees(i + 1 , end)
		for _ , lv := range LeftTree {
			for _ , rv := range RightTree {
				t := &TreeNode{}
				t.Val   = i
				t.Left  = lv
				t.Right = rv
				TreeNodes = append(TreeNodes, t)
			}
		}
	}
	return TreeNodes
}
func main() {
	
}
