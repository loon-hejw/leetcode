package main

import (
	"fmt"
	"math/rand"
)

/*
func kClosest(points [][]int, K int) [][]int {

	sort.Slice(points, func(i, j int) bool {
		p, q := points[i], points[j]
		return p[0]*p[0]+p[1]*p[1] < q[0]*q[0]+q[1]*q[1]
	})
	return points[:K]
}

*/

/*
type pair struct {
	dist  int
	point []int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less (i , j int) bool {
	return h[i].dist > h[j].dist
}

func (h hp) Swap( i , j int)  {
	h[i] , h[j] = h[j] ,h[i]
}

func (h *hp) Push (v interface{}) {
	*h = append(*h , v.(pair))
}

func (h *hp) Pop () interface{} {
	a := *h
	v := a[len(a) - 1]
	*h = a[:len(a) - 1]
	return v
}

func kClosest(points [][]int, K int) [][]int {

	h := make(hp , K)
	for i , p := range points[:K] {
		h[i] = pair{
			p[0] * p[0] + p[1] * p[1],
			p,
		}
	}
	heap.Init(&h)
	for _ , p := range points[K:] {
		if dist := p[0] * p[0] + p[1] * p[1] ; dist < h[0].dist {
			h[0] = pair{ dist, p}
			heap.Fix(&h , 0)
		}
	}
	ans := [][]int{}
	for _ , p := range h {
		ans = append(ans , p.point)
	}
	return ans
}
 */
func kClosest(points [][]int, K int) [][]int {

	rand.Shuffle(len(points) , func(i, j int) {
		points[i] , points[j] = points[j] , points[i]
	})
	
	var less func (i , j []int) bool
	less = func(i, j []int) bool {
		return i[0] * i[0] + i[1] * i[1] < j[0] * j[0] + j[1] * j[1]
	}
	
	var quickSeleck func(left , right int) 
	quickSeleck = func(left, right int) {

		if left == right {
			return
		}

		pivot := points[right]
		lessCount := left
		for i := left ; i < right ; i ++ {
			if less(points[i] , pivot) {
				points[i] , points[lessCount] = points[lessCount] , points[i]
				lessCount ++
			}
		}

		points[right] , points[lessCount] = points[lessCount] , points[right]
		if lessCount + 1 == K {
			return
		} else if lessCount + 1 < K {
			quickSeleck(lessCount + 1 , right)
		} else {
			quickSeleck(left , lessCount - 1)
		}
	}

	quickSeleck(0 , len(points) - 1)
	return points[:K]
}

func main() {
	fmt.Println(kClosest([][]int{{3,3},{5,-1},{-2,4}},2))
}
