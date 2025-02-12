package handlers

import (
	"log"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"regexp"
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

	if err == nil {
		doctors := ReadDoctorsJsonFile()
		outputData := []dto.Doctor{}

		sort.Ints(paramsID.ID)
		paramsID.ID = utils.RemoveDuplicateInt(paramsID.ID)

		for _, value := range paramsID.ID {
			if value < len(doctors) && value >= 0 {
				outputData = append(outputData, doctors[value])
			} else {
				return c.Status(fiber.StatusBadRequest).SendString("Список содержит несуществующий ID !")
			}
		}
		log.Println(outputData)
		return c.JSON(outputData)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func GetDoctorsListFilter(c *fiber.Ctx) error {
	queryFilters := new(dto.QueryDoctorsListFilter)
	err := c.QueryParser(queryFilters)
	utils.CheckErr(err)

	switch queryFilters.List {
	case "all":
		doctors := ReadDoctorsJsonFile()
		return c.JSON(doctors)
	case "filters":
		//log.Println(queryFilters.Specializations)
		//log.Println(len(queryFilters.Specializations))
		if len(queryFilters.Specializations) != 0 && queryFilters.Specializations[0] != "" {
			doctors := ReadDoctorsJsonFile()
			arrayIndex := []int{}
			for _, valSpecialization := range queryFilters.Specializations {
				arrayIndex = append(arrayIndex, returnIndexOfTargetSpecialization(valSpecialization, &doctors)...)
			}
			sort.Ints(arrayIndex)
			arrayIndex = utils.RemoveDuplicateInt(arrayIndex)
			outputDoctors := []dto.Doctor{}
			for _, value := range arrayIndex {
				outputDoctors = append(outputDoctors, doctors[value])
			}
			return c.JSON(outputDoctors)
		} else {
			return c.Status(fiber.StatusBadRequest).SendString("Пустой список специальностей !")
		}
	case "":
		return c.Status(fiber.StatusBadRequest).SendString("Пустой запрос ! !")
	default:
		return c.Status(fiber.StatusBadRequest).SendString("Неправильный запрос !")
	}
}

func CreateDoctor(c *fiber.Ctx) error {
	newEntry := new(dto.Doctor)

	if err := c.BodyParser(newEntry); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
	}

	isValidData, errorMessage := validateNewJsonDoctors(newEntry)
	if !isValidData {
		return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
	}

	caser := cases.Title(language.Russian)
	newEntry.Specialization = caser.String(utils.TrimSpaces(newEntry.Specialization))
	newEntry.Name = caser.String(newEntry.Name)
	newEntry.Family = caser.String(newEntry.Family)

	doctors := ReadDoctorsJsonFile()
	newEntry.ID = len(doctors)
	doctors = append(doctors, *newEntry)
	WriteDoctorsJsonFile(doctors)
	return c.Status(fiber.StatusCreated).SendString("Запись прошла успешно !")
}

func UpdateDoctor(c *fiber.Ctx) error {
	idValue := c.Params("id")

	intIdValue, err := strconv.Atoi(idValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат ID")
	}

	doctors := ReadDoctorsJsonFile()
	if intIdValue >= 0 && intIdValue < len(doctors) {

		newEntry := new(dto.Doctor)
		if err := c.BodyParser(newEntry); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
		}

		isValidData, errorMessage := validateUpdateJsonDoctors(newEntry)
		if !isValidData {
			return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
		}
		caser := cases.Title(language.Russian)

		if newEntry.Name != "" {
			doctors[intIdValue].Name = caser.String(newEntry.Name)
		}

		if newEntry.Family != "" {
			doctors[intIdValue].Family = caser.String(newEntry.Family)
		}

		if newEntry.Specialization != "" {
			newEntry.Specialization = utils.TrimSpaces(newEntry.Specialization)
			doctors[intIdValue].Specialization = newEntry.Specialization
		}

		if newEntry.DateOfBirth != "" {
			doctors[intIdValue].DateOfBirth = newEntry.DateOfBirth
		}

		if newEntry.Cabinet != 0 {
			doctors[intIdValue].Cabinet = newEntry.Cabinet
		}

		WriteDoctorsJsonFile(doctors)
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

	doctors := ReadDoctorsJsonFile()
	if intIdValue >= 0 && intIdValue < len(doctors) {
		doctor := doctors[intIdValue]
		doctors = append(doctors[:intIdValue], doctors[intIdValue+1:]...)
		for index := range doctors {
			doctors[index].ID = index
		}
		WriteDoctorsJsonFile(doctors)
		//c.Status(fiber.StatusOK).SendString("Запись с id=" + strconv.Itoa(intIdValue) + " успешно удалена !")
		return c.JSON(doctor)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный ID")
	}
}

func validateNewJsonDoctors(newEntry *dto.Doctor) (bool, string) {
	flag := true
	outputString := ""

	/*newEntry.Name*/
	if newEntry.Name == "" {
		flag = false
		outputString += "Пропущено имя !" + "\n"
	}
	if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Name) {
		flag = false
		outputString += "Имя должно содержать только буквы (кириллица) !" + "\n"
	}

	/*newEntry.Family*/
	if newEntry.Family == "" {
		flag = false
		outputString += "Пропущена фамилия !" + "\n"
	}
	if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Family) {
		flag = false
		outputString += "Фамилия должна содержать только буквы (кириллица) !" + "\n"
	}

	/*newEntry.Specialization*/
	if newEntry.Specialization == "" {
		flag = false
		outputString += "Пропущена специализация !" + "\n"
	}
	if regexp.MustCompile(`[^а-яА-Я\s]`).MatchString(newEntry.Specialization) {
		flag = false
		outputString += "Специализация должна содержать только буквы (кириллица) и пробелы !" + "\n"
	}

	/*newEntry.DateOfBirth*/
	val, err := time.Parse(time.DateOnly, newEntry.DateOfBirth)
	if err == nil {
		newEntry.DateOfBirth = val.String()
	} else {
		flag = false
		outputString += "Неверный формат либо пропущена дата рождения !" + "\n"
	}

	/*newEntry.Cabinet*/
	if newEntry.Cabinet <= 0 || newEntry.Cabinet >= 100 {
		flag = false
		outputString += "Неверное значение либо пропущен номер кабинета !" + "\n"
	}

	return flag, outputString
}

