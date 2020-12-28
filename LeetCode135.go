package main

import "fmt"

func candy(ratings []int) int {

	n := len(ratings)
	ans , inc , dec , pre := 1 , 1 , 0 ,1
	for i := 0 ; i < n ; i ++ {
		if ratings[i] >= ratings[i - 1] {
			dec = 0
			if ratings[i] == ratings[i - 1] {
				pre = 1
			} else {
				pre ++
			}
			ans += pre
			inc  = pre
		} else {
			dec ++
			if dec == inc {
				dec ++
			}
			ans += dec
			pre  = 1
		}
	}
	return ans
}

func main() {
	fmt.Println(candy([]int{1,2,2}))
}