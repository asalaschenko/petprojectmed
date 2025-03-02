package common

import (
	"log"
	"slices"
	"strings"
	"time"
)

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

func compareYear(date1 time.Time, date2 time.Time) bool {
	return date1.Year() == date2.Year()
}

func compareMonth(date1 time.Time, date2 time.Time) bool {
	return date1.Month() == date2.Month()
}

func compareDay(date1 time.Time, date2 time.Time) bool {
	return date1.Day() == date2.Day()
}

func ReturnIndexOfTargetDateOfBirth(dateOfBirth time.Time, m *map[int]time.Time, layout string) []int {
	outputArray := []int{}
	funcArray := [3]func(time.Time, time.Time) bool{compareYear, compareMonth, compareDay}
	form := strings.Split(layout, "-")

	for k, v := range *m {
		flag := true
		for index := range len(form) {
			flag = flag && funcArray[index](dateOfBirth, v)
		}
		if flag {
			outputArray = append(outputArray, k)
		}
	}

	log.Println(outputArray)
	return outputArray
}

func ReturnIndexOfTargetFilterValue(filterValue string, m *map[int]string) []int {
	outputArray := []int{}
	for k, v := range *m {
		if filterValue == v {
			outputArray = append(outputArray, k)
		}
	}
	return outputArray
}
