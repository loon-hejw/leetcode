package main

func numEquivDominoPairs(dominoes [][]int) int {
	q := make([]int , 100)
	ans := 0
	for _ , v := range dominoes {
		a , b := ReturnTwo(v[0] , v[1])
		t := a * 10  + b
		ans += q[t]
		q[t] ++
	}
	return ans
}

func ReturnTwo (a , b int) (int , int) {
	if a > b {
		return b , a
	}
	return a , b
}

func main() {
	
}
