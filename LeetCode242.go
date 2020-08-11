package main

func isAnagram(s string, t string) bool {
	if len(s) != len(t) { return false }
	lenght := map[int]int{}
	for k := 0 ; k < len(s) ; k ++ {

		if lenght[int(s[k])] == 0 {
			lenght[int(s[k])] = 1
		} else {
			lenght[int(s[k])] = lenght[int(s[k])] + 1
		}
		if lenght[int(t[k])] == 0 {
			lenght[int(t[k])] = -1
		} else {
			lenght[int(t[k])] = lenght[int(t[k])] - 1
		}
	}

	for _ , v := range(lenght) {
		if v != 0 {
			return false
		}
	}
	return true
}
func main() {
	isAnagram("anagram","nagaram")
}
