package main

import "fmt"

func predictPartyVictory(senate string) string {

	r , d := []int{} , []int{}
	for k , v := range senate {
		if v == 'R' {
			r = append(r, k)
		} else {
			d = append(d , k)
		}
	}
	// 分别循环两个数组
	for len(r) > 0 && len(d) > 0 {
		if r[0] < d[0] {
			r = append(r , r[0] + len(senate))
		} else {
			d = append(d , d[0] + len(senate))
		}
		r = r[1:]
		d = d[1:]
	}
	if len(r) > len(d) {
		return "Radiant"
	}
	return "Dire"
}

func main() {
	fmt.Println(predictPartyVictory("RDD"))
}
