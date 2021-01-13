package main

import "sort"

type DisjSets struct {
	data    []int
}

func InitDisj(n int) *DisjSets {
	d := new(DisjSets)
	d.data = make([]int, n)
	for i := 0; i < n; i++ {
		d.data[i] = i
	}
	return d
}

func (self *DisjSets) Union(i1, i2 int) {
	f1, f2 := self.Find(i1), self.Find(i2)
	if f1 == f2 {
		return
	}
	self.data[f2] = f1
}

func (self *DisjSets) Find(i int) (st int) {
	if self.data[i] == i {
		return i
	}
	self.data[i] = self.Find(self.data[i])
	return self.data[i]
}

func (self *DisjSets) GroupBy() map[int][]int {
	res := make(map[int][]int)
	for i, v := range self.data {
		res[self.Find(v)] = append(res[self.Find(v)], i)
	}
	return res
}

func smallestStringWithSwaps(s string, pairs [][]int) string {
	ls := len(s)
	newS := []byte(s)
	disj := InitDisj(ls)
	for i := 0; i < len(pairs); i++ {
		disj.Union(pairs[i][0], pairs[i][1])
	}
	m := disj.GroupBy()
	for _, ts := range m {
		t := make([]int, len(ts))
		copy(t, ts)
		sort.Slice(t, func(i, j int) bool {
			return s[t[i]] < s[t[j]]
		})
		for i := 0; i < len(t); i++ {
			newS[ts[i]] = s[t[i]]
		}
	}
	return string(newS)
}


func main() {
	
}
