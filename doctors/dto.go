package doctors

import (
	"petprojectmed/common"

	"github.com/go-playground/validator/v10"
)

type Doctor struct {
	Name           string `json:"name" validate:"omitnil,required,NoF,min=1,max=20" validateMessage:"Имя должно содержать только буквы (кириллица) !"`
	Family         string `json:"family" validate:"omitnil,required,NoF,min=1,max=20" validateMessage:"Фамилия должна содержать только буквы (кириллица) !"`
	Specialization string `json:"specialization" validate:"omitnil,required,SP,min=1" validateMessage:"Неверный формат либо пропущена дата рождения !"`
	DateOfBirth    string `json:"dateOfBirth" validate:"omitnil,required,DoB,min=1" validateMessage:"Invalid"`
	Cabinet        int    `json:"cabinet" validate:"gte=1,lte=90" validateMessage:"Неверное значение !"`
}

type QueryDoctorsListFilter struct {
	List            string   `query:"list"`
	DatesOfBirth    []string `query:"datesOfBirth"`
	Specializations []string `query:"specializations"`
}

type ParamsID struct {
	ID doctorsID `params:"id"`
}

type doctorsID []int

func (val *Doctor) validate() (error, string) {
	validate := validator.New()

	err := validate.RegisterValidation("NoF", common.ValidNameFamily)
	common.CheckErr(err)
	err = validate.RegisterValidation("SP", common.ValidSpecialization)
	common.CheckErr(err)
	err = validate.RegisterValidation("DoB", common.ValidDateofBirth)
	common.CheckErr(err)

	err = validate.Struct(val)
	customMessage := ""

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			customMessage += err.Error() // Extract the custom message (simplified)
		}
		return err, customMessage
	} else {
		return nil, OK
	}
}
