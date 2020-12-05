package main

import "fmt"

type N int
func v(n N) {
	fmt.Println("v %v:%p" , &n , n)
	n = 0
}

func v1(n *N) {
	fmt.Println("v %v:%p" , &n , n)
	*n = 0
}

func main() {
	var n N
	n = 1
	v(n)
	v1(&n)
	fmt.Println( "n %v:%p" ,&n , n)
}
