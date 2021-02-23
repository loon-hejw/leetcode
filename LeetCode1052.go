package main

func maxSatisfied(customers []int, grumpy []int, X int) int {

	total := 0
	n := len(customers)

	for i := 0 ; i < n ; i ++ {
		if grumpy[i] == 0 {
			total += customers[i]
		}
	}

	increase := 0
	for i := 0 ; i < X ; i ++ {
		increase += customers[i] * grumpy[i]
	}
	maxIncrase := 0
	for i := X ; i < n ; i ++ {
		increase = increase - customers[i - X] * grumpy[i - X] + customers[i] * grumpy[i]
		maxIncrase = max(increase , maxIncrase)
	}
	return total + maxIncrase
}

func main() {
	
}
