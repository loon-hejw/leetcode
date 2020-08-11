package main

func isPerfectSquare1(num int) bool {
	p := 1
	for num > 0 {
		num = num - p
		p = p + 2
	}
	return num == 0
}

func isPerfectSq(num int) bool {

	if num < 2 {
		return true
	}

	left , right := 2 , num / 2

	for left <= right {
		x := left + ( right - left ) / 2

		guss := x * x
		if guss == num {
			return true
		}

		if guss > num {
			right = x - 1
		} else {
			left = x + 1
		}
	}
	return false
}

func main() {
	
}
