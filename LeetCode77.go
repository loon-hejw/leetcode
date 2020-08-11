package main

import "fmt"

func combine2(n int, k int) [][]int {
	res := [][]int{}

	if k > n || k < 0 {
		return res
	}

	if k == 0 {
		res = append(res, []int{})
		return res
	}

	res = combine2(n - 1 , k - 1)
	for i := 0 ; i < len(res) ; i ++ {
		res[i] = append(res[i], n)
	}

	temp := combine2(n - 1 , k)
	for i := 0 ; i < len(temp) ; i ++ {
		res = append(res, temp[i])
	}

	return res
}

func combine(n int, k int) [][]int {
	return backcombine([]int{},[][]int{},n,k,1)
}

func backcombine(list []int ,res [][]int , n int , k int , start int) [][]int {

	if len(list) == k {
		temp := make([]int , len(list))
		copy(temp,list)
		res = append(res , temp)
		return res
	}

	for i := start ; i <= n ; i ++ {
		list = append(list , i)
		fmt.Println(list)
		res = backcombine(list , res , n , k ,i + 1)
		list = list[:len(list) - 1]
	}

	return res
}
func main() {
	fmt.Println(combine2(4 ,2))
}
