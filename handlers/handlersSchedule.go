package handlers

import (
	"context"
	"log"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"sort"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAppointments(c *fiber.Ctx) error {
	queryFilters := new(dto.QuerySheduleListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	log.Println(queryFilters.List)

	switch queryFilters.List {
	case "all":
		appointments := GetAllAppointments(conn)
		return c.JSON(appointments)
	case "filter":
		resultArray := [][]int{}
		flag := false
		schedule := GetAllAppointments(conn)

		if len(queryFilters.DoctorID) != 0 && queryFilters.DoctorID[0] != "" {
			flag = true
			arrayIndex := []int{}

			for _, valID := range queryFilters.DoctorID {
				arrayIndex = append(arrayIndex, returnIndexOfTargetDoctorID(valID, schedule)...)
			}

			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(queryFilters.PatientID) != 0 && queryFilters.PatientID[0] != "" {
			flag = true
			arrayIndex := []int{}
			for _, valID := range queryFilters.PatientID {
				arrayIndex = append(arrayIndex, returnIndexOfTargetPatientID(valID, schedule)...)
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
				arrayIndex = append(arrayIndex, returnIndexOfTargetDateTime(valDateTime, schedule)...)
			}
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if !flag {
			return c.Status(fiber.StatusBadRequest).SendString("Список фильтров пуст !")
		}

		outputArray := utils.FindIntersectionOfSetsValues(resultArray)
		if len(outputArray) == 0 {
			return c.Status(fiber.StatusOK).SendString("Нет данных с такими параметрами !")
		}

		outputAppointments := GetAppointmentsByID(conn, outputArray)
		return c.JSON(outputAppointments)
	case "":
		return c.Status(fiber.StatusBadRequest).SendString("Пустой запрос ! !")
	default:
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func CreateAppointment(c *fiber.Ctx) error {

	newEntryInput := new(dto.Appointment)

	if err := c.BodyParser(newEntryInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
	}

	isValidData, errorMessage := validateInputJsonAppointment(newEntryInput)
	if !isValidData {
		return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
	}

	isFreeHour := isFreeHourOfAppointment(newEntryInput)
	if !isFreeHour {
		return c.Status(fiber.StatusBadRequest).SendString("В графике приёма указанного врача данное время занято !")
	}

	newEntryOutput := new(dto.InsertAppointmentTable)
	_, timeValue := utils.CheckParseTimeValue(newEntryInput.Time)
	_, dateValue := utils.CheckParseDateValue(newEntryInput.Date)
	log.Println(dateValue.Format(time.DateOnly))
	log.Println(timeValue.Format(time.TimeOnly))
	dateTimeValue, _ := time.Parse(time.DateTime, dateValue.Format(time.DateOnly)+" "+timeValue.Format(time.TimeOnly))
	log.Println(dateTimeValue)
	trunc := time.Hour
	dateTimeValue = dateTimeValue.Truncate(trunc)
	newEntryOutput.DateAppointment = dateTimeValue
	newEntryOutput.DoctorID, _ = strconv.Atoi(newEntryInput.DoctorID)
	log.Println(newEntryOutput.DoctorID)
	newEntryOutput.PatientID, _ = strconv.Atoi(newEntryInput.PatientID)
	log.Println(newEntryOutput.PatientID)
	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	InsertAppointment(conn, newEntryOutput)
	return c.SendString("Готово")
}

func returnIndexOfTargetDoctorID(ID string, array *[]dto.AppointmentTable) []int {
	outputArray := []int{}
	intID, err := strconv.Atoi(ID)
	if err == nil {
		for _, value := range *array {
			if value.DoctorID == intID {
				outputArray = append(outputArray, value.ID)
			}
		}
	}
	return outputArray
}

func returnIndexOfTargetPatientID(ID string, array *[]dto.AppointmentTable) []int {
	outputArray := []int{}
	intID, err := strconv.Atoi(ID)
	if err == nil {
		for _, value := range *array {
			if value.PatientID == intID {
				outputArray = append(outputArray, value.ID)
			}
		}
	}
	return outputArray
}

func returnIndexOfTargetDateTime(dateTime string, array *[]dto.AppointmentTable) []int {
	outputArray := []int{}
	dateTimeValue, err := time.Parse(time.DateTime, dateTime)

	if err == nil {
		trunc := time.Hour
		dateTimeValue = dateTimeValue.Truncate(trunc)
		for _, value := range *array {
			if value.DateAppointment.Equal(dateTimeValue) {
				outputArray = append(outputArray, value.ID)
			}
		}
	}
	return outputArray
}

func DeleteAppointment(c *fiber.Ctx) error {
	idValue := c.Params("id")

	intIdValue, err := strconv.Atoi(idValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат ID")
	}

	isValidId := validateInputAppointmentID(intIdValue)

	if isValidId {
		conn := GetConnectionDB()
		defer conn.Close(context.Background())

		values := []int{}
		values = append(values, intIdValue)
		entryAppointment := GetAppointmentsByID(conn, values)

		DeleteAppointmentByID(conn, intIdValue)

		return c.JSON(entryAppointment)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный ID")
	}
}
