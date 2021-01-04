package main

import "fmt"

func fib(n int) int {
	f1 , f2 := 0 , 1
	for i := 2 ; i <= n + 1 ; i ++ {
		f1 , f2 = f2 , f1 + f2
	}
	return f1
}

func main() {
	fmt.Println(fib(0))
}
