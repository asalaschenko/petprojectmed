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
	DatesOfBirth []string `query: datesOfBirth`
	PhoneNumbers []string `query:"phoneNumber"`
}

type QuerySheduleListFilter struct {
	List            string   `query:"list"`
	DoctorID        []string `query:"doctorID"`
	PatientID       []string `query:"patientID"`
	DateAppointment []string `query:"dateAppointment"`
}
