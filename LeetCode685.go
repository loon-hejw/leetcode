package main

func findRedundantDirectedConnection(edges [][]int) []int {

	var find func(parent []int , j int) int
	find = func(parent []int, j int) int {
		for j != parent[j] {
			parent[j] = parent[parent[j]]
			j = parent[j]
		}
		return j
	}

	can1 := []int{-1 , -1}
	can2 := []int{-1 , -1}
	parent := make([]int,len(edges) + 1 )
	for i := 0 ; i < len(edges) ; i ++ {
		if parent[edges[i][1]] == 0 {
			parent[edges[i][1]] = edges[i][0]
		} else {
			can2 = []int{edges[i][0],edges[i][1]}
			can1 = []int{parent[edges[i][1]],edges[i][1]}
			edges[i][1] = 0
		}
	}

	for i := 0 ; i < len(edges) ; i++ {
		parent[i] = i
	}

	for i := 0 ; i < len(edges) ; i ++ {
		if edges[i][1] == 0 {
			continue
		}
		child  := edges[i][1]
		father := edges[i][0]
		if find(parent ,father) == child {
			if can1[0] == -1 {
				return edges[i]
			}
			return can1
		}
		parent[child] = father
	}
	return can2
}
func main() {
	
}
