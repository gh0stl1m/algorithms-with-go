package arrays2d

import (
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	// Arrange
	matrix := [][]uint{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	// Act
	PrintSpiralMatrix(matrix)
}
