package lc417

type position struct {
	r int
	c int
}

type positionSet map[position]struct{}

func (s *positionSet) put(p position) {
	(*s)[p] = struct{}{}
}

func (s *positionSet) has(p position) bool {
	_, ok := (*s)[p]
	return ok
}

func pacificAtlantic(heights [][]int) [][]int {
	a, p := make(positionSet), make(positionSet)

	// init
	for r := range heights {
		for c := range heights[r] {
			if r == 0 {
				p.put(position{r, c})
			}
			if c == 0 {
				p.put(position{r, c})
			}
			if r == len(heights)-1 {
				a.put(position{r, c})
			}
			if c == len(heights[r])-1 {
				a.put(position{r, c})
			}
		}
	}

	// dfs
	addWithDFS(&a, heights)
	addWithDFS(&p, heights)

	// return
	res := make([][]int, 0)
	for k := range a {
		if _, ok := p[k]; ok {
			res = append(res, []int{k.r, k.c})
		}
	}
	return res
}

func addWithDFS(set *positionSet, heights [][]int) {
	seen := make(positionSet)
	for pos := range *set {
		dfs(set, heights, pos, &seen)
	}
}

func dfs(set *positionSet, heights [][]int, pos position, seen *positionSet) {
	if seen.has(pos) {
		return
	}
	seen.put(pos)
	m, n := len(heights), len(heights[0])
	for _, d := range directions(pos, m, n) {
		if (heights[d.r][d.c] >= heights[pos.r][pos.c]) && !seen.has(d) {
			set.put(d)
			dfs(set, heights, d, seen)
		}
	}
}

func directions(pos position, m, n int) []position {
	dirs := make([]position, 0)
	u, r, d, l := pos, pos, pos, pos
	u.r -= 1
	r.c += 1
	d.r += 1
	l.c -= 1
	if u.r >= 0 {
		dirs = append(dirs, u)
	}
	if r.c < n {
		dirs = append(dirs, r)
	}
	if d.r < m {
		dirs = append(dirs, d)
	}
	if l.c >= 0 {
		dirs = append(dirs, l)
	}
	return dirs
}
