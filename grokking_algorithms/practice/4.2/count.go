package count

func Count(l []int) int {
	if len(l) == 0 {
		return 0
	}

	return 1 + Count(l[1:])
}
