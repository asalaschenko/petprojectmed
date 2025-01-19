package services

import (
	"petprojectmed/dao"
	"strconv"
	"strings"
)

func GetDataOfTables(url string, urlArg []string) string {

	databaseFileName := url
	outputString := ""

	if len(urlArg) == 0 {
		count := dao.GetSizeFileDB(databaseFileName)
		header := dao.GetHeaderFileDB(databaseFileName)
		body := ""
		for i := 0; i < count; i++ {
			index := dao.ID_DB(i)
			body += index.GetRecordFileDB(databaseFileName)
		}
		if body == "" {
			outputString = "<h1><span style=\"font-size:24px;\">Нет данных !</span></h1>"
		} else {
			outputString = wrapperTable(header, body)
		}
	} else {
		idArray := ConvertToIdValues(urlArg)
		header := dao.GetHeaderFileDB(databaseFileName)
		body := ""
		for _, value := range idArray {
			body += value.GetRecordFileDB(databaseFileName)
		}
		if body == "" {
			outputString = "<h1><span style=\"font-size:24px;\">Нет данных !</span></h1>"
		} else {
			outputString = wrapperTable(header, body)
		}
	}
	return outputString
}

func ConvertToIdValues(array []string) []dao.ID_DB {
	outputArray := []dao.ID_DB{}

	for _, idArray := range array {
		splitValues := strings.Fields(idArray)
		for _, idValue := range splitValues {
			value, err := strconv.Atoi(idValue)
			if err == nil {
				outputArray = append(outputArray, dao.ID_DB(value))
			}
		}
	}

	return removeDuplicateInt(outputArray)
}

func removeDuplicateInt(intSlice []dao.ID_DB) []dao.ID_DB {
	allKeys := make(map[int]bool)
	list := []dao.ID_DB{}
	for _, item := range intSlice {
		if _, value := allKeys[int(item)]; !value {
			allKeys[int(item)] = true
			list = append(list, item)
		}
	}
	return list
}
