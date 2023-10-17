package main

import (
	"fmt"
	"github.com/mjaliz/godsa/exercises"
)

func main() {
	sa := exercises.MergeSortedArrays([]int{12, 14, 25}, []int{10, 20, 32, 40})
	fmt.Println(sa)
}
