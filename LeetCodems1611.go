package main

func divingBoard(shorter int, longer int, k int) []int {
	lenght := []int{}
	if k == 0 {
		return lenght
	}
	if shorter == longer {
		return []int{shorter * k}
	}

	for i := 0 ; i <= k ; i ++ {
		lenght = append(lenght, shorter * (k - i) + longer * i)
	}
	return lenght
}
func main() {
	
}
