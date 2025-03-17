package doctorservices

import (
	"log"
	"petprojectmed/common"
	"petprojectmed/doctors"
	"petprojectmed/storage"
	"slices"
	"sort"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ListService struct {
	status      string
	Ids         IGetIDs
	DoctorsByID IGetDoctorsByID
}

func NewListService(Ids IGetIDs, DoctorsByID IGetDoctorsByID) *ListService {
	value := new(ListService)
	value.Ids = Ids
	value.DoctorsByID = DoctorsByID
	return value
}

func (l *ListService) ReturnStatus() string {
	return l.status
}

func (l *ListService) GetList(val *[]int) *[]storage.Doctor {
	sort.Ints(*val)
	*val = common.RemoveDuplicateInt(*val)

	arrayIndex := []int{}
	arrayIDs := l.Ids.Get()

	for _, value := range *val {
		if slices.Contains(*arrayIDs, value) {
			arrayIndex = append(arrayIndex, value)
		}
	}

	if len(arrayIndex) == 0 {
		l.status = common.NOT_FOUND
		return nil
	} else {
		l.status = common.OK
		return l.DoctorsByID.Get(*val)
	}
}

type FilterService struct {
	status              string
	AllDoctor           IGetAllDoctor
	IDandSpecialization IGetStringAndID
	IDandDateBirth      IGetDateAndID
	DoctorsByID         IGetDoctorsByID
}

func NewFilterService(AllDoctor IGetAllDoctor, StringAndID IGetStringAndID, DateAndID IGetDateAndID, DoctorsByID IGetDoctorsByID) *FilterService {
	value := new(FilterService)
	value.AllDoctor = AllDoctor
	value.IDandSpecialization = StringAndID
	value.IDandDateBirth = DateAndID
	value.DoctorsByID = DoctorsByID
	return value
}

func (l *FilterService) ReturnStatus() string {
	return l.status
}

func (f *FilterService) GetList(val *doctors.QueryDoctorsListFilter) *[]storage.Doctor {
	switch val.List {
	case "all":
		f.status = common.OK
		return f.AllDoctor.Get()
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(val.Specializations) != 0 && val.Specializations[0] != "" {
			flag = true

			caser := cases.Lower(language.Russian)
			for index, value := range val.Specializations {
				value = common.TrimSpaces(value)
				val.Specializations[index] = caser.String(value)
			}

			m := f.IDandSpecialization.Get()

			arrayIndex := []int{}
			for _, value := range val.Specializations {
				arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetFilterValueString(value, m)...)
			}
			sort.Ints(arrayIndex)
			arrayIndex = common.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(val.DatesOfBirth) != 0 && val.DatesOfBirth[0] != "" {
			flag = true

			for index, _ := range val.DatesOfBirth {
				common.TransformCharsForDateofBirth(&val.DatesOfBirth[index])
			}

			m := f.IDandDateBirth.Get()

			arrayIndex := []int{}
			for _, value := range val.DatesOfBirth {
				if b, layout := common.CheckAndParseDateValueForFilter(value); b {
					date := common.ReturnDateFormat(value, layout)
					arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetDateOfBirth(date, m, layout)...)
				}
			}

			log.Println(arrayIndex)

			sort.Ints(arrayIndex)
			arrayIndex = common.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if !flag {
			f.status = common.FILTER_EMPTY
			return nil
		}

		outputArray := common.FindIntersectionOfSetsValues(resultArray)

		if len(outputArray) == 0 {
			f.status = common.NOT_FOUND
			return nil
		}

		f.status = common.OK
		return f.DoctorsByID.Get(outputArray)
	default:
		f.status = common.INVALID_REQUEST
		return nil
	}
}

type CreateService struct {
	status string
	Doctor ICreateDoctor
}

func NewCreateService(Doctor ICreateDoctor) *CreateService {
	value := new(CreateService)
	value.Doctor = Doctor
	return value
}

func (c *CreateService) ReturnStatus() string {
	return c.status
}

func (c *CreateService) Create(val *doctors.Doctor) {
	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)
	_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
	date := common.ReturnDateFormat(val.DateOfBirth, layout)
	doctor := storage.NewDoctor(caserT.String(val.Name), caserT.String(val.Family), caserL.String(common.TrimSpaces(val.Specialization)), val.Cabinet, date)

	c.Doctor.Insert(doctor)
	c.status = common.OK
}

type UpdateService struct {
	status         string
	Name           IUpdateName
	Family         IUpdateFamily
	DateOfBirth    IUpdateDateOfBirth
	Specialization IUpdateSpecialization
	Cabinet        IUpdateCabinet
}

func NewUpdateService(Name IUpdateName, Family IUpdateFamily, DateOfBirth IUpdateDateOfBirth, Specialization IUpdateSpecialization, Cabinet IUpdateCabinet) *UpdateService {
	value := new(UpdateService)
	value.Name = Name
	value.Family = Family
	value.DateOfBirth = DateOfBirth
	value.Specialization = Specialization
	value.Cabinet = Cabinet
	return value
}

func (u *UpdateService) ReturnStatus() string {
	return u.status
}

func (u *UpdateService) Update(ID int, val *doctors.DoctorU) {
	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)

	if val.Name != "" {
		u.Name.Update(ID, caserT.String(val.Name))
	}
	if val.Family != "" {
		u.Family.Update(ID, caserT.String(val.Family))
	}
	if val.Specialization != "" {
		u.Specialization.Update(ID, caserL.String(common.TrimSpaces(val.Specialization)))
	}
	if val.DateOfBirth != "" {
		_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
		date := common.ReturnDateFormat(val.DateOfBirth, layout)
		u.DateOfBirth.Update(ID, date)
	}
	if val.Cabinet != 0 {
		u.Cabinet.Update(ID, val.Cabinet)
	}
	u.status = common.OK
}

type DeleteService struct {
	status      string
	DoctorsByID IGetDoctorsByID
	Doctor      IDeleteDoctor
}

func NewDeleteService(DoctorsByID IGetDoctorsByID, Doctor IDeleteDoctor) *DeleteService {
	value := new(DeleteService)
	value.DoctorsByID = DoctorsByID
	value.Doctor = Doctor
	return value
}

func (u *DeleteService) ReturnStatus() string {
	return u.status
}

func (d *DeleteService) Delete(ID *int) *storage.Doctor {
	output := *d.DoctorsByID.Get([]int{int(*ID)})
	d.Doctor.Delete(int(*ID))
	d.status = common.OK
	return &output[0]
}
