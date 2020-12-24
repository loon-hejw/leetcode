package main
func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	quene := []*TreeNode{}
	quene = append( quene , root)
	ans  := [][]int{}
	for len(quene) > 0 {
		list := []int{}
		q := []*TreeNode{}
		for i := range quene {
			if quene[i].Left != nil {
				q = append(q , quene[i].Left)
			}
			if quene[i].Right != nil {
				q = append(q , quene[i].Right)
			}

			list = append(list , quene[i].Val)
		}
		quene = q
		ans = append(ans  , list)
	}

	for i := 1 ; i < len(ans) ; i = i + 2 {
		left := 0
		right := len(ans[i]) - 1
		for left < right {
			ans[i][left] , ans[i][right] =  ans[i][right] , ans[i][left]
			left ++
			right --
		}
	}

	return ans
}
func main() {
	
}
