package main

type EdgesUnion struct {
	Parent, Size []int
	Count        int
}

func NewEdgesUnion(n int) *EdgesUnion {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &EdgesUnion{parent, size, n}
}

func (e *EdgesUnion) EdgeFind(x int) int {

	if e.Parent[x] != x {
		e.Parent[x] = e.EdgeFind(e.Parent[x])
	}
	return e.Parent[x]
}

func (e *EdgesUnion) CreateUnion(x, y int) bool {

	Fx, Fy := e.EdgeFind(x), e.EdgeFind(y)
	if Fx == Fy {
		return false
	}

	if e.Size[Fx] < e.Size[Fy] {
		Fx, Fy = Fy, Fx
	}
	e.Parent[Fy] = Fx
	e.Size[Fx] += e.Size[Fy]
	e.Count--
	return true
}

func (e *EdgesUnion) IsEdgesSet(x, y int) bool {
	return e.EdgeFind(x) == e.EdgeFind(y)
}

func maxNumEdgesToRemove(n int, edges [][]int) int {

	ans := len(edges)
	alice, bob := NewEdgesUnion(n), NewEdgesUnion(n)
	for _, v := range edges {
		x, y := v[1]-1, v[2]-1
		if v[0] == 3 && (!alice.IsEdgesSet(x, y) || !bob.IsEdgesSet(x, y)) {
			alice.CreateUnion(x, y)
			bob.CreateUnion(x, y)
			ans--
		}
	}
	for _, v := range edges {
		x, y := v[1]-1, v[2]-1
		if v[0] == 1 && alice.CreateUnion(x, y) {
			ans--
		}
		if v[0] == 2 && bob.CreateUnion(x, y) {
			ans--
		}
	}
	if alice.Count > 1 || bob.Count > 1 {
		return -1
	}
	return ans
}

func main() {

}
