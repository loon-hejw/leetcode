package main

func canPartition(nums []int) bool {

	maxFunc := func(a ,b int) int {
		if a > b {
			return a
		}
		return b
	}
	sum := 0
	max := 0
	for _ , v := range nums {
		sum = sum + v
		max = maxFunc(max , v)
	}
	if sum & 1 != 0 {
		return false
	}
	targer := sum >> 1
	if max > targer {
		return false
	}
	dp := make([][]bool,len(nums))
	for i := 0 ; i < len(nums) ; i ++ {
		dp[i] = make([]bool,targer + 1)
		dp[i][0] = true
	}

	dp[0][nums[0]] = true
	for i := 1 ; i < len(nums) ; i ++ {
		num := nums[i]
		for j := 0 ; j <= targer ; j ++ {
			dp[i][j] = dp[i - 1][j]
			if j >= num {
				dp[i][j] = dp[i - 1][j] || dp[i - 1][j - num]
			}
		}
	}
	return dp[len(nums) - 1][targer]
}

func main() {
	
}
