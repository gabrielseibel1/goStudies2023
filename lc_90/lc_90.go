package lc90

import (
	"fmt"
	"math"
	"slices"
)

func subsetsWithDup(nums []int) [][]int {
	powerSet := make(map[string][]int)

	state := make([]bool, len(nums))
	bitStrings := make([][]bool, int(math.Exp2(float64(len(nums)))))
	generateBitStrings(0, state, &bitStrings)

	for _, bits := range bitStrings {
		subset := getCombination(nums, bits)
		slices.Sort(subset)
		powerSet[fmt.Sprint(subset)] = subset
	}

	answer := make([][]int, len(powerSet))
	i := 0
	for _, set := range powerSet {
		answer[i] = set
		i++
	}

	return answer
}

func generateBitStrings(i int, state []bool, bitStrings *[][]bool) {
	if i == len(state) {
		*bitStrings = append(*bitStrings, slices.Clone(state))
		return
	}
	state[i] = true
	generateBitStrings(i+1, state, bitStrings)
	state[i] = false
	generateBitStrings(i+1, state, bitStrings)
}

func getCombination(nums []int, bits []bool) []int {
	subset := make([]int, 0)
	for i, bit := range bits {
		if bit {
			subset = append(subset, nums[i])
		}
	}
	return subset
}
