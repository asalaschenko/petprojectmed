package schedule

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
	case "D":
		return "Формат должен быть гггг-мм-дд/дд-мм-гггг"
	case "T":
		return "Формат должен быть чч/чч:мм/чч:мм:сс"
	case "required":
		return "Отсутствует поле"
	default:
		return ""
	}
}

func returnValidator() *validator.Validate {
	v := validator.New()
	err := v.RegisterValidation("D", common.ValidDate)
	common.CheckErr(err)
	err = v.RegisterValidation("T", common.ValidateTime)
	common.CheckErr(err)
	return v
}
