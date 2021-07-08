package main

import (
	"fmt"
	"math"
)

type RouteSet struct {
	Map map[int]bool
	Set []int
}

func newRouteSet() RouteSet  {
	return RouteSet{
		Set: []int{},
		Map: map[int]bool{},
	}
}

func numBusesToDestination(routes [][]int, source int, target int) int {

	if source == target {
		return 0
	}

	set := map[int][]int{}
	site := map[int][]int{}
	count := 0
	start := newRouteSet()
	end := newRouteSet()

	// 判断站上的车是不是通的。
	for i , route := range routes {
		for _ , v := range route {
			set[v] = append(set[v] , i)
			site[i] = append(site[i] , v)
		}
	}

	// 开始从那个车开始找
	for _ , v := range set[source]{
		start.Map[v] = true
		start.Set = append(start.Set , v)
	}

	// 找到那个车结束
	for _ , v := range set[target] {
		end.Map[v] = true
		end.Set = append(end.Set , v)
	}

	// 记录找过的站点
	siteMap := map[int]bool{}
	for len(start.Set) > 0 && len(end.Set) > 0 {
		if len(start.Set) > len(end.Set) {
			temp := start
			start = end
			end = temp
		}
		temp := newRouteSet()
		for _ , v := range start.Set {
			if end.Map[v] {
				return count + 1
			}
			for _ , c := range site[v] {
				if siteMap[c] {
					continue
				}
				for _ , s := range set[c] {
					if !temp.Map[s] {
						temp.Set = append(temp.Set , s)
						temp.Map[s] = true
					}
				}
				siteMap[c] = true
			}
		}
		count ++
		start = temp
	}
	return -1
}


// 建图
func numBusesToDestination2(routes [][]int, source, target int) int {
	if source == target {
		return 0
	}

	n := len(routes)
	edge := make([][]bool, n)
	for i := range edge {
		edge[i] = make([]bool, n)
	}
	rec := map[int][]int{}
	for i, route := range routes {
		for _, site := range route {
			for _, j := range rec[site] {
				edge[i][j] = true
				edge[j][i] = true
			}
			rec[site] = append(rec[site], i)
		}
	}

	dis := make([]int, n)
	for i := range dis {
		dis[i] = -1
	}
	q := []int{}
	for _, bus := range rec[source] {
		dis[bus] = 1
		q = append(q, bus)
	}
	for len(q) > 0 {
		x := q[0]
		q = q[1:]
		for y, b := range edge[x] {
			if b && dis[y] == -1 {
				dis[y] = dis[x] + 1
				q = append(q, y)
			}
		}
	}

	ans := math.MaxInt32
	for _, bus := range rec[target] {
		if dis[bus] != -1 && dis[bus] < ans {
			ans = dis[bus]
		}
	}
	if ans < math.MaxInt32 {
		return ans
	}
	return -1
}

func main() {

	fmt.Println(numBusesToDestination2([][]int{{1 , 2} , { 2 , 3 } , { 3 , 4 } , {4 , 5} , {5 , 6}} ,1 , 2))
}