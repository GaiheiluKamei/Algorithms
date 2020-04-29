package sum

func Sum(l []int) int {
	if len(l) == 0 {
		return 0
	}

	return l[0] + Sum(l[1:])
}
