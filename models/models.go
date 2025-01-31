package models

type Doctor struct {
	ID             int    `json:"id"`
	Name           string `json:"name" `
	Family         string `json:"family"`
	Specialization string `json:"specialization"`
	DateOfBirth    string `json:"date of birth"`
	Cabinet        int    `json:"cabinet"`
}

type Patient struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Family      string `json:"family"`
	DateOfBirth string `json:"date of birth,omitempty"`
	Gender      string `json:"gender,omitempty"`
	PhoneNumber string `json:"phone number"`
}
