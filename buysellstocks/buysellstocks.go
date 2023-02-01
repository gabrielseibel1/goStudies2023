package buysellstocks

import "math"

func MaxProfit(prices []int) int {
	profit := 0
	holding := false
	boughtPrice := math.MaxInt

	for i := 0; i < len(prices); i++ {
		if i+1 < len(prices) {
			if prices[i] <= prices[i+1] { // buy or hold if price will go up
				if !holding {
					holding = true
					boughtPrice = prices[i]
				}
			} else if holding { // sell if price will go down
				holding = false
				profit += prices[i] - boughtPrice
			}
		} else if holding { // sell if we are holding but there is no tomorrow
			profit += prices[i] - boughtPrice
		}
	}
	return profit
}
