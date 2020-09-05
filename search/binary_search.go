package search

// BinarySearch that get a number in a sorted array
func BinarySearch(sortedArr []int, num int) int {
	lo := 0
	hi := len(sortedArr)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if num < sortedArr[mid] {
			hi = mid - 1
		} else if num > sortedArr[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
