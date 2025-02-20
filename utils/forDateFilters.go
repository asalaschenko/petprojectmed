package utils

import (
	"slices"
	"time"
)

func CompareYear(date1 time.Time, date2 time.Time) bool {
	return date1.Year() == date2.Year()
}

func CompareMonth(date1 time.Time, date2 time.Time) bool {
	return date1.Month() == date2.Month()
}

func CompareDay(date1 time.Time, date2 time.Time) bool {
	return date1.Day() == date2.Day()
}

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
