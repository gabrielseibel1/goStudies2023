package lc746

func minCostClimbingStairs(cost []int) int {
	i := 2
	for ; i < len(cost); i++ {
		cost[i] += min(cost[i-1], cost[i-2])
	}
	return min(cost[i-1], cost[i-2])
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// mandar vaga aws pro klein
// mandar grupos pro klein
