package handlers

import (
	"log"
	"math"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"slices"
	"sort"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func GetAppointments(c *fiber.Ctx) error {
	queryFilters := new(dto.QuerySheduleListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)
	flag := false

	switch queryFilters.List {
	case "all":
		appointments := ReadScheduleJsonFile()
		return c.JSON(appointments)
	case "filter":
		resultArray := [][]int{}
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

		outputAppointments := []dto.Appointment{}
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

func CreateAppointment(c *fiber.Ctx) error {

	newEntryInput := new(dto.InputJsonAppointment)

	if err := c.BodyParser(newEntryInput); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
	}

	isValidData, errorMessage := validateNewJsonAppointment(newEntryInput)
	if !isValidData {
		return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
	}

	isFreeHour := isFreeHourOfAppointment(newEntryInput)
	if !isFreeHour {
		return c.Status(fiber.StatusBadRequest).SendString("В графике приёма указанного врача данное время занято !")
	}

	newEntryOutput := new(dto.Appointment)
	dateValue, _ := time.Parse(time.DateTime, newEntryInput.Date+" "+newEntryInput.Time+":00")
	trunc := time.Hour
	dateValue = dateValue.Truncate(trunc)
	newEntryOutput.DateAppointment = dateValue
	newEntryOutput.DoctorID, _ = strconv.Atoi(newEntryInput.DoctorID)
	newEntryOutput.PatientID, _ = strconv.Atoi(newEntryInput.PatientID)
	doctors := ReadDoctorsJsonFile()
	patients := ReadPatientsJsonFile()
	newEntryOutput.DoctorInitials = doctors[newEntryOutput.DoctorID].Name + " " + doctors[newEntryOutput.DoctorID].Family
	newEntryOutput.PatientInitials = patients[newEntryOutput.PatientID].Name + " " + patients[newEntryOutput.PatientID].Family
	newEntryOutput.Specialization = doctors[newEntryOutput.DoctorID].Specialization
	outputSchedule := ReadScheduleJsonFile()
	newEntryOutput.ID = len(outputSchedule)

	outputSchedule = append(outputSchedule, *newEntryOutput)
	WriteScheduleJsonFile(outputSchedule)
	return c.SendString("Готово")
}

func validateNewJsonAppointment(val *dto.InputJsonAppointment) (bool, string) {
	doctors := ReadDoctorsJsonFile()
	patients := ReadPatientsJsonFile()
	flag := true
	outputString := ""

	/*val.DoctorsID*/
	if val.DoctorID == "" {
		flag = false
		outputString += "Пропущен ID доктора !" + "\n"
	}
	doctroID, err := strconv.Atoi(val.DoctorID)
	if err != nil {
		flag = false
		outputString += "Неверный формат ID доктора !" + "\n"
	}
	if doctroID < 0 || doctroID >= len(doctors) {
		flag = false
		outputString += "Неверный ID доктора !" + "\n"
	}

	/*val.PatientID*/
	if val.PatientID == "" {
		flag = false
		outputString += "Пропущен ID пациента !" + "\n"
	}
	patientID, err := strconv.Atoi(val.PatientID)
	if err != nil {
		flag = false
		outputString += "Неверный формат ID пациента !" + "\n"
	}
	if patientID < 0 || patientID >= len(patients) {
		flag = false
		outputString += "Неверный ID пациента !" + "\n"
	}

	/*val.Date*/
	if val.Date == "" {
		flag = false
		outputString += "Пропущена дата приёма !" + "\n"
	}
	date, err := time.Parse(time.DateOnly, val.Date)
	if err != nil {
		flag = false
		outputString += "Неверный формат даты ! Формат должен быть гггг-мм-дд !" + "\n"
	}
	if int(date.Weekday()) == 6 || int(date.Weekday()) == 0 {
		flag = false
		outputString += "Выходной день !" + "\n"
	}
	t := time.Now()
	if !t.Before(date) {
		flag = false
		outputString += "Выбрана либо текущая дата, либо более более ранняя ! Выберите предстоящую дату !" + "\n"
	} else {
		diff := t.Sub(date) / time.Hour
		diffFloat := math.Abs(float64(diff / 24))
		log.Println(diffFloat)
		if diffFloat > 90 {
			flag = false
			outputString += "Запись открыта только на 90 дней вперёд !" + "\n"
		}
	}

	/*val.Time*/
	if val.Time == "" {
		flag = false
		outputString += "Пропущено время приёма !" + "\n"
	}
	timeValue, err := time.Parse(time.TimeOnly, val.Time+":00")
	if err != nil {
		flag = false
		outputString += "Неверный формат времени ! Формат должен быть чч:мм !" + "\n"
	}
	if timeValue.Hour() < 9 || timeValue.Hour() == 12 || timeValue.Hour() > 18 {
		flag = false
		outputString += "Врач принимает с 9-00 до 19-00, перерыв в 12-00 !" + "\n"
	}

	return flag, outputString
}

func isFreeHourOfAppointment(val *dto.InputJsonAppointment) bool {
	appointments := ReadScheduleJsonFile()

	for _, value := range appointments {
		intDoctorID, _ := strconv.Atoi(val.DoctorID)
		if value.DoctorID == intDoctorID {
			dateValue, _ := time.Parse(time.DateTime, val.Date+" "+val.Time+":00")
			d := time.Hour
			dateValue = dateValue.Truncate(d)
			if value.DateAppointment.Equal(dateValue) {
				return false
			}
		}
	}
	return true
}

func returnIndexOfTargetDoctorID(ID string, array *[]dto.Appointment) []int {
	outputArray := []int{}
	intID, err := strconv.Atoi(ID)
	if err == nil {
		for index, value := range *array {
			if value.DoctorID == intID {
				outputArray = append(outputArray, index)
			}
		}
	}
	return outputArray
}

func returnIndexOfTargetPatientID(ID string, array *[]dto.Appointment) []int {
	outputArray := []int{}
	intID, err := strconv.Atoi(ID)
	if err == nil {
		for index, value := range *array {
			if value.PatientID == intID {
				outputArray = append(outputArray, index)
			}
		}
	}
	return outputArray
}

func returnIndexOfTargetDateTime(dateTime string, array *[]dto.Appointment) []int {
	outputArray := []int{}
	dateTimeValue, err := time.Parse(time.DateTime, dateTime)
	trunc := time.Hour
	dateTimeValue = dateTimeValue.Truncate(trunc)

	if err == nil {
		for index, value := range *array {
			if value.DateAppointment.Equal(dateTimeValue) {
				outputArray = append(outputArray, index)
			}
		}
	}
	return outputArray
}
