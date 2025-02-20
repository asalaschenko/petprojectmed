package dto

import (
	"time"
)

type DoctorTable struct {
	ID             int
	Name           string
	Family         string
	Specialization string
	Cabinet        int
	DateOfBirth    time.Time
}

type PatientTable struct {
	ID          int
	Name        string
	Family      string
	DateOfBirth time.Time
	Gender      string
	PhoneNumber string
}

type AppointmentTable struct {
	ID              int
	DoctorID        int
	DoctorInitials  string
	Specialization  string
	PatientID       int
	PatientInitials string
	DateAppointment time.Time
}

type InsertAppointmentTable struct {
	DoctorID        int
	PatientID       int
	DateAppointment time.Time
}
