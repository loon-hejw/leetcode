package main

import (
	"fmt"
	"math/rand"
)

// 冒泡排序
func bubbleSort(array []int) []int {
	for i := 0 ; i < len(array) - 2 ; i ++ {
		for j := 0 ; j < len(array) - 1 - i ; j ++ {
			if array[j] > array[ j + 1 ] {
				array[j] ,array[j + 1] = array[j + 1] , array[j]
			}
		}
	}
	return array
}

// 选择排序
func selectionSort(array []int) []int {
	for i := 0 ; i < len(array) - 1 ; i ++ {
		for j := i + 1 ; j < len(array) ; j ++ {
			if array[i] > array[j] {
				array[i] , array[j] = array[j] , array[i]
			}
		}
	}
	return array
}

// 插入排序
func insertSort(array []int) []int {
	for i := 1 ; i < len(array) ; i ++ {
		index   := i - 1
		current := array[i]
		for index >= 0 && array[index] > current {
			array[ index + 1 ] = array[index]
			index --
		}
		array[index + 1] = current
	}
	return array
}

// 希尔排序
func shellSort(array []int) []int {
	for i := len(array) / 2 ; i > 0 ; i = i / 2 {
		for j := i ; j < len(array) ; j ++ {
			k , n := j , array[j]
			for  k - i >= 0 && n < array[k - i] {
				array[k] = array[k - i]
				k = k - i
			}
			array[k] = n
		}
	}
	return array
}

// 归并排序
func mergeSort(array []int , left , right int){
	if right <= left { return }
	mid := (left + right) >> 1
	mergeSort( array , left , mid )
	mergeSort( array , mid + 1 , right)
 	merge( array , left , mid , right)
}

func merge(array []int , left , mid , right int) {
	temp := make([]int , right - left + 1)
	i , j , k := left , mid + 1 , 0
	for i <= mid && j <= right {
		if array[i] <= array[j] {
			temp[k] = array[i] ; i ++ ; k ++
		} else {
			temp[k] = array[j] ; j ++ ; k ++
		}
	}
	for i <= mid   { temp[k] = array[i] ; k ++ ; i ++ }
	for j <= right { temp[k] = array[j] ; k ++ ; j ++ }
	for i := 0 ; i < len(temp) ; i ++ { array[left + i] = temp[i] }
}

// 快速排序
func quickSork (array []int) []int {
	length := len(array)
	if length <= 1 {
		arrayCopy := make([]int,length)
		copy(arrayCopy,array)
		return arrayCopy
	}
	m := array[rand.Intn(length)]
	less := make([]int , 0 , length)
	middle := make([]int , 0 , length)
	more  := make([]int , 0 , length)
	for _ , item := range array {
		if item < m {
			less   = append( less , item )
		} else if item == m {
			middle = append( middle , item )
		} else {
			more   = append( more , item )
		}
	}
	less , more = quickSork(less) , quickSork(more)
	less = append( less , middle... )
	less = append( less , more... )
	return less
}

// 堆排序
func heapsort (array []int) []int {
	ep := (len(array) - 1) >> 1
	for i := ep ; i >= 0 ; i -- {
		heapt(array , i , len(array) - 1)
	}
	for i := len(array) - 1 ; i > 0 ; i -- {
		array[0] ,  array[i] = array[i] , array[0]
		heapt(array, 0 , i - 1)
	}
	return array
}
func heapt (array []int , start , end int) {
	le := start * 2 + 1
	re := le + 1
	if le > end { return }
	tmp := le
	if re <= end && array[re] > array[le] {
		tmp = re
	}
	if array[tmp] > array[start] {
		array[start] , array[tmp] = array[tmp] , array[start]
		heapt( array , tmp , end )
	}
}

// 计数排序
func countingSort(array []int) []int {
	max := 0
	for _ , v := range array {
		if max < v { max = v }
	}
	tempArray := make([]int, max + 1)
	for _ , v := range array { tempArray[v] ++ }
	sortArray := []int{}
	for k , v := range tempArray {
		for v > 0 {
			sortArray = append(sortArray,k)
			v --
		}
	}
	return sortArray
}

// 桶排序
//func bucketSort(array []int) []int {
	//max := 0
	//for _ , v := range array {
	//	if max < v { max = v }
	//}
	//tempArray := make([]int, max + 1)
	//for _ , v := range array { tempArray[v] ++ }
	//sortArray := []int{}
	//for k , v := range tempArray {
	//	for v > 0 {
	//		sortArray = append(sortArray,k)
	//		v --
	//	}
	//}
	//return sortArray
//}

// 基数排序
func radixSort(array []int) []int {

	baseArray := make([][]int,10)
	mod , dev , max  := 10 , 1 , 0
	for _ , v := range array {
		if v > max { max = v }
	}
	for max > 0  {
		for _ , v := range array {
			temp := (v % mod) / dev
			baseArray[temp] = append(baseArray[temp],v )
		}
		array = []int{}
		for i := 0 ; i < len(baseArray) ; i ++{
			for j := 0 ; j < len(baseArray[i]) ; j ++ {
				array = append(array , baseArray[i][j])
			}
		}
		baseArray = make([][]int,10)
		max = max / 10 ; mod = mod * 10 ; dev = dev * 10
	}
	return array
}

func main() {
	temp := []int{3,44,38,5,47,15,36,26,27,2,46,4,19,50,48}
	//fmt.Println(selectionSort(temp))
	//fmt.Println(bubbleSort(temp))
	//fmt.Println(shellSort(temp))
	//mergeSort(temp,0,len(temp) - 1)
	//fmt.Println(quickSork(temp))
	//fmt.Println(heapsort(temp))
	//fmt.Println(bucketSort(temp))
	fmt.Println(radixSort(temp))
}
