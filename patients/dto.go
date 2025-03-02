package patients

type Patient struct {
	Name        string `json:"name"`
	Family      string `json:"family"`
	DateOfBirth string `json:"dateOfBirth"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phoneNumber"`
}
