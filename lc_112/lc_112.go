package lc112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	return dfs(root, targetSum)
}

func dfs(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}

	if root.Left != nil {
		root.Left.Val += root.Val
		if dfs(root.Left, targetSum) {
			return true
		}
	}

	if root.Right != nil {
		root.Right.Val += root.Val
		if dfs(root.Right, targetSum) {
			return true
		}
	}

	return false
}
