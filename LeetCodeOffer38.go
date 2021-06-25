package main

import (
	"fmt"
	"sort"
)

func permutation(s string) []string {

	var (
		ans  []string
		data []byte
		perm []byte
	)

	data = append(data, s...)
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	vis := make([]bool , len(data))

	var dfs func(index int)
	dfs = func(index int) {

		if index == len(data) {
			ans = append(ans, string(perm))
			return
		}

		for i , b := range data {
			if vis[i] || i > 0 && !vis[i - 1] && b == data[i - 1] {
				continue
			}

			vis[i] = true
			perm = append(perm , b)
			dfs(index + 1)
			perm = perm[:len(perm) - 1]
			vis[i] = false
		}
	}
	dfs(0)
	return ans
}

func main() {
	fmt.Println(permutation("aac"))
}
