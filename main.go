package main

import (
	"fmt"
	"godsa/linklist"
)

func main() {
	l := linklist.NewLinkList(4)
	l.Append(3)
	fmt.Println(l)
	tmp := l.Pop()
	l.Prepend(5)
	fmt.Println(tmp)
}
