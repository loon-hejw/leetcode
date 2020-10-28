package main

func uniqueOccurrences(arr []int) bool {

	countMap := map[int]int{}
	for _ , v := range arr {
		countMap[v] ++
	}
	lenghtMap := map[int]bool{}
	for _ , v := range countMap {
		if _ , ok := lenghtMap[v] ; ok {
			return false
		}
		lenghtMap[v] = true
	}
	return true
}
func main() {
	
}
