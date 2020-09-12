package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// Arrange
	var values []int = []int{1, 0, 2, 6, 8, 3, 5, 4, 7, 9}
	var expectedArr []int = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Act
	sortedArr := QuickSort(values)
	areArrEqual := reflect.DeepEqual(sortedArr, expectedArr)

	// Asserts
	if !areArrEqual {
		t.Error("The array must be sorted")
	}
}
