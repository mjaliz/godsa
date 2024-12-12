package linklist

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{
		Val: val,
	}
}

type LinkList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func NewLinkList(val int) *LinkList {
	newNode := NewNode(val)
	return &LinkList{
		Head:   newNode,
		Tail:   newNode,
		Length: 1,
	}
}

func (l *LinkList) Append(val int) bool {
	newNode := NewNode(val)
	if l.Head == nil {
		l.Head = newNode
	} else {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
	l.Length++
	return true
}
