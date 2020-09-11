package search

import "testing"

func TestInterpolationSearch(t *testing.T) {
	// Arrange
	var values []int = []int{0, 1, 2, 3, 4, 4, 6, 7, 8, 9}

	// Act
	valueFound := InterpolationSearch(values, 3)

	// Asserts
	if valueFound != 3 {
		t.Error("The values must be found")
	}
}
