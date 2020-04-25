package selectionsort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	list := []int{12, 8, 0, -3, 91, 62, 53, 22}
	sorted := SelectionSort(list)

	t.Log(sorted)
	if !reflect.DeepEqual(sorted, []int{-3, 0, 8, 12, 22, 53, 62, 91}) {
		t.Error("incorrect sort result: ", sorted)
	}
}
