package handlers

import (
	"context"
	"log"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"slices"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func CreateAppointment(c *fiber.Ctx) error {
	queryFilters := new(dto.QuerySheduleListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	switch queryFilters.List {
	case "all":
		appointments := ReadScheduleJsonFile()
		return c.JSON(appointments)
	case "filter":
		resultArray := [][]int{}
		flag := false
		schedule := ReadScheduleJsonFile()

		if len(queryFilters.DoctorID) != 0 && queryFilters.DoctorID[0] != "" {
			flag = true
			arrayIndex := []int{}
			for _, valID := range queryFilters.DoctorID {
				arrayIndex = append(arrayIndex, returnIndexOfTargetDoctorID(valID, &schedule)...)
			}
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			log.Println("queryFilters.DoctorID", arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(queryFilters.PatientID) != 0 && queryFilters.PatientID[0] != "" {
			flag = true
			arrayIndex := []int{}
			for _, valID := range queryFilters.PatientID {
				arrayIndex = append(arrayIndex, returnIndexOfTargetPatientID(valID, &schedule)...)
			}
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			log.Println("queryFilters.PatientID", arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(queryFilters.DateAppointment) != 0 && queryFilters.DateAppointment[0] != "" {
			flag = true
			arrayIndex := []int{}
			for _, valDateTime := range queryFilters.DateAppointment {
				valDateTime += ":00"
				arrayIndex = append(arrayIndex, returnIndexOfTargetDateTime(valDateTime, &schedule)...)
			}
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			log.Println("queryFilters.DateAppointment", arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if !flag {
			return c.Status(fiber.StatusBadRequest).SendString("Список фильтров пуст !")
		}

		outputArray := []int{}
		var count int = 0
		for _, val1 := range resultArray[0] {
			for _, val2 := range resultArray {
				if slices.Contains(val2, val1) {
					count++
				}
			}
			if count == len(resultArray) {
				outputArray = append(outputArray, val1)
			}
			count = 0
		}

		log.Println(outputArray)

		outputAppointments := []Appointment{}
		for _, value := range outputArray {
			outputAppointments = append(outputAppointments, schedule[value])
		}
		return c.JSON(outputAppointments)
	case "":
		return c.Status(fiber.StatusBadRequest).SendString("Пустой запрос ! !")
	default:
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}
