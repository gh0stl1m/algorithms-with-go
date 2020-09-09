package arrays1d

import "fmt"

// Problem: Find longest sequence of 1’s with one flip in an
//          Array composed by 0s and 1s

// FindLongestSequence returns the longest sequence
func FindLongestSequence(arr []uint) uint {
	onesMappedArr := mapOnesValues(arr)

	return getMaxSequence(onesMappedArr)
}

func mapOnesValues(arr []uint) []uint {
	var arrMapped []uint
	var sequenceOnesTotal uint = 0

	for i := 0; i < len(arr); i++ {
		if arr[i] == 1 {
			sequenceOnesTotal++
			if sequenceOnesTotal > 0 && i == len(arr)-1 {
				arrMapped = append(arrMapped, sequenceOnesTotal)
			}
		} else {
			if sequenceOnesTotal > 0 {
				arrMapped = append(arrMapped, sequenceOnesTotal)
			}

			arrMapped = append(arrMapped, arr[i])
			sequenceOnesTotal = 0
		}
	}

	fmt.Println("Arr mapped: ", arrMapped)

	return arrMapped
}

func getMaxSequence(arr []uint) uint {
	var maxSequence uint = 0

	for i, value := range arr {
		currentLength := value

		if (i + 1) < len(arr) {
			currentLength++
		}

		if (i + 2) < len(arr) {
			currentLength += arr[i+2]
		}

		if currentLength > maxSequence {
			maxSequence = currentLength
		}
	}

	return maxSequence
}
