package patients

import (
	"petprojectmed/storage"
)

type IFilterService interface {
	ReturnStatus() string
	GetList(*QueryPatientsListFilter) *[]storage.Patient
}

type IListService interface {
	ReturnStatus() string
	GetList(*[]int) *[]storage.Patient
}

type ICreateService interface {
	ReturnStatus() string
	Create(*Patient)
}

type IDeleteService interface {
	ReturnStatus() string
	Delete(*int) *storage.Patient
}

type IUpdateService interface {
	ReturnStatus() string
	Update(int, *PatientU)
}
