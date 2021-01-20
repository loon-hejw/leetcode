package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {

	emailIndex := map[string]int{}
	emailName  := map[string]string{}

	for _ , v := range accounts {
		name := v[0]
		for i := 1 ; i < len(v) ; i ++ {
			if _, has := emailIndex[v[i]]; !has {
				emailIndex[v[i]] = len(emailIndex)
				emailName[v[i]] = name
			}
		}
	}

	parent := make([]int , len(emailIndex))
	for i := range parent {
		parent[i] = i
	}

	var find func( int ) int
	find = func (x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func( from , to int ) {
		parent[find(from)] = find(to)
	}

	for _ , v := range accounts {
		firstIndex := emailIndex[v[1]]
		for _ , email := range v[2:] {
			union(emailIndex[email] , firstIndex)
		}
	}

	indexToEmails := map[int][]string{}
	for email , index := range emailIndex {
		index = find(index)
		indexToEmails[index] = append(indexToEmails[index] , email)
	}
	ans :=  [][]string{}
	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return ans
}

func main() {
	
}
