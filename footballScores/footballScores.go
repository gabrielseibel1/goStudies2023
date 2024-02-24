package footballScores

import "sort"

func countMatches(teamA, teamB []int) []int {
	result := make([]int, 0)

	sort.Slice(teamA, func(i, j int) bool {
		return teamA[i] < teamA[j]
	})

	for _, goalsB := range teamB {
		count := sort.Search(len(teamA), func(i int) bool {
			return teamA[i] > goalsB
		})
		result = append(result, count)
	}

	return result
}
