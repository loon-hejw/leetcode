package main

import "fmt"

func findLength(A []int, B []int) int {
	N , M  := len(A) , len(B)
	ans    := 0

	for i := 0 ; i < N ; i ++ {
		Len := Min(M , N - i)
		ans = MaxLenght(A , B , i , 0 , Len , ans)
	}
	for i := 0 ; i < M ; i ++ {
		Len := Min(N , M - i)
		ans  = MaxLenght(A , B , 0 , i , Len , ans)
	}

	return ans
}

func MaxLenght(A , B []int , addA , addB , len int , ans int) int {
	res ,k := 0 , 0
	for i := 0 ; i < len ; i ++ {
		if A[addA + i] == B[addB + i] {
			k ++
		} else {
			k = 0
		}
		res = Max(res , k)
	}
	return Max(ans,res)
}

func Max(a int , b int ) int {
	if a > b {
		return a
	}
	return b
}
func Min (a int , b int) int {
	if a > b {
		return b
	}
	return a
}
func main() {
	a := findLength([]int{1,1,0,0,1},[]int{1,1,0,0,0})
	fmt.Println(a)
}
