package binarysearch

import "testing"

// TestBinarySearch test the function BinarySearch.
func TestBinarySearch(t *testing.T) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17}

	goal := BinarySearch(list, 16)
	if goal != 15 {
		t.Errorf("Expected 15 but got %d", goal)
	}

	no := BinarySearch(list, 0)
	if no != -1 {
		t.Errorf("Expected -1 but got %d", no)
	}
}
