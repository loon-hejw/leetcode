package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {

	queue := []*TreeNode{}
	queue = append(queue,root)
	res := []float64{}
	for len(queue) > 0 {
		sum := 0
		temp := []*TreeNode{}
		for _ , v :=  range queue {
			sum = sum + v.Val
			if v.Left != nil {
				temp = append(temp,v.Left)
			}
			if v.Right != nil {
				temp = append(temp,v.Right)
			}
		}
		res = append(res,float64(sum / len(queue)))
		queue = temp
	}
	return res
}
func main() {
	
}
