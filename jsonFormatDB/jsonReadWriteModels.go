package jsonFormatDB

import "time"

type Doctor struct {
	ID             int    `json:"id"`
	Name           string `json:"name" `
	Family         string `json:"family"`
	Specialization string `json:"specialization"`
	DateOfBirth    string `json:"dateOfBirth"`
	Cabinet        int    `json:"cabinet"`
}

type Patient struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Family      string `json:"family"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}

type Appointment struct {
	ID              int       `json:"id"`
	DoctorID        int       `json:"doctorID"`
	DoctorInitials  string    `json:"doctorInitials"`
	Specialization  string    `json:"specialization"`
	PatientID       int       `json:"patientID"`
	PatientInitials string    `json:"patientInitials"`
	DateAppointment time.Time `json:"dateAppointment"`
}

type InputJsonAppointment struct {
	DoctorID  string `json:"doctorsID"`
	PatientID string `json:"patientsID"`
	Date      string `json:"data"`
	Time      string `json:"time"`
}
