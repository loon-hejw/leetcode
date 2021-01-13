package main

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {

	// 给方程组中的每个变量编号
	id := map[string]int{}
	for _, eq := range equations {
		a, b := eq[0], eq[1]
		if _, has := id[a]; !has {
			id[a] = len(id)
		}
		if _, has := id[b]; !has {
			id[b] = len(id)
		}
	}

	fa := make([]int, len(id))
	w := make([]float64, len(id))
	for i := range fa {
		fa[i] = i
		w[i] = 1
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			f := find(fa[x])
			w[x] *= w[fa[x]]
			fa[x] = f
		}
		return fa[x]
	}
	merge := func(from, to int, val float64) {
		fFrom, fTo := find(from), find(to)
		w[fFrom] = val * w[to] / w[from]
		fa[fFrom] = fTo
	}

	for i, eq := range equations {
		merge(id[eq[0]], id[eq[1]], values[i])
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasS := id[q[0]]
		end, hasE := id[q[1]]
		if hasS && hasE && find(start) == find(end) {
			ans[i] = w[start] / w[end]
		} else {
			ans[i] = -1
		}
	}
	return ans
}

func main() {
	fmt.Println(calcEquation([][]string{{"a","b"}, {"c","d"}} , []float64{1.0,1.0} ,[][]string{{"a","c"},{"b","d"},{"b","a"},{"d","c"}}))
}
