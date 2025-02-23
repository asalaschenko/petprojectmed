package handlers

import (
	"context"
	"math"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"regexp"
	"strconv"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func validateInputJsonDoctorsForCreate(newEntry *dto.Doctor) (bool, string) {
	flag := true
	outputString := ""

	/*newEntry.Name*/
	if newEntry.Name == "" {
		flag = false
		outputString += "Пропущено имя !" + "\n"
	} else if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Name) {
		flag = false
		outputString += "Имя должно содержать только буквы (кириллица) !" + "\n"
	}

	/*newEntry.Family*/
	if newEntry.Family == "" {
		flag = false
		outputString += "Пропущена фамилия !" + "\n"
	} else if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Family) {
		flag = false
		outputString += "Фамилия должна содержать только буквы (кириллица) !" + "\n"
	}

	/*newEntry.Specialization*/
	if newEntry.Specialization == "" {
		flag = false
		outputString += "Пропущена специализация !" + "\n"
	} else if regexp.MustCompile(`[^а-яА-Я\s]`).MatchString(newEntry.Specialization) {
		flag = false
		outputString += "Специализация должна содержать только буквы (кириллица) и пробелы !" + "\n"
	}

	/*newEntry.DateOfBirth*/
	valid, _ := utils.CheckParseDateValue(newEntry.DateOfBirth)
	if !valid {
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

func validateInputJsonDoctorsForUpdate(newEntry *dto.Doctor) (bool, string) {
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

	// /*newEntry.DateOfBirth*/
	if newEntry.DateOfBirth != "" {
		valid, _ := utils.CheckParseDateValue(newEntry.DateOfBirth)
		if !valid {
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

func validateInputJsonPatientsForCreate(newEntry *dto.Patient) (bool, string) {
	flag := true
	outputString := ""
	caser := cases.Lower(language.Russian)

	/*newEntry.Name*/
	if newEntry.Name == "" {
		flag = false
		outputString += "Пропущено имя !" + "\n"
	} else if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Name) {
		flag = false
		outputString += "Имя должно содержать только буквы (кириллица) !" + "\n"
	}

	/*newEntry.Family*/
	if newEntry.Family == "" {
		flag = false
		outputString += "Пропущена фамилия !" + "\n"
	} else if regexp.MustCompile(`[^а-яА-Я]`).MatchString(newEntry.Family) {
		flag = false
		outputString += "Фамилия должна содержать только буквы (кириллица) !" + "\n"
	}

	/*newEntry.PhoneNumber*/
	if newEntry.PhoneNumber == "" {
		flag = false
		outputString += "Пропущен номер телефона !" + "\n"
	} else if !regexp.MustCompile(`^7[0-9]{10}$`).MatchString(newEntry.PhoneNumber) {
		flag = false
		outputString += "Номер телефона должен иметь следующий формат: 7xxxxxxxxxx !" + "\n"
	} else if isTherePhoneNumberInOtherPatients(newEntry.PhoneNumber) {
		flag = false
		outputString += "Вы ввели уже существующий номер телефона !" + "\n"
	}

	/*newEntry.DateOfBirth*/
	valid, _ := utils.CheckParseDateValue(newEntry.DateOfBirth)
	if !valid {
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

func validateInputJsonPatientsForUpdate(newEntry *dto.Patient) (bool, string) {
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
		if isTherePhoneNumberInOtherPatients(newEntry.PhoneNumber) {
			flag = false
			outputString += "Вы ввели уже существующий номер телефона !"
		}
	}

	/*newEntry.DateOfBirth*/
	if newEntry.DateOfBirth != "" {
		valid, _ := utils.CheckParseDateValue(newEntry.DateOfBirth)
		if !valid {
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

func isTherePhoneNumberInOtherPatients(phoneNumber string) bool {
	conn := GetConnectionDB()
	defer conn.Close(context.Background())
	entries := GetAllPatients(conn)
	for _, value := range *entries {
		if value.PhoneNumber == phoneNumber {
			return true
		}
	}
	return false
}

func validateInputJsonAppointment(val *dto.Appointment) (bool, string) {
	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	doctors := GetAllDoctors(conn)
	patients := GetAllPatients(conn)
	flag := true
	outputString := ""

	/*val.DoctorsID*/
	if val.DoctorID == "" {
		flag = false
		outputString += "Пропущен ID доктора !" + "\n"
	} else {
		doctroID, err := strconv.Atoi(val.DoctorID)
		if err != nil {
			flag = false
			outputString += "Неверный формат ID доктора !" + "\n"
		} else if doctroID < 0 || doctroID >= len(*doctors) {
			flag = false
			outputString += "Неверный ID доктора !" + "\n"
		}
	}

	/*val.PatientID*/
	if val.PatientID == "" {
		flag = false
		outputString += "Пропущен ID пациента !" + "\n"
	} else {
		patientID, err := strconv.Atoi(val.PatientID)
		if err != nil {
			flag = false
			outputString += "Неверный формат ID пациента !" + "\n"
		} else if patientID < 0 || patientID >= len(*patients) {
			flag = false
			outputString += "Неверный ID пациента !" + "\n"
		}
	}

	/*val.Date*/
	if val.Date == "" {
		flag = false
		outputString += "Пропущена дата приёма !" + "\n"
	} else {
		valid, date := utils.CheckParseDateValue(val.Date)
		if !valid {
			flag = false
			outputString += "Неверный формат даты ! Формат должен быть гггг-мм-дд !" + "\n"
		} else if int(date.Weekday()) == 6 || int(date.Weekday()) == 0 {
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
			//log.Println(diffFloat)
			if diffFloat > 90 {
				flag = false
				outputString += "Запись открыта только на 90 дней вперёд !" + "\n"
			}
		}
	}

	/*val.Time*/
	if val.Time == "" {
		flag = false
		outputString += "Пропущено время приёма !" + "\n"
	} else {
		valid, timeValue := utils.CheckParseTimeValue(val.Time)
		if !valid {
			flag = false
			outputString += "Неверный формат времени !" + "\n"
		} else if timeValue.Hour() < 9 || timeValue.Hour() == 12 || timeValue.Hour() > 18 {
			flag = false
			outputString += "Врач принимает с 9-00 до 19-00, перерыв в 12-00 !" + "\n"
		}
	}

	return flag, outputString
}

func isFreeHourOfAppointment(val *dto.Appointment) bool {
	conn := GetConnectionDB()
	defer conn.Close(context.Background())

	intDoctorID, _ := strconv.Atoi(val.DoctorID)
	appointments := GetAllAppointments(conn)

	for _, value := range *appointments {
		if value.DoctorID == intDoctorID {
			_, timeValue := utils.CheckParseTimeValue(val.Time)
			dateValue, _ := time.Parse(time.DateTime, val.Date+" "+timeValue.Format(time.TimeOnly))
			d := time.Hour
			dateValue = dateValue.Truncate(d)
			if value.DateAppointment.Equal(dateValue) {
				return false
			}
		}
	}
	return true
}

func validateInputAppointmentID(appointmentID int) bool {
	conn := GetConnectionDB()
	defer conn.Close(context.Background())
	entries := GetAllAppointments(conn)

	for _, value := range *entries {
		if value.ID == appointmentID {
			return true
		}
	}

	return false
}

func validateInputDoctorID(doctorID int) bool {
	conn := GetConnectionDB()
	defer conn.Close(context.Background())
	entries := GetAllDoctors(conn)

	for _, value := range *entries {
		if value.ID == doctorID {
			return true
		}
	}

	return false
}

func validateInputPatientID(patientID int) bool {
	conn := GetConnectionDB()
	defer conn.Close(context.Background())
	entries := GetAllPatients(conn)

	for _, value := range *entries {
		if value.ID == patientID {
			return true
		}
	}

	return false
}
