package dto

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
