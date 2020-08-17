package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

type Node struct {
	Val int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return  node
	}
	nodeMap := map[*Node]*Node{}
	quene   := []*Node{}
	nodeMap[node] = &Node{ node.Val , []*Node{}}
	quene = append(quene,node)
	for len(quene) > 0 {
		temp := quene[0]
		quene = quene[1:]
		for _ , v := range temp.Neighbors {
			if _ , ok := nodeMap[v] ; !ok {
				nodeMap[v] = &Node{ v.Val , []*Node{}}
				quene = append(quene, v)
			}
			nodeMap[temp].Neighbors = append(nodeMap[temp].Neighbors,nodeMap[v])
		}
	}
	return nodeMap[node]
}
func main() {
	
}
