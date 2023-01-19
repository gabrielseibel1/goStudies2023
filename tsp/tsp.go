package tsp

import "math"

func MinimumCostTour(m [][]int, s int) (int, []int) {
	if len(m) == 1 { // only one node
		return m[0][0], []int{0, 0}
	}
	if len(m) == 2 {
		return m[0][1] + m[1][0], []int{0, 1, 0}
	}

	tab := initTab(len(m), int(math.Pow(2, float64(len(m)))))

	setupTab(&m, tab, s)
	solveUpToLastNode(&m, tab, s)
	dist := findMinDistToReturn(&m, tab, s)
	tour := listTour(&m, tab, s)
	return dist, tour
}

func initTab(row, col int) *[][]*int {
	tab := make([][]*int, row)
	for i := 0; i < row; i++ {
		tab[i] = make([]*int, col)
	}
	return &tab
}

func setupTab(m *[][]int, tab *[][]*int, s int) {
	for i := 0; i < len(*m); i++ {
		if i == s {
			continue
		}
		(*tab)[i][1<<s|1<<i] = &((*m)[s][i])
	}
}

func solveUpToLastNode(m *[][]int, tab *[][]*int, s int) {
	// for each level (with l nodes to visit)
	for l := 3; l <= len(*m); l++ {
		// for each possible combination of l visited nodes that include origin
		for _, subset := range combinations(l, len(*m)) {
			if notIn(subset, s) {
				continue
			}
			// for each node to be considered the next (just added to path)
			for next := 0; next < len(*m); next++ {
				if notIn(subset, next) || next == s {
					continue
				}
				// find last/anterior node which minimizes distance to next based on state
				state := subset ^ (1 << next)
				minDist := math.MaxInt
				for ant := 0; ant < len(*m); ant++ {
					if ant == s || ant == next || notIn(subset, ant) {
						continue
					}
					dist := *(*tab)[ant][state] + (*m)[ant][next]
					minDist = int(math.Min(float64(dist), float64(minDist)))
				}
				(*tab)[next][subset] = &minDist
			}
		}
	}
}

func notIn(subset, i int) bool {
	return (subset & (1 << i)) == 0
}

func combinations(on, n int) []int {
	combs := make([]int, 0)
	if on == 0 {
		return []int{0}
	}
	combinationsRecursive(on, n, 0, 0, 0, &combs)
	combinationsRecursive(on, n, 0|(1<<0), 0, 1, &combs)
	return combs
}

func combinationsRecursive(mustOn, n, state, i, ones int, combs *[]int) {
	if ones == mustOn {
		*combs = append(*combs, state)
		return
	}
	if mustOn-ones > n-i-1 { // no space left for more ones
		return
	}

	combinationsRecursive(mustOn, n, state, i+1, ones, combs)
	combinationsRecursive(mustOn, n, state|(1<<(i+1)), i+1, ones+1, combs)
}

func findMinDistToReturn(m *[][]int, tab *[][]*int, s int) int {
	endState := (1 << len(*m)) - 1 // 2^n - 1, all ones

	// minimize tour + dist to origin
	minTourCost := math.MaxInt
	for last := 0; last < len(*m); last++ {
		if last == s {
			continue
		}
		tourCost := *(*tab)[last][endState] + (*m)[last][s]
		if tourCost < minTourCost {
			minTourCost = tourCost
		}
	}
	return minTourCost
}

func listTour(m *[][]int, tab *[][]*int, s int) []int {
	lastNode := s
	state := (1 << len(*m)) - 1 // 2^n - 1, all ones

	tour := make([]int, len(*m)+1)
	tour[0], tour[len(*m)] = s, s

	// for each level of the tour (level has l traversed nodes)
	for l := len(*m) - 1; l >= 1; l-- {
		// which node was the best at the level
		best := -1
		for n := 0; n < len(*m); n++ {
			if n == s || notIn(state, n) {
				continue
			}
			if best == -1 {
				best = n
			}
			nDist := *(*tab)[n][state] + (*m)[n][lastNode]
			bDist := *(*tab)[best][state] + (*m)[best][lastNode]
			if nDist < bDist {
				best = n
			}
		}

		tour[l] = best
		state = state ^ (1 << best)
		lastNode = best
	}

	return tour
}
