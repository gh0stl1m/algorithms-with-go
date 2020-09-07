package arrays2d

import "fmt"

// Problem: Spiral matrix

// PrintSpiralMatrix prints a two dimentional array in an spiral way
func PrintSpiralMatrix(matrix [][]uint) {
	firstRow := 0
	lastRow := len(matrix) - 1
	firstColumn := 0
	lastColumn := len(matrix[0]) - 1

	for firstRow < lastRow && firstColumn < lastColumn {
		for i := firstColumn; i <= lastColumn; i++ {
			fmt.Println(matrix[firstRow][i])
		}

		for i := firstRow + 1; i <= lastRow; i++ {
			fmt.Println(matrix[i][firstColumn])
		}

		for i := lastColumn - 1; i >= firstColumn; i-- {
			fmt.Println(matrix[lastRow][i])
		}

		for i := lastRow - 1; i > firstColumn; i-- {
			fmt.Println(matrix[i][firstColumn])
		}

		firstRow++
		lastRow--
		firstColumn++
		lastColumn--
	}
}
