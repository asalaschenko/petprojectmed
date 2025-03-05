package schedule

import (
	"context"
	"petprojectmed/storage"
	"slices"
)

type Appointment struct {
	DoctorID  doctorID  `json:"doctorID" validate:"required,gte=1"`
	PatientID patientID `json:"patientID" validate:"required,gte=1"`
	Date      string    `json:"date" validate:"required,D"`
	Time      string    `json:"time" validate:"required,T"`
}

type QuerySheduleListFilter struct {
	List            string   `query:"list"`
	DoctorID        []int    `query:"doctorID"`
	PatientID       []int    `query:"patientID"`
	DateAppointment []string `query:"dateAppointment"`
}

type patientID int
type doctorID int
type appointmentID int

func (ID *doctorID) verify() bool {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	values := storage.GetIDofDoctors(conn)
	return slices.Contains(*values, int(*ID))
}

func (ID *patientID) verify() bool {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	values := storage.GetIDofPatients(conn)
	return slices.Contains(*values, int(*ID))
}

func (ID *appointmentID) verify() bool {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	values := storage.GetIDofAppointments(conn)
	return slices.Contains(*values, int(*ID))
}

func (val *Appointment) validate() (string, error) {
	err := returnValidator().Struct(val)
	return checkErr(err)
}
