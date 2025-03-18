package patients

import (
	"context"
	"petprojectmed/common"
	"petprojectmed/storage"
	"slices"
)

type Patient struct {
	Name        string `json:"name" validate:"required,NoF,min=4,max=20"`
	Family      string `json:"family" validate:"required,NoF,min=2,max=20"`
	DateOfBirth string `json:"dateOfBirth" validate:"required,DoB"`
	Gender      string `json:"gender" validate:"required,G"`
	PhoneNumber string `json:"phoneNumber" validate:"required,PN"`
}

type PatientU struct {
	Name        string `json:"name" validate:"omitempty,NoF,min=4,max=20"`
	Family      string `json:"family" validate:"omitempty,NoF,min=2,max=20"`
	DateOfBirth string `json:"dateOfBirth" validate:"omitempty,DoB"`
	Gender      string `json:"gender" validate:"omitempty,G"`
	PhoneNumber string `json:"phoneNumber" validate:"omitempty,PN"`
}

type QueryPatientsListFilter struct {
	List         string   `query:"list"`
	DatesOfBirth []string `query:"datesOfBirth"`
	PhoneNumbers []string `query:"phoneNumbers"`
}

type ParamsID struct {
	ID []int `params:"id"`
}

func Verify(ID *int) bool {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())
	a := storage.NewIDofPatients(conn)
	values := a.Get()
	return slices.Contains(*values, int(*ID))
}

func (val *Patient) validate() (string, error) {
	common.TransformCharsForDateofBirth(&val.DateOfBirth)
	common.TransformCharsForPhoneNumber(&val.PhoneNumber)
	err := returnValidator().Struct(val)
	return checkErr(err)
}

func (val *PatientU) validate() (string, error) {
	common.TransformCharsForDateofBirth(&val.DateOfBirth)
	common.TransformCharsForPhoneNumber(&val.PhoneNumber)
	err := returnValidator().Struct(val)
	return checkErr(err)
}
