package doctors

import (
	"context"
	"petprojectmed/common"
	"petprojectmed/storage"
	"slices"
)

type Doctor struct {
	Name           string `json:"name" validate:"required,NoF,min=4,max=20"`
	Family         string `json:"family" validate:"required,NoF,min=2,max=20"`
	Specialization string `json:"specialization" validate:"required,SP,min=4"`
	DateOfBirth    string `json:"dateOfBirth" validate:"required,DoB"`
	Cabinet        int    `json:"cabinet" validate:"required,gte=1,lte=90"`
}

type DoctorU struct {
	Name           string `json:"name" validate:"omitempty,NoF,min=4,max=20"`
	Family         string `json:"family" validate:"omitempty,NoF,min=2,max=20"`
	Specialization string `json:"specialization" validate:"omitempty,SP"`
	DateOfBirth    string `json:"dateOfBirth" validate:"omitempty,DoB"`
	Cabinet        int    `json:"cabinet" validate:"omitempty,gte=1,lte=90"`
}

type QueryDoctorsListFilter struct {
	List            string   `query:"list"`
	DatesOfBirth    []string `query:"datesOfBirth"`
	Specializations []string `query:"specializations"`
}

type ParamsID struct {
	ID doctorsID `params:"id"`
}

type doctorID int

type doctorsID []int

func (ID *doctorID) verify() bool {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	values := storage.GetIDofDoctors(conn)
	return slices.Contains(*values, int(*ID))
}

func (val *Doctor) validate() (error, string) {
	common.TransformCharsForDateofBirth(&val.DateOfBirth)
	err := returnValidator().Struct(val)
	return checkErr(err)
}

func (val *DoctorU) validate() (error, string) {
	common.TransformCharsForDateofBirth(&val.DateOfBirth)
	err := returnValidator().Struct(val)
	return checkErr(err)
}
