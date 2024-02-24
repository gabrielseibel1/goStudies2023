package lc199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	view := make([]int, 0)
	level := make([]*TreeNode, 1)
	level[0] = root

	for len(level) != 0 {
		view = append(view, level[len(level)-1].Val)
		newLevel := make([]*TreeNode, 0, len(level)*2)

		for _, node := range level {
			if node.Left != nil {
				newLevel = append(newLevel, node.Left)
			}
			if node.Right != nil {
				newLevel = append(newLevel, node.Right)
			}
		}

		level = newLevel
	}

	return view
}
