package cruds

import (
	"petprojectmed/DAO"
	"strconv"
	"strings"
)

func ConvertToIdArray(array []string) []DAO.ID_DB {
	outputArray := []DAO.ID_DB{}

	for _, str1 := range array {
		splitValues := strings.Fields(str1)
		for _, str2 := range splitValues {
			value, err := strconv.Atoi(str2)
			if err == nil {
				outputArray = append(outputArray, DAO.ID_DB(value))
			}
		}
	}

	return removeDuplicateInt(outputArray)
}

func removeDuplicateInt(intSlice []DAO.ID_DB) []DAO.ID_DB {
	allKeys := make(map[int]bool)
	list := []DAO.ID_DB{}
	for _, item := range intSlice {
		if _, value := allKeys[int(item)]; !value {
			allKeys[int(item)] = true
			list = append(list, item)
		}
	}
	return list
}
