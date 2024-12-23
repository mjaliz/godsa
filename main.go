package main

import (
	"fmt"
	"godsa/linklist"
)

func main() {
	l := linklist.NewLinkList(4)
	l.Append(3)
	l.Append(2)
	l.Insert(1, 5)
	fmt.Println(l)
	tmp := l.Get(1)
	fmt.Println(tmp)
}
