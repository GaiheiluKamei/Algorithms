package max

import (
	"math"
	"testing"
)

func TestMax(t *testing.T) {
	l1 := []int{}
	l2 := []int{10}
	l3 := []int{9, 8}
	l4 := []int{1, 2, 33, 6, 18, 9, 10}

	if Max(l1) != math.MinInt64 || Max(l2) != 10 || Max(l3) != 9 || Max(l4) != 33 {
		t.Error("your code is shit.")
	}
}
