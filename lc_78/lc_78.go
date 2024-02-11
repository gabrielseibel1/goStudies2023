package lc78

import "slices"

func subsets(nums []int) [][]int {
	ps := make([][]int, 0)

	ps = append(ps, []int{})
	for i := range nums {
		ps = append(ps, slices.Clone(combinations(nums, i+1))...)
	}

	return ps
}

func combinations(slice []int, size int) [][]int {
	if len(slice) <= 0 || size <= 0 || len(slice) < size {
		return [][]int{}
	}
	if len(slice) == size {
		return [][]int{slice}
	}

	combos := make([][]int, 0)
	for i, n := range slice {
		if len(slice)-i < size {
			break
		}
		nextCombos := slices.Clone(combinations(slice[i+1:], size-1))

		if len(nextCombos) == 0 {
			combos = append(combos, []int{n})
			continue
		}

		for i, nextCombination := range nextCombos {
			combosStartingInI := append([]int{n}, nextCombination...)
			nextCombos[i] = slices.Clone(combosStartingInI)
		}
		combos = append(combos, nextCombos...)
	}
	return combos
}
