package main

func maxProfit(k int, prices []int) int {
	dp := make([]int,k)
	max := func( a , b int ) int {
		if a > b {
			return a
		}
		return b
	}
	for k := range prices {
		for i := 0 ; i < k ; i ++ {
			dp[i] = max( prices[k] , dp[i] )
		}
	}
	return 0
}
func main() {
	
}
