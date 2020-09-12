package sorting

// BubbleSort sorts an array of integers
func BubbleSort(arr []int) []int {
	var temp int

	for i := range arr {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j-1] > arr[j] {
				temp = arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			}

		}
	}

	return arr
}
