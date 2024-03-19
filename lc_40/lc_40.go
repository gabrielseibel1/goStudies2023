package lc40

import (
	"slices"
)

type combination struct {
	numbers []int
	sum     int
}

func combinationSum2(candidates []int, target int) [][]int {
	// find all combinations that sum up to target
	// there are only positive numbers and positive target
	// no repeated indices used (repeated numbers are allowed if they are in the slice)
	// no duplicate combinations

	state := combination{numbers: make([]int, 0)}
	combinations := make([][]int, 0)

	initCombinationsThatSumToTarget(candidates, target, &state, &combinations)

	slices.SortFunc(combinations, func(s1, s2 []int) int {
		if len(s1) < len(s2) {
			return -1
		}
		if len(s1) > len(s2) {
			return 1
		}
		slices.Sort(s1)
		slices.Sort(s2)
		for i := range s1 {
			if s1[i] < s2[i] {
				return -1
			}
			if s1[i] > s2[i] {
				return 1
			}
		}
		return 0
	})
	compact := slices.CompactFunc(combinations, func(s1, s2 []int) bool {
		eq := true
		if len(s1) != len(s2) {
			return false
		}
		for _, v := range s1 {
			if !slices.Contains(s2, v) {
				eq = false
			}
		}
		for _, v := range s2 {
			if !slices.Contains(s1, v) {
				eq = false
			}
		}
		return eq
	})
	return compact
}

func initCombinationsThatSumToTarget(candidates []int, target int, state *combination, combinations *[][]int) {
	state.sum += candidates[0]
	state.numbers = append(state.numbers, candidates[0])
	combinationsThatSumToTarget(1, candidates, target, state, combinations)

	state.sum -= candidates[0]
	state.numbers = state.numbers[:len(state.numbers)-1]
	combinationsThatSumToTarget(1, candidates, target, state, combinations)
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
