package lc46

import (
	"fmt"
	"maps"
	"math"
	"slices"
)

func permute(nums []int) [][]int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}
	state := make([]int, len(nums))
	for i := range state {
		state[i] = math.MinInt
	}
	memo := make(map[string]map[string][]int)
	perms := permutations(m, state, make([]bool, len(nums)), memo)
	result := make([][]int, 0, len(perms))
	for _, p := range perms {
		result = append(result, p)
	}
	return result
}

func permutations(m map[int]struct{}, state []int, fixed []bool, memo map[string]map[string][]int) map[string][]int {
	if p, ok := memo[memoKey(m, state, fixed)]; ok {
		return p
	}

	// if no position available just return the state
	anyAvailable := false
	for _, f := range fixed {
		if !f {
			anyAvailable = true
			break
		}
	}
	if !anyAvailable {
		return map[string][]int{stateKey(state): slices.Clone(state)}
	}

	// if no elements available return nil
	if len(m) == 0 {
		return map[string][]int{}
	}

	// fill state positions that are available with m values
	perms := make(map[string][]int)
	mAux := make(map[int]struct{}, len(m))
	maps.Copy(mAux, m)
	for i := range state {
		if !fixed[i] {
			for n := range m {
				delete(mAux, n)
				state[i] = n
				fixed[i] = true

				for _, perm := range permutations(mAux, state, fixed, memo) {
					perms[stateKey(perm)] = perm
				}

				state[i] = math.MinInt
				fixed[i] = false
				mAux[n] = struct{}{}
			}
		}
	}

	memo[memoKey(m, state, fixed)] = perms

	return perms
}

func stateKey(state []int) string {
	return fmt.Sprintf("%v", state)
}

func memoKey(m map[int]struct{}, state []int, fixed []bool) string {
	return fmt.Sprintf("%v_%v_%v", m, state, fixed)
}
