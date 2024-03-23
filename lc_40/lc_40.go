package lc40

import (
	"slices"
)

type combination struct {
	numbers []int
	sum     int
}

func combinationSum2(candidates []int, target int) [][]int {
	state := combination{numbers: make([]int, 0)}
	combinations := make([][]int, 0)
	slices.Sort(candidates)
	combinationsThatSumToTarget(0, candidates, target, &state, &combinations)
	return combinations
}

func combinationsThatSumToTarget(i int, candidates []int, target int, state *combination, combinations *[][]int) {
	if state.sum == target {
		*combinations = append(*combinations, slices.Clone(state.numbers))
		return
	}
	if state.sum > target {
		return
	}
	if i >= len(candidates) {
		return
	}

	// traverse combinations tree with ith element selected
	state.sum += candidates[i]
	state.numbers = append(state.numbers, candidates[i])
	combinationsThatSumToTarget(i+1, candidates, target, state, combinations)

	// traverse combinations tree with ith element skipped
	state.sum -= candidates[i]
	state.numbers = state.numbers[:len(state.numbers)-1]
	var j int
	for j = i + 1; j < len(candidates) && candidates[j] == candidates[i]; j++ {
	}
	combinationsThatSumToTarget(j, candidates, target, state, combinations)
}
