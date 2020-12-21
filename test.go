package main

import (
	"fmt"
	"strconv"
)

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

 t  , err := strconv.ParseFloat("23.98",64)

 fmt.Println(t)
 fmt.Println(err)
}
