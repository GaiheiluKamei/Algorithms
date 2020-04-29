package max

import "math"

func Max(l []int) int {
	if len(l) == 0 {
		return math.MinInt64
	}

	if len(l) == 1 {
		return l[0]
	}

	if len(l) == 2 {
		if l[0] > l[1] {
			return l[0]
		}
		return l[1]
	}

	subL := l[1:]

	if l[0] > Max(subL) {
		return l[0]
	}

	return Max(subL)
}