func validateUpdateJsonDoctors(newEntry *dto.Doctor) (bool, string) {
	flag := true
	outputString := ""

	/*newEntry.Name*/
	if newEntry.Name != "" {
		if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Name) {
			flag = false
			outputString += "Имя должно содержать только буквы (кириллица) !" + "\n"
		}
	}

	/*newEntry.Family*/
	if newEntry.Family != "" {
		if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Family) {
			flag = false
			outputString += "Фамилия должна содержать только буквы (кириллица) !" + "\n"
		}
	}

	/*newEntry.Specialization*/
	if newEntry.Specialization != "" {
		if regexp.MustCompile(`[^а-яА-Я\s]`).MatchString(newEntry.Specialization) {
			flag = false
			outputString += "Специализация должна содержать только буквы и пробелы !" + "\n"
		}
	}

	/*newEntry.DateOfBirth*/
	if newEntry.DateOfBirth != "" {
		val, err := time.Parse(time.DateOnly, newEntry.DateOfBirth)
		if err == nil {
			newEntry.DateOfBirth = val.String()
		} else {
			flag = false
			outputString += "Неверный формат либо пропущена дата рождения !" + "\n"
		}
	}

	/*newEntry.Cabinet*/
	if newEntry.Cabinet != 0 {
		if newEntry.Cabinet <= 1 || newEntry.Cabinet >= 100 {
			flag = false
			outputString += "Неверное значение !" + "\n"
		}
	}

	return flag, outputString
}

func returnIndexOfTargetSpecialization(specialization string, array *[]dto.Doctor) []int {
	outputArray := []int{}
	caser := cases.Lower(language.Russian)
	for index, value := range *array {
		if caser.String(value.Specialization) == caser.String(specialization) {
			outputArray = append(outputArray, index)
		}
	}
	return outputArray
}
