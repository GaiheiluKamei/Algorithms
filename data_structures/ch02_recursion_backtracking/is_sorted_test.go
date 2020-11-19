package rb

import "testing"

func TestIsSorted(t *testing.T) {
	if !isSorted([]int{1, 2, 3, 4, 5}) {
		t.Error("error")
	}

	if isSorted([]int{1, 2, 4, 3, 5}) {
		t.Error("error")
	}
}