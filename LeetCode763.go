package main

func partitionLabels(S string) []int {
	aChar := [26]int{}
	for i , c := range S {
		aChar[c - 'a'] = i
	}

	res   := []int{}
	start := 0
	end   := 0
	for i , c := range S {
		if aChar[c - 'a'] > end {
			end = aChar[c - 'a']
		}
		if i == end {
			res = append(res , end - start + 1)
			start = end + 1
		}
	}
	return res
}

func main() {
	
}
