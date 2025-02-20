package dto

type Doctor struct {
	Name           string `json:"name" `
	Family         string `json:"family"`
	Specialization string `json:"specialization"`
	DateOfBirth    string `json:"dateOfBirth"`
	Cabinet        int    `json:"cabinet"`
}

type Patient struct {
	Name        string `json:"name"`
	Family      string `json:"family"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}

type Appointment struct {
	DoctorID  string `json:"doctorsID"`
	PatientID string `json:"patientsID"`
	Date      string `json:"data"`
	Time      string `json:"time"`
}
