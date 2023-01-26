package sortedlisttobst

import "github.com/gabrielseibel1/goStudies2023/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedListToBST(head *list.Node) *TreeNode {
	return bst(head)
}

func bst(head *list.Node) *TreeNode {
	if head == nil { // list is empty
		return nil
	}
	if head.Next == nil { // list has only one element
		return &TreeNode{Val: head.Val}
	}

	// cut left list at mid and find center and right list
	mid := countElements(head) / 2
	left := head
	for i := 1; i < mid; i++ {
		left = left.Next
	}
	center := left.Next
	right := center.Next
	center.Next = nil
	left.Next = nil
	left = head

	return &TreeNode{
		Val:   center.Val,
		Left:  bst(left),
		Right: bst(right),
	}
}

func countElements(head *list.Node) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}
