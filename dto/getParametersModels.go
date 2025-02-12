package dto

type ParamsListID struct {
	ID []int `params:"id"`
}

type QueryDoctorsListFilter struct {
	List            string   `query:"list"`
	Specializations []string `query:"specializations"`
}

type QueryPatientsListFilter struct {
	List         string   `query:"list"`
	PhoneNumbers []string `query:"phoneNumber"`
}

type QuerySheduleListFilter struct {
	List            string   `query:"list"`
	DoctorID        []string `query:"doctorID"`
	PatientID       []string `query:"patientID"`
	DateAppointment []string `query:"dateAppointment"`
}

type InputJsonAppointment struct {
	DoctorID  string `json:"doctorsID"`
	PatientID string `json:"patientsID"`
	Date      string `json:"data"`
	Time      string `json:"time"`
}
