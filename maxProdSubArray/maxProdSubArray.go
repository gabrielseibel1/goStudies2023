package maxprodsubarray

import "fmt"

func MaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	memo := make(map[string][]int)
	_, max := minMaxProds(nums, memo, 0, false)
	return max
}

func minMaxProds(nums []int, memo map[string][]int, i int, contiguous bool) (min int, max int) {
	if ans, ok := memo[fmt.Sprint(i, contiguous)]; ok {
		return ans[0], ans[1]
	}

	if i+1 == len(nums) {
		min, max = nums[i], nums[i]
	} else {
		nextMinContiguous, nextMaxContiguous := 0, 0
		if nums[i] != 0 {
			nextMinContiguous, nextMaxContiguous = minMaxProds(nums, memo, i+1, true)
		}
		if contiguous {
			min = minOf(nums[i], nums[i]*nextMinContiguous, nums[i]*nextMaxContiguous)
			max = maxOf(nums[i], nums[i]*nextMinContiguous, nums[i]*nextMaxContiguous)
		} else {
			nextMinGeneral, nextMaxGeneral := minMaxProds(nums, memo, i+1, false)
			min = minOf(nums[i], nums[i]*nextMinContiguous, nums[i]*nextMaxContiguous, nextMinGeneral, nextMaxGeneral)
			max = maxOf(nums[i], nums[i]*nextMinContiguous, nums[i]*nextMaxContiguous, nextMinGeneral, nextMaxGeneral)
		}
	}

	ans := []int{min, max}
	memo[fmt.Sprint(i, contiguous)] = ans

	return min, max
}

func minOf(s ...int) int {
	min := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] < min {
			min = s[i]
		}
	}
	return min
}

func maxOf(s ...int) int {
	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
		}
	}
	return max
}
