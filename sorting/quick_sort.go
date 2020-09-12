package sorting

import "math/rand"

// QuickSort sorts an array of elements
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left := 0
	right := len(arr) - 1
	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}
