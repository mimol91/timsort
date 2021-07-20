package timsort

import "testing"

func isEqual(elementsA []T, elementsB []T) bool {
	if len(elementsA) != len(elementsB) {
		return false
	}
	for i, value := range elementsA {
		if value != elementsB[i] {
			return false
		}
	}

	return true
}
func TestSort(t *testing.T) {
	var tests = []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{2, 1, -2, 0, 0, 2, 72, -33}, []int{-33, -2, 0, 0, 1, 2, 2, 72}},
		{[]int{0}, []int{0}},
	}

	for _, test := range tests {
		actual := Sort(test.input)

		if isEqual(actual, test.expected) == false {
			t.Errorf("Sort result for %d is not correct. Expected %d, has %d ", test.input, test.expected, actual)
		}
	}
}
