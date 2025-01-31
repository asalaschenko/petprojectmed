package handlers

import (
	"log"
	"petprojectmed/models"
	"petprojectmed/utils"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CreatePatient(c *fiber.Ctx) error {
	newEntry := new(models.Patient)

	if err := c.BodyParser(newEntry); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
	}

	isValidData, errorMessage := validateNewJsonPatients(newEntry)
	if !isValidData {
		return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
	}

	newEntry.PhoneNumber = utils.TrimSpaces(newEntry.PhoneNumber)
	caser := cases.Title(language.Russian)
	newEntry.Name = caser.String(newEntry.Name)
	newEntry.Family = caser.String(newEntry.Family)
	newEntry.Gender = caser.String(newEntry.Gender)

	patients := ReadPatientsJsonFile()
	newEntry.ID = len(patients)
	patients = append(patients, *newEntry)
	WritePatientsJsonFile(patients)
	return c.Status(fiber.StatusCreated).SendString("Запись прошла успешно !")
}

func UpdatePatient(c *fiber.Ctx) error {
	idValue := c.Params("id")

	intIdValue, err := strconv.Atoi(idValue)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный формат ID")
	}

	patients := ReadPatientsJsonFile()
	if intIdValue >= 0 && intIdValue < len(patients) {

		newEntry := new(models.Patient)
		if err := c.BodyParser(newEntry); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Неверный формат запроса")
		}

		isValidData, errorMessage := validateUpdateJsonPatients(newEntry)
		if !isValidData {
			return c.Status(fiber.StatusBadRequest).SendString(errorMessage)
		}
		caser := cases.Title(language.Russian)

		if newEntry.Name != "" {
			patients[intIdValue].Name = caser.String(newEntry.Name)
		}

		if newEntry.Family != "" {
			patients[intIdValue].Family = caser.String(newEntry.Family)
		}

		if newEntry.Gender != "" {
			patients[intIdValue].Gender = caser.String(newEntry.Gender)
		}

		if newEntry.DateOfBirth != "" {
			patients[intIdValue].DateOfBirth = newEntry.DateOfBirth
		}

		if newEntry.PhoneNumber != "" {
			patients[intIdValue].PhoneNumber = newEntry.PhoneNumber
		}

		WritePatientsJsonFile(patients)
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

	patients := ReadPatientsJsonFile()
	if intIdValue >= 0 && intIdValue < len(patients) {
		patient := patients[intIdValue]
		patients = append(patients[:intIdValue], patients[intIdValue+1:]...)
		for index := range patients {
			patients[index].ID = index
		}
		WritePatientsJsonFile(patients)
		//c.Status(fiber.StatusOK).SendString("Запись с id=" + strconv.Itoa(intIdValue) + " успешно удалена !")
		return c.JSON(patient)
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Неверный ID")
	}
}

func GetPatientsList(c *fiber.Ctx) error {

	patients := ReadPatientsJsonFile()

	m := c.Queries()
	log.Println(m)
	valAll, okAll := m["all"]
	if okAll {
		switch valAll {
		case "":
			return c.JSON(patients)
		case "filter":
			valPhoneNumber, okPhoneNumber := m["phone number"]
			log.Println(valPhoneNumber)
			if okPhoneNumber && valPhoneNumber != "" {
				arrayIndex := returnIndexOfTargetPhoneNumber(valPhoneNumber, &patients)
				log.Println(arrayIndex)
				outputPatients := []models.Patient{}
				for _, value := range arrayIndex {
					outputPatients = append(outputPatients, patients[value])
				}
				return c.JSON(outputPatients)
			}
		}
	}

	structID := new(id)
	flag := c.Query("id", "")
	if flag != "" {
		err := c.QueryParser(structID)
		utils.CheckErr(err)
		outputData := []models.Patient{}
		arrayID := []int{}
		for _, value := range structID.ID {
			id, err := strconv.Atoi(value)
			if err == nil {
				arrayID = append(arrayID, id)
			}
		}
		sort.Ints(arrayID)
		arrayID = utils.RemoveDuplicateInt(arrayID)
		for _, value := range arrayID {
			if value < len(patients) {
				outputData = append(outputData, patients[value])
			}
		}
		return c.JSON(outputData)
	}
	return c.Status(fiber.StatusForbidden).SendString("Пустой или неправильный запрос !")
}

func validateNewJsonPatients(newEntry *models.Patient) (bool, string) {
	flag := true
	outputString := ""
	caser := cases.Lower(language.Russian)

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

	/*newEntry.PhoneNumber*/
	if newEntry.PhoneNumber == "" {
		flag = false
		outputString += "Пропущена специализация !" + "\n"
	}
	if !regexp.MustCompile(`^7[0-9]{10}$`).MatchString(newEntry.PhoneNumber) {
		flag = false
		outputString += "Номер телефона должен иметь следующий формат: 7xxxxxxxxxx !" + "\n"
	}

	/*newEntry.DateOfBirth*/
	val, err := time.Parse(time.DateOnly, newEntry.DateOfBirth)
	if err == nil {
		newEntry.DateOfBirth = val.String()
	} else {
		flag = false
		outputString += "Неверный формат либо пропущена дата рождения !" + "\n"
	}

	/*newEntry.Gender*/
	if caser.String(newEntry.Gender) != "мужской" && caser.String(newEntry.Gender) != "женский" {
		flag = false
		outputString += "Неверное значение пола (\"мужской\" или \"женский\") !" + "\n"
	}

	return flag, outputString
}

func validateUpdateJsonPatients(newEntry *models.Patient) (bool, string) {
	flag := true
	outputString := ""
	caser := cases.Lower(language.Russian)

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

	/*newEntry.PhoneNumber*/
	if newEntry.PhoneNumber != "" {
		if !regexp.MustCompile(`^7[0-9]{10}$`).MatchString(newEntry.PhoneNumber) {
			flag = false
			outputString += "Номер телефона должен иметь следующий формат: 7xxxxxxxxxx !" + "\n"
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
	if newEntry.Gender != "" {
		if caser.String(newEntry.Gender) != "мужской" && caser.String(newEntry.Gender) != "женский" {
			flag = false
			outputString += "Неверное значение пола (\"мужской\" или \"женский\") !" + "\n"
		}
	}

	return flag, outputString
}

func returnIndexOfTargetPhoneNumber(phone_number string, array *[]models.Patient) []int {
	outputArray := []int{}
	for index, value := range *array {
		log.Println(value.PhoneNumber, phone_number)
		if value.PhoneNumber == phone_number {
			outputArray = append(outputArray, index)
		}
	}
	return outputArray
}
