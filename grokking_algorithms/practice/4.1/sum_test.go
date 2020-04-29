package sum

import "testing"

func TestSum(t *testing.T) {
	l1 := []int{}
	l2 := []int{10}
	l3 := []int{1, 2, 3, 6, 8, 9, 10}

	if Sum(l1) != 0 || Sum(l2) != 10 || Sum(l3) != 39 {
		t.Error("your code is shit.")
	}
}
