package quicksort

func QuickSort(list []int) []int {
	if len(list) < 2 {
		return list
	}

	pivot := list[0]
	var less, greater []int

	for _, v := range list[1:] {
		if v < pivot {
			less = append(less, v)
		} else {
			greater = append(greater, v)
		}
	}

	return append(append(QuickSort(less), pivot), QuickSort(greater)...)
}
