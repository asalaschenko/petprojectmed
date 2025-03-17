package doctorservices

import (
	"petprojectmed/storage"
	"time"
)

type ICreateDoctor interface {
	Insert(*storage.Doctor)
}

type IDeleteDoctor interface {
	Delete(int)
}

type IGetAllDoctor interface {
	Get() *[]storage.Doctor
}

type IGetDoctorsByID interface {
	Get([]int) *[]storage.Doctor
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

type IUpdateSpecialization interface {
	Update(int, string)
}

type IUpdateCabinet interface {
	Update(int, int)
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
