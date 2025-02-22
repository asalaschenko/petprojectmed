package utils

import "slices"

func FindIntersectionOfSetsValues(arrays [][]int) []int {
	outputArray := []int{}
	var count int = 0
	for _, val1 := range arrays[0] {
		for _, val2 := range arrays {
			if slices.Contains(val2, val1) {
				count++
			}
		}
		if count == len(arrays) {
			outputArray = append(outputArray, val1)
		}
		count = 0
	}
	return outputArray
}
