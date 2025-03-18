package schedule

import (
	"context"
	"petprojectmed/storage"
	"slices"
)

type Appointment struct {
	DoctorID  int    `json:"doctorID" validate:"required,gte=1"`
	PatientID int    `json:"patientID" validate:"required,gte=1"`
	Date      string `json:"date" validate:"required,D"`
	Time      string `json:"time" validate:"required,T"`
}

type QuerySheduleListFilter struct {
	List            string   `query:"list"`
	DoctorID        []int    `query:"doctorID"`
	PatientID       []int    `query:"patientID"`
	DateAppointment []string `query:"dateAppointment"`
}

func verify(ID *int) bool {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())
	a := storage.NewIDofAppointments(conn)
	values := a.Get()
	return slices.Contains(*values, int(*ID))
}

func (val *Appointment) validate() (string, error) {
	err := returnValidator().Struct(val)
	return checkErr(err)
}
