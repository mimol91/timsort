package timsort

import "math"

type Ordered interface {
	type int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, float32, float64, string
}

func binarySearch[T Ordered](data []T, item T, start int, end int) int {
	if start == end {
		if data[start] > item {
			return start
		}
		return start + 1
	}

	if start > end {
		return start
	}

	mid := int(math.Round((float64(start + end)) / 2))
	if data[mid] < item {
		return binarySearch(data, item, mid+1, end)
	}
	if data[mid] > item {
		return binarySearch(data, item, start, mid-1)
	}

	return mid
}

func insertionSort[T Ordered](data []T) []T {
	for i := 1; i < len(data); i++ {
		var result []T
		value := data[i]
		pos := binarySearch(data, value, 0, i-1)

		result = append(result, data[0:pos]...)
		result = append(result, value)
		result = append(result, data[pos:i]...)
		result = append(result, data[i+1:]...)

		data = result
	}

	return data
}

func merge[T Ordered](left []T, right []T) []T {
	var result []T

	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}

	if left[0] < right[0] {
		result = merge(left[1:], right)
		result = append([]T{left[0]}, result...)
	} else {
		result = merge(left, right[1:])
		result = append([]T{right[0]}, result...)
	}

	return result
}

// Sort sort elements using Timsort algorithm.
func Sort[T Ordered](elements []T) []T {
	if len(elements) < 2 {
		return elements
	}

	var runs [][]T
	var sortedRuns [][]T
	newRun := []T{elements[0]}

	for i := 1; i < len(elements); i++ {
		if i == len(elements)-1 {
			newRun = append(newRun, elements[i])
			runs = append(runs, newRun)
			break
		}

		if elements[i] < elements[i-1] {
			runs = append(runs, newRun)
			newRun = []T{elements[i]}
		} else {
			newRun = append(newRun, elements[i])
		}
	}

	for _, item := range runs {
		sortedRuns = append(sortedRuns, insertionSort(item))
	}
	var sortedArray []T
	for _, item := range sortedRuns {
		sortedArray = merge(sortedArray, item)
	}

	return sortedArray
}
