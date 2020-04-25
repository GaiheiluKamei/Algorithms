package selectionsort

func findSmallest(list []int) int {
	smallest := list[0]
	smallestIndex := 0

	for i := range list {
		if list[i] < smallest {
			smallest = list[i]
			smallestIndex = i
		}
	}

	return smallestIndex
}

// SelectionSort sort an array by using the selection sort algorithm.
func SelectionSort(list []int) []int {
	sorted := make([]int, len(list))
	for i := range list {
		smallestIndex := findSmallest(list)
		sorted[i] = list[smallestIndex]
		list = append(list[:smallestIndex], list[smallestIndex+1:]...)
	}
	return sorted
}
