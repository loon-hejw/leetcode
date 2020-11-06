package main

func insert(intervals [][]int, newInterval []int) [][]int {
	left  := newInterval[0]
	right := newInterval[1]
	min := func(a , b int) int {
		if a > b {
			return b
		}
		return a
	}
	merged := false
	ans := [][]int{}
	for _ , v := range intervals {
		if right < v[0] {
			if !merged {
				ans = append(ans , []int{left , right})
				merged = true
			}
			ans = append(ans , v)
		} else if left > v[1] {
			ans = append(ans , v)
		} else {
			left  = (left + v[0]) - min(left , v[0])
			right = min(right, v[1])
		}
	}
	if !merged {
		ans = append(ans , []int{left , right})
	}
	return ans
}

func main() {
	
}
