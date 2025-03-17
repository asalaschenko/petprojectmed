package scheduleservices

import (
	"petprojectmed/storage"
	"time"
)

type ICreateAppointment interface {
	Insert(*storage.CreateAppointment)
}

type IDeleteAppointment interface {
	Delete(int)
}

type IGetAllAppointment interface {
	Get() *[]storage.GetAppointment
}

type IGetAppointmentsByID interface {
	Get(appointmentID []int) *[]storage.GetAppointment
}

type IGetIDs interface {
	Get() *[]int
}

type IGetIDsAndIDs interface {
	Get() *map[int]int
}

type IGetDateAndID interface {
	Get() *map[int]time.Time
}

type IGetDateByID interface {
	Get(int) *[]time.Time
}
