package utils

import (
	"regexp"
	"strings"
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
