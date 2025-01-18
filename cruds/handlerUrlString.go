package cruds

import (
	"petprojectmed/DAO"
	"strconv"
	"strings"
)

func ConvertToIdValues(array []string) []DAO.ID_DB {
	outputArray := []DAO.ID_DB{}

	for _, idArray := range array {
		splitValues := strings.Fields(idArray)
		for _, idValue := range splitValues {
			value, err := strconv.Atoi(idValue)
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
