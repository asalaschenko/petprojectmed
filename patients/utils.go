package patients

import (
	"petprojectmed/common"

	"github.com/go-playground/validator/v10"
)

func returnErrorDescribe(tag string, field string) string {
	switch tag {
	case "NoF":
		if field == "Family" {
			return "Фамилия должна содержать только буквы (кириллица)"
		} else if field == "Name" {
			return "Имя должно содержать только буквы (кириллица)"
		} else {
			return ""
		}
	case "G":
		return "Неверное значение пола (\"мужской\" или \"женский\")"
	case "min", "max":
		if field == "Name" {
			return "Имя содержит от 4 до 20 символов"
		} else if field == "Family" {
			return "Фамилия содержит от 2 до 20 символов"
		} else {
			return ""
		}
	case "DoB":
		return "Неверный формат даты рождения"
	case "PN":
		return "Номер телефона должен иметь следующий формат: 7xxxxxxxxxx"
	case "required":
		return "Отсутствует поле"
	default:
		return ""
	}
}

func returnValidator() *validator.Validate {
	v := validator.New()
	err := v.RegisterValidation("NoF", common.ValidNameFamily)
	common.CheckErr(err)
	err = v.RegisterValidation("G", common.ValidGender)
	common.CheckErr(err)
	err = v.RegisterValidation("DoB", common.ValidDate)
	common.CheckErr(err)
	err = v.RegisterValidation("PN", common.ValidPhoneNumber)
	common.CheckErr(err)
	return v
}

func checkErr(err error) (string, error) {
	customMessage := ""
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			customMessage += err.StructField() + " " + err.Error() + "\n" // Extract the custom message (simplified)
			customMessage += "Descpription: " + returnErrorDescribe(err.Tag(), err.StructField()) + "\n"
		}
		return customMessage, err
	} else {
		return common.OK, nil
	}
}
