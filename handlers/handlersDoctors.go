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
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GetDoctorsListID(c *fiber.Ctx) error {
	paramsID := new(dto.ParamsListID)
	err := c.ParamsParser(paramsID)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	if err == nil {
		sort.Ints(paramsID.ID)
		paramsID.ID = utils.RemoveDuplicateInt(paramsID.ID)

		outputData := GetDoctorsByID(conn, paramsID.ID)
		return c.JSON(outputData)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func oldGetDoctorsListFilter(c *fiber.Ctx) error {
	log.Println("OK")
	queryFilters := new(dto.QueryDoctorsListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	switch queryFilters.List {
	case "all":
		doctors := GetAllDoctors(conn)
		return c.JSON(doctors)
	case "filter":
		if len(queryFilters.Specializations) != 0 && queryFilters.Specializations[0] != "" {
			arrayIndex := returnIndexOfTargetSpecialization(queryFilters.Specializations)
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			outputDoctors := GetDoctorsByID(conn, arrayIndex)
			return c.JSON(outputDoctors)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString("Пустой список специальностей !")
		}
	case "":
		return c.Status(fiber.StatusBadRequest).SendString("Пустой запрос или неправльный запрос !")
	default:
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func GetDoctorsListFilter(c *fiber.Ctx) error {
	log.Println("OK")
	queryFilters := new(dto.QueryDoctorsListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	switch queryFilters.List {
	case "all":
		doctors := GetAllDoctors(conn)
		return c.JSON(doctors)
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(queryFilters.Specializations) != 0 && queryFilters.Specializations[0] != "" {
			flag = true

			caser := cases.Lower(language.Russian)
			for index, value := range queryFilters.Specializations {
				value = utils.TrimSpaces(value)
				queryFilters.Specializations[index] = caser.String(value)
			}

			arrayIndex := returnIndexOfTargetSpecialization(queryFilters.Specializations)
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

		outputDoctors := GetDoctorsByID(conn, outputArray)
		return c.JSON(outputDoctors)
	case "":
		return c.Status(fiber.StatusBadRequest).SendString("Пустой запрос или неправльный запрос !")
	default:
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func CreateDoctor(c *fiber.Ctx) error {
	inputJson := new(dto.Doctor)

	if err := c.BodyParser(inputJson); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
	}

	isValidData, errorMessage := validateInputJsonDoctorsForCreate(inputJson)
	if !isValidData {
		return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
	}

	caser := cases.Title(language.Russian)
	newEntryDB := new(dto.DoctorTable)
	newEntryDB.Name = caser.String(inputJson.Name)
	newEntryDB.Family = caser.String(inputJson.Family)
	caser = cases.Lower(language.Russian)
	newEntryDB.Specialization = caser.String(utils.TrimSpaces(inputJson.Specialization))
	newEntryDB.Cabinet = inputJson.Cabinet
	newEntryDB.DateOfBirth, _ = time.Parse(time.DateOnly, inputJson.DateOfBirth)

	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	InsertDoctor(conn, newEntryDB)
	return c.Status(fiber.StatusCreated).SendString("Запись прошла успешно !")
}

func UpdateDoctor(c *fiber.Ctx) error {
	idValue := c.Params("id")

	intIdValue, err := strconv.Atoi(idValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат ID")
	}

	isValidId := validateInputDoctorID(intIdValue)

	if isValidId {

		inputJson := new(dto.Doctor)
		if err := c.BodyParser(inputJson); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
		}

		isValidData, errorMessage := validateInputJsonDoctorsForUpdate(inputJson)
		if !isValidData {
			return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
		}

		conn := GetConnectionDB()
		defer conn.Close(context.Background())

		values := []int{}
		values = append(values, intIdValue)
		pointer := GetDoctorsByID(conn, values)
		updateEntryDoctors := *pointer
		updateEntryDoctor := updateEntryDoctors[0]

		caser := cases.Title(language.Russian)

		if inputJson.Name != "" {
			updateEntryDoctor.Name = caser.String(inputJson.Name)
		}

		if inputJson.Family != "" {
			updateEntryDoctor.Family = caser.String(inputJson.Family)
		}

		if inputJson.Specialization != "" {
			inputJson.Specialization = utils.TrimSpaces(inputJson.Specialization)
			updateEntryDoctor.Specialization = inputJson.Specialization
		}

		if inputJson.DateOfBirth != "" {
			value, _ := time.Parse(time.DateOnly, inputJson.DateOfBirth)
			updateEntryDoctor.DateOfBirth = value
		}

		if inputJson.Cabinet != 0 {
			updateEntryDoctor.Cabinet = inputJson.Cabinet
		}

		UpdateDoctorByID(conn, intIdValue, &updateEntryDoctor)
		return c.Status(fiber.StatusOK).SendString("Запись с id=" + strconv.Itoa(intIdValue) + " успешно обновлена !")
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный ID")
	}
}

func DeleteDoctor(c *fiber.Ctx) error {
	idValue := c.Params("id")

	intIdValue, err := strconv.Atoi(idValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат ID")
	}

	isValidId := validateInputDoctorID(intIdValue)

	if isValidId {
		conn := GetConnectionDB()
		defer conn.Close(context.Background())

		values := []int{}
		values = append(values, intIdValue)
		entryDoctor := GetDoctorsByID(conn, values)

		DeleteDoctorByID(conn, intIdValue)

		return c.JSON(entryDoctor)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный ID")
	}
}

func returnIndexOfTargetSpecialization(specializations []string) []int {
	outputArray := []int{}
	conn := GetConnectionDB()
	doctors := GetDoctorsBySpecialization(conn, specializations)

	for _, entry := range *doctors {
		outputArray = append(outputArray, entry.ID)
	}

	return outputArray
}
