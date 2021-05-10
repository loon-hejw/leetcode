package main

import "math"
type entry struct {
	first, firstIdx, second int
}

func minCost(houses []int, cost [][]int, m, n, target int) int {
	const inf = math.MaxInt64 / 2 // 防止整数相加溢出

	// 将颜色调整为从 0 开始编号，没有被涂色标记为 -1
	for i := range houses {
		houses[i]--
	}

	// dp 所有元素初始化为极大值
	dp := make([][][]int, m)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, target)
			for k := range dp[i][j] {
				dp[i][j][k] = inf
			}
		}
	}
	best := make([][]entry, m)
	for i := range best {
		best[i] = make([]entry, target)
		for j := range best[i] {
			best[i][j] = entry{inf, -1, inf}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if houses[i] != -1 && houses[i] != j {
				continue
			}

			for k := 0; k < target; k++ {
				if i == 0 {
					if k == 0 {
						dp[i][j][k] = 0
					}
				} else {
					dp[i][j][k] = dp[i-1][j][k]
					if k > 0 {
						// 使用 best(i-1,k-1) 直接得到 dp(i,j,k) 的值
						if b := best[i-1][k-1]; j == b.firstIdx {
							dp[i][j][k] = min(dp[i][j][k], b.second)
						} else {
							dp[i][j][k] = min(dp[i][j][k], b.first)
						}
					}
				}

				if dp[i][j][k] != inf && houses[i] == -1 {
					dp[i][j][k] += cost[i][j]
				}

				// 用 dp(i,j,k) 更新 best(i,k)
				if b := &best[i][k]; dp[i][j][k] < b.first {
					b.second = b.first
					b.first = dp[i][j][k]
					b.firstIdx = j
				} else if dp[i][j][k] < b.second {
					b.second = dp[i][j][k]
				}
			}
		}
	}

	ans := inf
	for _, res := range dp[m-1] {
		ans = min(ans, res[target-1])
	}
	if ans == inf {
		return -1
	}
	return ans
}

func main() {

}
