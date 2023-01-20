package addlist

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	lenL1 := countDigits(l1)
	lenL2 := countDigits(l2)
	diff := int(math.Abs(float64(lenL1 - lenL2)))
	if diff > 0 {
		zeroes := new(ListNode)
		n := zeroes
		for i := 1; i < diff; i++ {
			n.Next = new(ListNode)
			n = n.Next
		}

		if lenL1 < lenL2 {
			l1 = joinList(l1, zeroes)
		} else if lenL2 < lenL1 {
			l2 = joinList(l2, zeroes)
		}
	}

	l := new(ListNode)
	n := l
	carry := 0
	for l1 != nil {
		v1, v2 := 0, 0

		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		}

		v3 := v1 + v2 + carry
		if v3 > 9 {
			carry = 1
		} else {
			carry = 0
		}
		v3 = v3 % 10

		n.Val = v3
		if l1 != nil {
			n.Next = new(ListNode)
			n = n.Next
		}
	}

	if carry > 0 {
		n.Next = new(ListNode)
		n.Next.Val = carry
	}

	return l
}

func joinList(l1, l2 *ListNode) *ListNode {
	n := l1
	for n.Next != nil {
		n = n.Next
	}
	n.Next = l2
	return l1
}

func countDigits(l *ListNode) int {
	digits := 0
	for l != nil {
		digits++
		l = l.Next
	}
	return digits
}
