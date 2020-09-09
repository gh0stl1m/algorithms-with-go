package hashmap

import "testing"

func TestFindUniqueNumber(t *testing.T) {
	// Arrange
	var arrayOfValues1 []int = []int{1, 2, 2, 5, 5, 1, 4}
	var arrayOfValues2 []int = []int{1, 2, 2, 4, 5, 1, 4, 6}

	// Act
	uniqueNumber1 := FindUniqueNumber(arrayOfValues1)
	uniqueNumber2 := FindUniqueNumber(arrayOfValues2)

	// Asserts
	if uniqueNumber1 != 4 {
		t.Error("Unexpected value")
	}

	if uniqueNumber2 != 6 {
		t.Error("Unexpected value")
	}
}
