package binarysearch

// BinarySearch returns the index of the goal number if it exists in list, -1 or else.
func BinarySearch(list []int, goal int) int {
	low := 0
	high := len(list) - 1

	for high >= low {
		mid := (high + low) / 2
		guess := list[mid]
		if guess == goal {
			return mid
		}

		if guess > goal {
			high = mid - 1
		}

		low = mid + 1
	}

	return -1
}
