package main

import "fmt"

func connect(root *Node) *Node {

	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		tmp := queue
		queue = nil
		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}

func main() {
	
}
