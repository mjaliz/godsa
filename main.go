package main

import (
	"fmt"
	"godsa/linklist"
)

func main() {
	l := linklist.NewLinkList(4)
	l.Append(3)
	fmt.Println(l)
	tmp := l.PopFirst()
	tmp = l.PopFirst()
	tmp = l.PopFirst()
	fmt.Println(tmp)
}
