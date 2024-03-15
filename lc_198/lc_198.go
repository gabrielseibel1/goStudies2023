package lc198

func Rob(nums []int) int {
	switch len(nums) {
	case 0:
		return 0
	case 1:
		return nums[0]
	case 2:
		return max(nums[0], nums[1])
	}
	i := 2
	for ; i < len(nums); i++ {
		nums[i] += max(nums[i-2], atOr0(nums, i-3))
	}
	return max(nums[i-1], nums[i-2])
}

func atOr0(nums []int, i int) int {
	if i < 0 {
		return 0
	}
	return nums[i]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
