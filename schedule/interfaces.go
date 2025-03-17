package schedule

import (
	"petprojectmed/storage"
)

type IFilterService interface {
	ReturnStatus() string
	GetList(*QuerySheduleListFilter) *[]storage.GetAppointment
}

type ICreateService interface {
	ReturnStatus() string
	Create(*Appointment)
}

type IDeleteService interface {
	ReturnStatus() string
	Delete(*int) *storage.GetAppointment
}
