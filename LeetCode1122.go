package main

import (
	"fmt"
	"sort"
)

//func relativeSortArray(arr1 []int, arr2 []int) []int {
//
//}
func main() {

	arr1 := []int{2,3,1,3,2,4,6,7,9,2,19}
	arr2 := []int{2,1,4,3,9,6}
	array := make([]int,1001)

	sort.Ints(arr1)
	for _ , i := range arr1{
		array[i] = array[i] + 1
	}
	arr3 := []int{}
	for _ , v := range arr2{
		for array[v] > 0 {
			arr3 = append(arr3 , v)
			array[v] = array[v] - 1
		}
	}

	for _ , v := range arr1 {
		if array[v] != 0 {
			arr3 = append(arr3 , v)
			array[v] = array[v] - 1
		}
	}

	fmt.Println(arr3)
}