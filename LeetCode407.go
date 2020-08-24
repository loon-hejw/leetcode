package main

import "fmt"

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 { return 0 }
	min := func ( a ,b int) int {
		if  a > b {
			return b
		}
		return a
	}
	// 初始变量
	ans := 0
	for i := 1 ; i < len(heightMap) - 1 ; i ++ {
		height := heightMap[i]
		l,r := 0 , len(height) - 1
		// 寻找边界
		for l < r && height[l] <= height[l + 1] { l ++ }
		for l < r && height[r] <= height[r - 1] { r -- }
		left , right , top , dump := 0 , 0 , 0 , 0
		for l < r {
			left = height[l]
			right = height[r]

			if left <= right {
				for l = l + 1 ; l < r && left >= height[l] ; l ++ {
					u  , t := 0 , len(heightMap) - 1
					for  u < i && heightMap[u][l] <= heightMap[u + 1][l] { u ++ }
					for  t > i && heightMap[t][l] <= heightMap[t - 1][l] { t -- }
					top = heightMap[u][l] ; dump = heightMap[t][l]
					minH := min(min(top,dump),left)
					if  minH >= height[l] {
						ans = ans + (minH - height[l])
					}
				}
			} else {
				for r = r - 1 ; l < r && right >= height[r]; r --  {
					u  , t := 0 , len(heightMap) - 1
					for  u < i && heightMap[u][r] <= heightMap[u + 1][r] { u ++ }
					for  t > i && heightMap[t][r] <= heightMap[t - 1][r] { t -- }
					top = heightMap[u][r] ; dump = heightMap[t][r]

					minH := min(min(top,dump),right)
					if  minH >= height[r] {
						ans = ans + (minH - height[r])
					}
				}
			}
		}
	}
	return ans
}
func main() {
	fmt.Println(trapRainWater([][]int{{12,13,1,12},{13,4,13,12},{13,8,10,12},{12,13,12,12},{13,13,13,13}}))
}
