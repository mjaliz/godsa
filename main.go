package main

import (
	"fmt"
	"github.com/mjaliz/godsa/array"
)

func main() {
	arr := array.NewArray()
	fmt.Println(arr)
	arr.Push("Ryan")
	arr.Push("Hana")
	arr.Push("Kian")
	arr.Push("Kaveh")
	arr.Delete(2)
	fmt.Println(arr)
	fmt.Println(arr.Get(0))
}
