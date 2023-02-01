package lis

func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lis := make([]int, len(nums))
	lis[0] = 1
	bestLIS := 1
	for i := 1; i < len(lis); i++ {
		bestLastLIS := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				bestLastLIS = max(bestLastLIS, lis[j])
			}
		}
		lis[i] = bestLastLIS + 1
		bestLIS = max(bestLIS, lis[i])
	}
	return bestLIS
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
