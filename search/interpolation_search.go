package search

// InterpolationSearch that get a number in a sorted array
func InterpolationSearch(sortedArr []int, num int) int {
	lo := 0
	hi := len(sortedArr) - 1
	for lo <= hi {
		mid := lo + ((num - sortedArr[lo]) * (hi - lo) / (sortedArr[hi] - sortedArr[lo]))
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
