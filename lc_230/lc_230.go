package lc230

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// could use a max heap, add all from the tree and select the top one
// in place: go to smallest, backtrack k steps
// get the view from left to right of the tree, of len k, then return last
// go to smallest (always left, found when no left anymore) the next elements are the ones from the right tree or from the parent+parent right tree
func KthSmallest(root *TreeNode, k int) int {
	nodes := leftToRight(root)
	for _, node := range nodes {
		fmt.Println(node)
	}
	return nodes[k-1].Val
}

func leftToRight(node *TreeNode) []*TreeNode {
	nodes := make([]*TreeNode, 0)
	if node.Left != nil {
		nodes = append(nodes, leftToRight(node.Left)...)
	}
	nodes = append(nodes, node)
	if node.Right != nil {
		nodes = append(nodes, leftToRight(node.Right)...)
	}
	return nodes
}
