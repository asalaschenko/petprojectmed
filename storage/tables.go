package storage

import "time"

type Doctor struct {
	ID             int
	Name           string
	Family         string
	Specialization string
	Cabinet        int
	DateOfBirth    time.Time
}

type Patient struct {
	ID          int
	Name        string
	Family      string
	DateOfBirth time.Time
	Gender      string
	PhoneNumber string
}

type GetAppointment struct {
	ID              int
	DoctorID        int
	DoctorInitials  string
	Specialization  string
	PatientID       int
	PatientInitials string
	DateAppointment time.Time
}

type InsertAppointment struct {
	DoctorID        int
	PatientID       int
	DateAppointment time.Time
}

func NewDoctor(Name string, Family string, Specialization string, Cabinet int, DateOfBirth time.Time) *Doctor {
	doctor := new(Doctor)
	doctor.Name = Name
	doctor.Family = Family
	doctor.Specialization = Specialization
	doctor.Cabinet = Cabinet
	doctor.DateOfBirth = DateOfBirth
	return doctor
}
