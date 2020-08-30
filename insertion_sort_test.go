package insertsort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	items := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5}

	results, err := sort(items)

	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(results, expected) != true {
		t.Error("expected result is different", results)
	}
}

func TestEmpty(t *testing.T) {
	items := make([]int, 0)
	_, err := sort(items)

	if err == nil {
		t.Error("Should return error with empty items")
	}
}
