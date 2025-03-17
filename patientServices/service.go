package patientservices

import (
	"petprojectmed/common"
	"petprojectmed/patients"
	"petprojectmed/storage"
	"slices"
	"sort"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ListService struct {
	status       string
	Ids          IGetIDs
	PatientsByID IGetPatientsByID
}

func NewListService(Ids IGetIDs, PatientsByID IGetPatientsByID) *ListService {
	value := new(ListService)
	value.Ids = Ids
	value.PatientsByID = PatientsByID
	return value
}

func (l *ListService) ReturnStatus() string {
	return l.status
}

func (l *ListService) GetList(val *[]int) *[]storage.Patient {
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
		return l.PatientsByID.Get(*val)
	}
}

type FilterService struct {
	status           string
	AllPatient       IGetAllPatient
	IDandPhoneNumber IGetStringAndID
	IDandDateBirth   IGetDateAndID
	PatientsByID     IGetPatientsByID
}

func NewFilterService(AllPatient IGetAllPatient, StringAndID IGetStringAndID, DateAndID IGetDateAndID, PatientsByID IGetPatientsByID) *FilterService {
	value := new(FilterService)
	value.AllPatient = AllPatient
	value.IDandPhoneNumber = StringAndID
	value.IDandDateBirth = DateAndID
	value.PatientsByID = PatientsByID
	return value
}

func (l *FilterService) ReturnStatus() string {
	return l.status
}

func (f *FilterService) GetList(val *patients.QueryPatientsListFilter) *[]storage.Patient {
	switch val.List {
	case "all":
		f.status = common.OK
		return f.AllPatient.Get()
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(val.PhoneNumbers) != 0 && val.PhoneNumbers[0] != "" {
			flag = true

			for index, _ := range val.PhoneNumbers {
				common.TransformCharsForPhoneNumber(&val.PhoneNumbers[index])
			}

			m := f.IDandPhoneNumber.Get()

			arrayIndex := []int{}
			for _, value := range val.PhoneNumbers {
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
		return f.PatientsByID.Get(outputArray)
	default:
		f.status = common.INVALID_REQUEST
		return nil
	}
}

type CreateService struct {
	status                string
	PhoneNumberOfPatients IGetString
	Patient               ICreatePatient
}

func NewCreateService(Patient ICreatePatient) *CreateService {
	value := new(CreateService)
	value.Patient = Patient
	return value
}

func (c *CreateService) ReturnStatus() string {
	return c.status
}

func (c *CreateService) Create(val *patients.Patient) {
	entries := c.PhoneNumberOfPatients.Get()
	if slices.Contains(*entries, val.PhoneNumber) {
		c.status = common.INVALID_PHONE_NUMBER
	}

	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)
	_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
	date := common.ReturnDateFormat(val.DateOfBirth, layout)
	patient := storage.NewPatient(caserT.String(val.Name), caserT.String(val.Family), val.PhoneNumber, caserL.String(val.Gender), date)

	c.Patient.Insert(patient)
	c.status = common.OK
}

type UpdateService struct {
	status                string
	PhoneNumberOfPatients IGetString
	Name                  IUpdateName
	Family                IUpdateFamily
	DateOfBirth           IUpdateDateOfBirth
	PhoneNumber           IUpdatePhoneNumber
	Gender                IUpdateGender
}

func NewUpdateService(PhoneNumberOfPatients IGetString, Name IUpdateName, Family IUpdateFamily, DateOfBirth IUpdateDateOfBirth, PhoneNumber IUpdatePhoneNumber, Gender IUpdateGender) *UpdateService {
	value := new(UpdateService)
	value.PhoneNumberOfPatients = PhoneNumberOfPatients
	value.Name = Name
	value.Family = Family
	value.DateOfBirth = DateOfBirth
	value.PhoneNumber = PhoneNumber
	value.Gender = Gender
	return value
}

func (u *UpdateService) ReturnStatus() string {
	return u.status
}

func (u *UpdateService) Update(ID int, val *patients.PatientU) {
	caserT := cases.Title(language.Russian)
	caserL := cases.Lower(language.Russian)

	if val.Name != "" {
		u.Name.Update(ID, caserT.String(val.Name))
	}
	if val.Family != "" {
		u.Family.Update(ID, caserT.String(val.Family))
	}
	if val.Gender != "" {
		u.Gender.Update(ID, caserL.String(val.Gender))
	}
	if val.DateOfBirth != "" {
		_, layout := common.CheckAndParseDateValue(val.DateOfBirth)
		date := common.ReturnDateFormat(val.DateOfBirth, layout)
		u.DateOfBirth.Update(ID, date)
	}
	if val.PhoneNumber != "" {
		entries := u.PhoneNumberOfPatients.Get()
		if slices.Contains(*entries, val.PhoneNumber) {
			u.status = common.INVALID_PHONE_NUMBER
		}

		u.PhoneNumber.Update(ID, val.PhoneNumber)
	}
	u.status = common.OK
}

type DeleteService struct {
	status       string
	PatientsByID IGetPatientsByID
	Patient      IDeletePatient
}

func NewDeleteService(PatientsByID IGetPatientsByID, Patient IDeletePatient) *DeleteService {
	value := new(DeleteService)
	value.PatientsByID = PatientsByID
	value.Patient = Patient
	return value
}

func (u *DeleteService) ReturnStatus() string {
	return u.status
}

func (d *DeleteService) Delete(ID *int) *storage.Patient {
	output := *d.PatientsByID.Get([]int{*ID})
	d.Patient.Delete(*ID)
	d.status = common.OK

	return &output[0]
}
