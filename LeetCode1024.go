package main

func videoStitching(clips [][]int, T int) int {

	max := make([]int , T)
	last , pre := 0 ,0
	for _ , c := range clips {
		if c[0] < T && c[1] > max[1] {
			max[c[0]] = c[1]
		}
	}
	ans := 0
	for i ,v := range max {
		if v > last {
			last = v
		}
		if i == last {
			return -1
		}
		if i == pre {
			ans++
			pre  = last
		}
	}
	return ans
}

func main() {
	
}
