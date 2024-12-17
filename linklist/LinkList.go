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

func (l *LinkList) Pop() *Node {
	if l.Length == 0 {
		return nil
	}
	temp := l.Head
	pre := l.Head
	for temp.Next != nil {
		pre = temp
		temp = temp.Next
	}
	l.Tail = pre
	l.Tail.Next = nil
	l.Length--
	if l.Length == 0 {
		l.Head = nil
		l.Tail = nil
	}
	return temp
}

func (l *LinkList) Prepend(val int) {
	newNode := NewNode(val)
	if l.Length == 0 {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
	l.Length++
}
