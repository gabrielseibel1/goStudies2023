package list

type Node struct {
	Val  int
	Next *Node
}

func New(s []int) *Node {
	if len(s) == 0 {
		return nil
	}
	l := new(Node)
	n := l
	for i, v := range s {
		n.Val = v
		if i+1 < len(s) {
			n.Next = new(Node)
			n = n.Next
		}
	}
	return l
}
