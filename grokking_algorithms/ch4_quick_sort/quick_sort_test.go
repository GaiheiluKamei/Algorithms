package quicksort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	testCases := []struct {
		desc   string
		input  []int
		output []int
	}{
		{
			desc:   "test zero elements",
			input:  []int{},
			output: []int{},
		},
		{
			desc:   "test one element",
			input:  []int{1},
			output: []int{1},
		},
		{
			desc:   "test two elements",
			input:  []int{10, 9},
			output: []int{9, 10},
		},
		{
			desc:   "test three+ elements",
			input:  []int{9, 5, 6, 8, 3},
			output: []int{3, 5, 6, 8, 9},
		},
	}

	for _, c := range testCases {
		t.Run(c.desc, func(t *testing.T) {
			r := QuickSort(c.input)

			if !reflect.DeepEqual(r, c.output) {
				t.Errorf("expected %v, got %v", c.output, r)
			}
		})
	}
}