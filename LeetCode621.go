package main

import "fmt"

func leastInterval(tasks []byte, n int) int {

	cnt := map[byte]int{}
	for _, t := range tasks {
		cnt[t]++
	}

	maxExec, maxExecCnt := 0, 0
	for _, c := range cnt {
		if c > maxExec {
			maxExec, maxExecCnt = c, 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	if time := (maxExec-1)*(n+1) + maxExecCnt; time > len(tasks) {
		return time
	}
	return len(tasks)
}

func main() {
	fmt.Println(leastInterval([]byte{'A','A','B','C','D','E','F','G'},2))
}
