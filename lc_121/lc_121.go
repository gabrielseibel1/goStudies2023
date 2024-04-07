package lc121

func maxProfit(prices []int) int {
	maxV := prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] > maxV {
			maxV = prices[i]
		}
		prices[i] = maxV - prices[i]
	}
	profitV := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] > profitV {
			profitV = prices[i]
		}
	}
	return profitV
}
