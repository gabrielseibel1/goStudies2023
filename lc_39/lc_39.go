package lc39

import (
	"slices"
)

func CombinationSum(candidates []int, target int) [][]int {
	slices.Sort(candidates)
	tip := 0
	for _, n := range candidates {
		tip++
		if n > target {
			break
		}
	}
	candidates = candidates[:tip]
	slices.Reverse(candidates)
	return combinationSumFromProperInput(candidates, target)
}

func combinationSumFromProperInput(candidates []int, target int) [][]int {
	if target == 1 || len(candidates) == 0 {
		return [][]int{}
	}
	combos := make([][]int, 0)
	for i, n := range candidates {
		if n == target {
			combos = append(combos, []int{n})
		}
		if n > target {
			continue
		}
		if newTarget := target - n; newTarget > 0 {
			next := combinationSumFromProperInput(candidates[i:], newTarget)
			for j := range next {
				next[j] = append([]int{n}, next[j]...)
			}
			combos = append(combos, next...)
		}
	}
	return combos
}
