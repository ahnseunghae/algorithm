package insertsort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	var testCases = []struct {
		items    []int
		expected []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{1, 2, 2, 3, 4}, []int{1, 2, 2, 3, 4}},
		{[]int{5, 5, 3, 2, 1}, []int{1, 2, 3, 5, 5}},
	}

	for _, tc := range testCases {
		results, err := sort(tc.items)

		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(results, tc.expected) != true {
			t.Errorf("expected: %v, got: %v", tc.expected, results)
		}
	}

}

func TestEmpty(t *testing.T) {
	items := make([]int, 0)
	_, err := sort(items)

	if err == nil {
		t.Error("Should return error with empty items")
	}
}
