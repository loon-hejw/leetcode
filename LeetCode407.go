package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

func trapRainWater(heightMap [][]int) int {

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	visit := make([][]bool, len(heightMap))
	for i := range visit {
		visit[i] = make([]bool, len(heightMap[0]))
	}
	pq := make(PriorityQueue, 0)
	for i := 0; i < len(heightMap); i++ {
		pq = append(pq, &Item{
			value: heightMap[i][0],
			i:     i,
			j:     0,
		})
		visit[i][0] = true
		pq = append(pq, &Item{
			value: heightMap[i][len(heightMap[0])-1],
			i:     i,
			j:     len(heightMap[0]) - 1,
		})
		visit[i][len(heightMap[0])-1] = true
	}
	for i := 1; i < len(heightMap[0])-1; i++ {
		pq = append(pq, &Item{
			value: heightMap[0][i],
			i:     0,
			j:     i,
		})
		visit[0][i] = true
		pq = append(pq, &Item{
			value: heightMap[len(heightMap)-1][i],
			i:     len(heightMap) - 1,
			j:     i,
		})
		visit[len(heightMap)-1][i] = true
	}
	heap.Init(&pq)
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	ret := 0
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		for _, dir := range dirs {
			i, j := item.i+dir[0], item.j+dir[1]
			if i >= 0 && i < len(heightMap) && j >= 0 && j < len(heightMap[0]) && !visit[i][j] {
				ret += max(0, item.value-heightMap[i][j])
				visit[i][j] = true
				heap.Push(&pq, &Item{
					value: max(item.value, heightMap[i][j]),
					i:     i,
					j:     j,
				})
			}
		}
	}
	return ret
}



type Item struct {
	value int
	i     int
	j     int
}

func (i *Item) String() string {
	return strconv.Itoa(i.value)
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}


func main() {
	fmt.Println(trapRainWater([][]int{{12, 13, 1, 12},
		                              {13, 4, 13, 12},
		                              {13, 8, 10, 12},
		                              {12, 13, 12, 12},
		                              {13, 13, 13, 13}}))
}
