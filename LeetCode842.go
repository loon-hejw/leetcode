package main

import "math"

func splitIntoFibonacci(S string) []int {

	n := len(S)
	ans := []int{}
	var backtrack func(index , sum , prev int) bool
	backtrack = func(index, sum, prev int) bool {

		if index == n {
			return len(ans) >= 3
		}
		cur := 0
		for i := index ; i < n ; i ++ {
			if i > index && S[index] == '0' {
				return false
			}

			cur = cur * 10 + int(S[i] - '0')
			if cur > math.MaxInt32 {
				return false
			}

			if len(ans) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}

			ans = append(ans , cur)
			if backtrack(i + 1 , prev + cur , cur) {
				return true
			}
			ans = ans[:len(ans) - 1]
		}
		return false
	}
	backtrack(0 , 0 , 0)
	return ans
}

func main() {
	
}
