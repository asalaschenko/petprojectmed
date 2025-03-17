package patientservices

import (
	"petprojectmed/storage"
	"time"
)

type ICreatePatient interface {
	Insert(*storage.Patient)
}

type IDeletePatient interface {
	Delete(int)
}

type IGetAllPatient interface {
	Get() *[]storage.Patient
}

type IGetPatientsByID interface {
	Get([]int) *[]storage.Patient
}

type IGetIDs interface {
	Get() *[]int
}

type IGetDateAndID interface {
	Get() *map[int]time.Time
}

type IGetStringAndID interface {
	Get() *map[int]string
}

type IGetString interface {
	Get() *[]string
}

type IUpdateName interface {
	Update(int, string)
}

type IUpdateFamily interface {
	Update(int, string)
}

type IUpdateDateOfBirth interface {
	Update(int, time.Time)
}

type IUpdatePhoneNumber interface {
	Update(int, string)
}

type IUpdateGender interface {
	Update(int, string)
}
