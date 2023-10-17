package exercises

func MergeSortedArrays(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	i := 0
	j := 0

	var sa []int
	for i < len(a) || j < len(b) {
		if i < len(a) && a[i] < b[j] {
			sa = append(sa, a[i])
			i++
		} else {
			sa = append(sa, b[j])
			j++
		}
	}

	return sa
}
