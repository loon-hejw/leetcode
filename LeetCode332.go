package main

import (
	"fmt"
	"sort"
)

func findItinerary(tickets [][]string) []string {
	var (
		m = map[string][]string{}
		res []string
	)

	for _ , v := range tickets {
		m[v[0]] = append(m[v[0]] , v[1])
	}

	for k := range m {
		sort.Strings(m[k])
	}

	var dfs func(star string)
	dfs = func(star string) {
		for {
			if  v , ok := m[star] ; !ok || len(v) == 0 {
				break
			}
			temp := m[star][0]
			m[star] = m[star][1:]
			dfs(temp)
		}
		res = append(res , star)
	}
	dfs("JFK")
	for i := 0 ; i < len(res) >> 1 ; i ++ {
		res[i] , res[len(res) - 1 - i] = res[len(res) - 1 - i] , res[i]
	}
	return res
}

func main() {
	fmt.Println(findItinerary([][]string{{"MUC","LHR"},{"JFK","MUC"},{"SFO","SJC"},{"LHR","SFO"}}))
}
