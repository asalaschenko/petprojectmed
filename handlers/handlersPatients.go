package handlers

import (
	"context"
	"log"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetPatientsListID(c *fiber.Ctx) error {
	paramsID := new(dto.ParamsListID)
	err := c.ParamsParser(paramsID)
	utils.CheckErr(err)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	if err == nil {
		sort.Ints(paramsID.ID)
		paramsID.ID = utils.RemoveDuplicateInt(paramsID.ID)

		outputData := GetPatientsByID(conn, paramsID.ID)
		return c.JSON(outputData)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func OldGetPatientsListFilter(c *fiber.Ctx) error {
	queryFilters := new(dto.QueryPatientsListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	switch queryFilters.List {
	case "all":
		patients := GetAllPatients(conn)
		return c.JSON(patients)
	case "filter":
		if len(queryFilters.PhoneNumbers) != 0 && queryFilters.PhoneNumbers[0] != "" {
			arrayIndex := returnIndexOfTargetPhoneNumber(queryFilters.PhoneNumbers)
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			outputPatients := GetPatientsByID(conn, arrayIndex)
			return c.JSON(outputPatients)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString("Пустой список номеров или неправильный запрос !")
		}
	case "":
		return c.Status(fiber.StatusBadRequest).SendString("Пустой запрос или неправильный запрос !")
	default:
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func GetPatientsListFilter(c *fiber.Ctx) error {
	log.Println("OK")
	queryFilters := new(dto.QueryPatientsListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	switch queryFilters.List {
	case "all":
		patients := GetAllPatients(conn)
		return c.JSON(patients)
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(queryFilters.PhoneNumbers) != 0 && queryFilters.PhoneNumbers[0] != "" {
			flag = true

			for index, value := range queryFilters.PhoneNumbers {
				utils.TransformCharsForPhoneNumber(&value)
				queryFilters.PhoneNumbers[index] = value
			}

			arrayIndex := returnIndexOfTargetPhoneNumber(queryFilters.PhoneNumbers)
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(queryFilters.DatesOfBirth) != 0 && queryFilters.DatesOfBirth[0] != "" {
			flag = true
			arrayIndex := []int{}
			for _, value := range queryFilters.DatesOfBirth {
				utils.TransformCharsForDateofBirth(&value)

				if date, err := time.Parse("2006-01-02", value); err == nil {
					arrayIndex = append(arrayIndex, returnIndexOfTargetDateOfBirth(date, "2006-01-02")...)
				} else if date, err := time.Parse("2006-01", value); err == nil {
					arrayIndex = append(arrayIndex, returnIndexOfTargetDateOfBirth(date, "2006-01")...)
				} else if date, err := time.Parse("2006", value); err == nil {
					arrayIndex = append(arrayIndex, returnIndexOfTargetDateOfBirth(date, "2006")...)
				} else if date, err := time.Parse("01-2006", value); err == nil {
					arrayIndex = append(arrayIndex, returnIndexOfTargetDateOfBirth(date, "01-2006")...)
				} else if date, err := time.Parse("02-01-2006", value); err == nil {
					arrayIndex = append(arrayIndex, returnIndexOfTargetDateOfBirth(date, "02-01-2006")...)
				}
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

		outputPatients := GetPatientsByID(conn, outputArray)
		return c.JSON(outputPatients)
	case "":
		return c.Status(fiber.StatusBadRequest).SendString("Пустой запрос или неправильный запрос !")
	default:
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func CreatePatient(c *fiber.Ctx) error {
	inputJson := new(dto.Patient)

	if err := c.BodyParser(inputJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
	}

	isValidData, errorMessage := validateInputJsonPatientsForCreate(inputJson)
	if !isValidData {
		return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
	}

	caser := cases.Title(language.Russian)
	newEntryDB := new(dto.PatientTable)
	newEntryDB.Name = caser.String(inputJson.Name)
	newEntryDB.Family = caser.String(inputJson.Family)
	newEntryDB.DateOfBirth, _ = time.Parse(time.DateOnly, inputJson.DateOfBirth)
	caser = cases.Lower(language.Russian)
	newEntryDB.Gender = caser.String(inputJson.Gender)
	newEntryDB.PhoneNumber = inputJson.PhoneNumber

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	InsertPatient(conn, newEntryDB)
	return c.Status(fiber.StatusCreated).SendString("Запись прошла успешно !")
}

func UpdatePatient(c *fiber.Ctx) error {
	idValue := c.Params("id")

	intIdValue, err := strconv.Atoi(idValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат ID")
	}

	isValidId := validateInputPatientID(intIdValue)

	if isValidId {

		inputJson := new(dto.Patient)
		if err := c.BodyParser(inputJson); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
		}

		isValidData, errorMessage := validateInputJsonPatientsForUpdate(inputJson)
		if !isValidData {
			return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
		}

		conn := GetConnectionDB()
		defer conn.Close(context.Background())

		values := []int{}
		values = append(values, intIdValue)
		pointer := GetPatientsByID(conn, values)
		updateEntryPatients := *pointer
		updateEntryPatient := updateEntryPatients[0]

		caser := cases.Title(language.Russian)
		if inputJson.Name != "" {
			updateEntryPatient.Name = caser.String(inputJson.Name)
		}

		if inputJson.Family != "" {
			updateEntryPatient.Family = caser.String(inputJson.Family)
		}

		if inputJson.DateOfBirth != "" {
			value, _ := time.Parse(time.DateOnly, inputJson.DateOfBirth)
			updateEntryPatient.DateOfBirth = value
		}

		caser = cases.Lower(language.Russian)
		if inputJson.Gender != "" {
			inputJson.Gender = caser.String(inputJson.Gender)
			updateEntryPatient.Gender = inputJson.Gender
		}

		if inputJson.PhoneNumber != "" {
			updateEntryPatient.PhoneNumber = inputJson.PhoneNumber
		}

		UpdatePatientByID(conn, intIdValue, &updateEntryPatient)
		return c.Status(fiber.StatusOK).SendString("Запись с id=" + strconv.Itoa(intIdValue) + " успешно обновлена !")
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный ID")
	}
}

func DeletePatient(c *fiber.Ctx) error {
	idValue := c.Params("id")

	intIdValue, err := strconv.Atoi(idValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат ID")
	}

	isValidId := validateInputPatientID(intIdValue)

	if isValidId {
		conn := GetConnectionDB()
		defer conn.Close(context.Background())

		values := []int{}
		values = append(values, intIdValue)
		entryPatient := GetPatientsByID(conn, values)

		DeletePatientByID(conn, intIdValue)

		return c.JSON(entryPatient)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный ID")
	}
}

func returnIndexOfTargetPhoneNumber(phoneNunmbers []string) []int {
	outputArray := []int{}
	conn := GetConnectionDB()
	defer conn.Close(context.Background())
	patients := GetPatientsByPhoneNumber(conn, phoneNunmbers)

	for _, entry := range *patients {
		outputArray = append(outputArray, entry.ID)
	}

	return outputArray
}

func returnIndexOfTargetDateOfBirth(dateOfBirth time.Time, layout string) []int {
	outputArray := []int{}
	funcArray := [3]func(time.Time, time.Time) bool{utils.CompareYear, utils.CompareMonth, utils.CompareDay}
	conn := GetConnectionDB()
	defer conn.Close(context.Background())
	patients := GetAllPatients(conn)
	form := strings.Split(layout, "-")

	for _, value := range *patients {
		flag := true
		for index := range len(form) {
			flag = flag && funcArray[index](dateOfBirth, value.DateOfBirth)
		}
		if flag {
			outputArray = append(outputArray, value.ID)
		}
	}

	log.Println(outputArray)
	return outputArray
}
