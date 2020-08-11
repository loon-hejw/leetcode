package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {

	indegree := make([]int,numCourses)
	for _ , v := range prerequisites {
		indegree[v[0]] ++
	}
	var queue []int
	var res   []int
	for i := 0 ; i < numCourses ; i ++ {
		if indegree[i] == 0 {
			queue = append(queue , i)
		}
	}

	for len(queue) > 0  {
		temp := queue[0]
		res = append(res , temp)
		queue = queue[1:]
		for _ , v := range prerequisites {
			if v[1] == temp {
				indegree[v[0]] --
				if indegree[v[0]] == 0 {
					queue = append(queue,v[0])
				}
			}
		}
	}
	return res
}
func main() {
	fmt.Println(findOrder(2,[][]int{{1,0}}))
}
