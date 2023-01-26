package quadtree

// https://leetcode.com/problems/construct-quad-tree/

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func QuadTree(grid [][]int) *Node {
	if len(grid) == 0 {
		return nil
	}
	return quadTree(grid, 0, len(grid)-1, 0, len(grid)-1)
}

func quadTree(grid [][]int, rs, re, cs, ce int) *Node {
	if (rs == re) && (cs == ce) {
		return &Node{Val: intToBool(grid[rs][cs]), IsLeaf: true}
	}

	rm := (rs + re) / 2
	cm := (cs + ce) / 2
	tl := quadTree(grid, rs, rm, cs, cm)
	tr := quadTree(grid, rs, rm, cm+1, ce)
	bl := quadTree(grid, rm+1, re, cs, cm)
	br := quadTree(grid, rm+1, re, cm+1, ce)

	if v := tl.Val; (tl.IsLeaf && (tl.Val == v)) &&
		(tr.IsLeaf && (tr.Val == v)) &&
		(bl.IsLeaf && (bl.Val == v)) &&
		(br.IsLeaf && (br.Val == v)) {
		tr = nil
		bl = nil
		br = nil
		return tl
	}
	return &Node{TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
}

func intToBool(i int) bool {
	return i == 1
}
