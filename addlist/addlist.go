package addlist

// https://leetcode.com/problems/add-two-numbers/description/

import (
	"math"

	"github.com/gabrielseibel1/goStudies2023/list"
)

func AddTwoNumbers(l1 *list.Node, l2 *list.Node) *list.Node {
	lenL1 := countDigits(l1)
	lenL2 := countDigits(l2)
	diff := int(math.Abs(float64(lenL1 - lenL2)))
	if diff > 0 {
		zeroes := new(list.Node)
		n := zeroes
		for i := 1; i < diff; i++ {
			n.Next = new(list.Node)
			n = n.Next
		}

		if lenL1 < lenL2 {
			l1 = joinList(l1, zeroes)
		} else if lenL2 < lenL1 {
			l2 = joinList(l2, zeroes)
		}
	}

	l := new(list.Node)
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
			n.Next = new(list.Node)
			n = n.Next
		}
	}

	if carry > 0 {
		n.Next = new(list.Node)
		n.Next.Val = carry
	}

	return l
}

func joinList(l1, l2 *list.Node) *list.Node {
	n := l1
	for n.Next != nil {
		n = n.Next
	}
	n.Next = l2
	return l1
}

func countDigits(l *list.Node) int {
	digits := 0
	for l != nil {
		digits++
		l = l.Next
	}
	return digits
}
