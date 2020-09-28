package main

type Node struct {
	 Val int
	 Left *Node
	 Right *Node
	 Next *Node
}
func connect(root *Node) *Node {

	if root != nil {
		return root
	}
	// BFS
	var NodeBox []*Node
	NodeBox = append(NodeBox,root)
	for len(NodeBox) > 0 {
		tempBox := NodeBox
		NodeBox = nil
		for k , node := range tempBox {
			if k + 1 < len(tempBox) {
				node.Next = tempBox[k + 1]
			}
			if node.Left != nil {
				NodeBox = append(NodeBox , node.Left)
			}
			if node.Right != nil {
				NodeBox = append(NodeBox , node.Right)
			}
		}
	}
	return root
}
func main() {
	
}
