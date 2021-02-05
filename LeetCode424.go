package main

func characterReplacement(s string, k int) int {

	cnt := [26]int{}
	maxCnt , left := 0 ,0
	max := func( a , b int) int {
		if a < b {
			return b
		}
		return a
	}
	for  right , ch := range s {
		cnt[ch - 'A'] ++
		maxCnt = max(maxCnt , cnt[ch - 'A'])
		if right - left + 1 - maxCnt > k {
			cnt[s[left] - 'A'] --
			left ++
		}
	}
	return len(s) - left
}

func main() {
	
}
