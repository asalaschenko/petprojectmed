package doctors

import (
	"petprojectmed/storage"
)

type IFilterService interface {
	ReturnStatus() string
	GetList(*QueryDoctorsListFilter) *[]storage.Doctor
}

type IListService interface {
	ReturnStatus() string
	GetList(*[]int) *[]storage.Doctor
}

type ICreateService interface {
	ReturnStatus() string
	Create(*Doctor)
}

type IDeleteService interface {
	ReturnStatus() string
	Delete(*int) *storage.Doctor
}

type IUpdateService interface {
	ReturnStatus() string
	Update(int, *DoctorU)
}
