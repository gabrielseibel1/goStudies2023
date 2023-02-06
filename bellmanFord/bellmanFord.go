package bellmanford

// https://leetcode.com/problems/path-with-maximum-probability

func MaxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	dists := make([]float64, n)
	dists[start] = 1
	for i := 0; i < n-1; i++ {
		for j, e := range edges {
			p := succProb[j]
			relax(dists, e[0], e[1], p)
			relax(dists, e[1], e[0], p)
		}
	}
	return dists[end]
}

func relax(dists []float64, from, to int, p float64) {
	if dists[from]*p > dists[to] {
		dists[to] = dists[from] * p
	}
}
