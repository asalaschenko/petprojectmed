package doctors

import (
	"petprojectmed/common"

	"github.com/go-playground/validator/v10"
)

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
	case "SP":
		return "Специализация должна содержать только буквы и пробелы"
	case "DoB":
		return "Неверный формат либо пропущена дата рождения"
	case "lte", "gte":
		if field == "Cabinet" {
			return "Номер кабинета от 1 до 90"
		} else {
			return ""
		}
	case "min", "max":
		if field == "Name" {
			return "Имя содержит от 4 до 20 символов"
		} else if field == "Family" {
			return "Фамилия содержит от 2 до 20 символов"
		} else {
			return ""
		}
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
	err = v.RegisterValidation("SP", common.ValidSpecialization)
	common.CheckErr(err)
	err = v.RegisterValidation("DoB", common.ValidDate)
	common.CheckErr(err)
	return v
}
