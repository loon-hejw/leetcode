package main

import "fmt"

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	m := map[int]int{}
	for i , v := range nums {
		id := getid(v , t + 1)
		if _ , ok := m[id] ; ok {
			return true
		}
		if c , ok := m[id - 1] ; ok && abs(v - c) <= t {
			return true
		}
		if c , ok := m[id + 1] ; ok && abs(v - c) <= t {
			return true
		}
		m[id] = v
		if i >= k {
			delete(m , getid( nums[i - k], t + 1))
		}
	}
	return false
}

func getid(a , b int) int {
	if a >= 0 {
		return a / b
	}
	return (a + 1) / b - 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(containsNearbyAlmostDuplicate([]int{1,2,3,1} , 3 , 0))
}
