package handlers

import (
	"context"
	"log"
	"strings"
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

func returnIndexOfTargetDateOfBirth(dateOfBirth time.Time, layout string, s rune) []int {
	outputArray := []int{}
	funcArray := [3]func(time.Time, time.Time) bool{CompareYear, CompareMonth, CompareDay}
	conn := GetConnectionDB()
	defer conn.Close(context.Background())
	form := strings.Split(layout, "-")

	if s == 'p' {
		patients := GetAllPatients(conn)
		for _, value := range *patients {
			flag := true
			for index := range len(form) {
				flag = flag && funcArray[index](dateOfBirth, value.DateOfBirth)
			}
			if flag {
				outputArray = append(outputArray, value.ID)
			}
		}
	} else if s == 'd' {
		doctors := GetAllDoctors(conn)
		for _, value := range *doctors {
			flag := true
			for index := range len(form) {
				flag = flag && funcArray[index](dateOfBirth, value.DateOfBirth)
			}
			if flag {
				outputArray = append(outputArray, value.ID)
			}
		}
	}

	log.Println(outputArray)
	return outputArray
}
