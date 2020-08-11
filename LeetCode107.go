package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var res [][]int
	var cont []*TreeNode
	if root == nil {
		return [][]int{}
	}
	cont = append(cont , root)
	for len(cont) > 0 {
		contTemp := []*TreeNode{}
		size := len(cont)
		resTemp := []int{}
		for size > 0 {
			temp := cont[0]
			cont = cont[1:]
			resTemp = append(resTemp,temp.Val)
			if temp.Left != nil {
				contTemp = append(contTemp , temp.Left)
			}
			if temp.Right != nil {
				contTemp = append(contTemp , temp.Right)
			}
			size --
		}
		res = append(res,resTemp)
		cont = contTemp
	}

	left , right := 0 , len(res) - 1
	for left < right {
		res[left] , res[right] = res[right], res[left]
		left ++ ; right --
	}
	return res
}
func main() {
	
}
