package exercises

import (
	"slices"
	"strings"
)

func Reverse(s string) string {
	arr := strings.Split(s, "")
	slices.Reverse(arr)
	return strings.Join(arr, "")
}
