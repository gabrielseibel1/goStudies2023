package lc1448

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return countGoodNodes(root, root.Val)
}

func countGoodNodes(root *TreeNode, maxOfPath int) int {
	if root == nil {
		return 0
	}
	good := 0
	if root.Val >= maxOfPath {
		good++
		maxOfPath = root.Val
	}
	good += countGoodNodes(root.Left, maxOfPath)
	good += countGoodNodes(root.Right, maxOfPath)
	return good
}
