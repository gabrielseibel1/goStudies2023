package lc19

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ptr := head
	if ptr == nil {
		return nil
	}

	// atravessa a lista ate o fim para descobrir o tamanho
	size := 1
	for ptr.Next != nil {
		ptr = ptr.Next
		size++
	}

	if size-n == 0 {
		// remove o primeiro
		head = head.Next
	} else {
		// atravessa a lista ate tamanho-n
		ptr = head
		for i := 0; i < size-n-1; i++ {
			ptr = ptr.Next
		}
		// remove o nodo seguinte
		if ptr.Next != nil {
			ptr.Next = ptr.Next.Next
		}
	}

	return head
}
