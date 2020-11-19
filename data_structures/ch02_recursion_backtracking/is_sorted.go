package rb

/*
Problem-2: Given an array, check whether the array is in sorted order with recursion.
*/

func isSorted(a []int) bool {
	n := len(a)

	if n == 1 {
		return true
	}

	if a[n-1] < a[n-2] {
		return false
	}

	return isSorted(a[:n-1])
}

