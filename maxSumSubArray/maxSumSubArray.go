package maxsumsubarray

// https://leetcode.com/problems/maximum-subarray

import "math"

func MaxSubArray(nums []int) int {
	return maxSubArraySum(nums, 0, len(nums)-1)
}

func maxSubArraySum(a []int, from, to int) int {
	if from == to {
		return a[from]
	}

	mean := (from + to) / 2
	leftSum := maxSubArraySum(a, from, mean)
	rightSum := maxSubArraySum(a, mean+1, to)
	centerSum := centerArraySum(a, from, mean, to)

	return int(math.Max(
		math.Max(float64(leftSum), float64(rightSum)),
		float64(centerSum)))
}

func centerArraySum(a []int, from, mean, to int) int {
	leftSum := 0
	maxLeftSum := math.MinInt
	for i := mean; i >= from; i-- {
		leftSum += a[i]
		maxLeftSum = int(math.Max(
			float64(maxLeftSum), float64(leftSum)))
	}
	rightSum := 0
	maxRightSum := math.MinInt
	for i := mean + 1; i <= to; i++ {
		rightSum += a[i]
		maxRightSum = int(math.Max(
			float64(maxRightSum), float64(rightSum)))
	}
	return maxLeftSum + maxRightSum
}
