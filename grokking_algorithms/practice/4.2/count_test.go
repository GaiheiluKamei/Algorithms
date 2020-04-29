package count

import "testing"

func TestCount(t *testing.T) {
	l1 := []int{}
	l2 := []int{10}
	l3 := []int{1, 2, 3, 6, 8, 9, 10}

	if Count(l1) != 0 || Count(l2) != 1 || Count(l3) != 7 {
		t.Error("your code is shit.")
	}
}
