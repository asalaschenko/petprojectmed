package common

import (
	"regexp"
	"strings"
	"time"
)

func RemoveDuplicateInt(intSlice []int) []int {
	allKeys := make(map[int]bool)
	list := []int{}
	for _, item := range intSlice {
		if _, value := allKeys[int(item)]; !value {
			allKeys[int(item)] = true
			list = append(list, item)
		}
	}
	return list
}

func TrimSpaces(s string) string {
	s = strings.TrimSpace(s)
	s = regexp.MustCompile(`\s+`).ReplaceAllString(s, " ")
	return s
}

func TransformCharsForDateofBirth(value *string) {
	*value = regexp.MustCompile(`[^0-9]`).ReplaceAllString(*value, " ")
	*value = strings.TrimSpace(*value)
	*value = regexp.MustCompile(`\s+`).ReplaceAllString(*value, "-")
}

func TransformCharsForPhoneNumber(value *string) {
	*value = regexp.MustCompile(`[^0-9]`).ReplaceAllString(*value, " ")
	*value = strings.TrimSpace(*value)
}

func ReturnDateFormat(date string, layout string) time.Time {
	parseDate, _ := time.Parse(layout, date)
	return parseDate
}

func ReturnTimeFormat(time1 string, layout string) time.Time {
	parseTime, _ := time.Parse(layout, time1)
	return parseTime
}

func ReturnDateTimeFormat(date string, time1 string) time.Time {
	parseTime, _ := time.Parse(time.DateTime, date+" "+time1)
	return parseTime
}
