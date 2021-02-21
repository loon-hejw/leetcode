package main

import "fmt"

func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	// 滑动窗口大小为 n-k
	windowSize := n - k
	// 选前 n-k 个作为初始值
	sum := 0
	for _, pt := range cardPoints[:windowSize] {
		sum += pt
	}
	minSum := sum
	for i := windowSize; i < n; i++ {
		// 滑动窗口每向右移动一格，增加从右侧进入窗口的元素值，并减少从左侧离开窗口的元素值
		sum += cardPoints[i] - cardPoints[i-windowSize]
		minSum = min(minSum, sum)
	}
	total := 0
	for _, pt := range cardPoints {
		total += pt
	}
	return total - minSum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxScore([]int{1,2,3,4,5,6,1},3))
}
