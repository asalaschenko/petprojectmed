package scheduleservices

import (
	"math"
	"petprojectmed/common"
	"petprojectmed/doctors"
	"petprojectmed/patients"
	"petprojectmed/schedule"
	"petprojectmed/storage"
	"slices"
	"sort"
	"strings"
	"time"
)

type FilterService struct {
	status           string
	AllAppointment   IGetAllAppointment
	IDandDoctorID    IGetIDsAndIDs
	IDandPatientID   IGetIDsAndIDs
	DateAndID        IGetDateAndID
	AppointmentsByID IGetAppointmentsByID
}

func NewFilterService(AllAppointment IGetAllAppointment, IDandDoctorID IGetIDsAndIDs, IDandPatientID IGetIDsAndIDs, DateAndID IGetDateAndID, AppointmentsByID IGetAppointmentsByID) *FilterService {
	value := new(FilterService)
	value.AllAppointment = AllAppointment
	value.IDandDoctorID = IDandDoctorID
	value.IDandPatientID = IDandPatientID
	value.DateAndID = DateAndID
	value.AppointmentsByID = AppointmentsByID
	return value
}

func (l *FilterService) ReturnStatus() string {
	return l.status
}

func (f *FilterService) GetList(val *schedule.QuerySheduleListFilter) *[]storage.GetAppointment {
	switch val.List {
	case "all":
		f.status = common.OK
		return f.AllAppointment.Get()
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(val.DoctorID) != 0 && val.DoctorID[0] != 0 {
			flag = true
			m := f.IDandDoctorID.Get()

			arrayIndex := []int{}
			for _, value := range val.DoctorID {
				arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetFilterValueInt(value, m)...)
			}

			sort.Ints(arrayIndex)
			arrayIndex = common.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(val.PatientID) != 0 && val.PatientID[0] != 0 {
			flag = true
			m := f.IDandPatientID.Get()

			arrayIndex := []int{}
			for _, value := range val.PatientID {
				arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetFilterValueInt(value, m)...)
			}

			sort.Ints(arrayIndex)
			arrayIndex = common.RemoveDuplicateInt(arrayIndex)
			resultArray = append(resultArray, arrayIndex)
		}

		if len(val.DateAppointment) != 0 && val.DateAppointment[0] != "" {
			flag = true
			m := f.DateAndID.Get()

			arrayIndex := []int{}
			for _, value := range val.DateAppointment {
				if b, layout := common.CheckAndParseDateValueForFilter(value); b {
					date := common.ReturnDateFormat(value, layout)
					arrayIndex = append(arrayIndex, common.ReturnIndexOfTargetDateTimeAppointment(date, m, layout)...)
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
		return f.AppointmentsByID.Get(outputArray)
	default:
		f.status = common.INVALID_REQUEST
		return nil
	}
}

const doctorAppointmentPeriod = 90
const beginWork, endWork, breakWork = 9, 18, 12
const SATURDAY, SUNDAY = 6, 0

type CreateService struct {
	status      string
	Dates       IGetDateByID
	Appointment ICreateAppointment
}

func NewCreateService(Dates IGetDateByID, Appointment ICreateAppointment) *CreateService {
	value := new(CreateService)
	value.Dates = Dates
	value.Appointment = Appointment
	return value
}

func (c *CreateService) ReturnStatus() string {
	return c.status
}

func (c *CreateService) Create(val *schedule.Appointment) {
	status := []string{}

	if !doctors.Verify(&val.DoctorID) {
		status = append(status, common.INVALID_DOCTOR_ID)
	}
	if !patients.Verify(&val.PatientID) {
		status = append(status, common.INVALID_PATIENT_ID)
	}

	_, layout := common.CheckAndParseDateValue(val.Date)
	date := common.ReturnDateFormat(val.Date, layout)

	if int(date.Weekday()) == SATURDAY || int(date.Weekday()) == SUNDAY {
		status = append(status, common.DAY_IS_OFF)
	}
	if !time.Now().Before(date) {
		status = append(status, common.EXPIRED_DATE)
	} else {
		diffFloat := math.Abs(float64((time.Since(date) / time.Hour) / 24))
		if diffFloat > doctorAppointmentPeriod {
			status = append(status, common.TOO_LATE_DATE)
		}
	}

	_, layout = common.CheckAndParseTimeValue(val.Time)
	time1 := common.ReturnTimeFormat(val.Time, layout)
	trunc := time.Hour
	time1 = time1.Truncate(trunc)

	if time1.Hour() < beginWork || time1.Hour() == breakWork || time1.Hour() > endWork {
		status = append(status, common.NON_WORKING_HOUR)
	}

	dateTime := common.ReturnDateTimeFormat(date.Format(time.DateOnly), time1.Format(time.TimeOnly))
	m := *c.Dates.Get(val.DoctorID)

	if slices.Contains(m, dateTime) {
		status = append(status, common.TIME_IS_BUSY)
	}

	if len(status) != 0 {
		c.status = strings.Join(status, ",")
	} else {
		appointment := storage.NewCreateAppointment(int(val.DoctorID), int(val.PatientID), dateTime)
		c.Appointment.Insert(appointment)
		c.status = common.OK
	}
}

type DeleteService struct {
	status           string
	AppointmentsByID IGetAppointmentsByID
	Appointment      IDeleteAppointment
}

func NewDeleteService(AppointmentsByID IGetAppointmentsByID, Appointment IDeleteAppointment) *DeleteService {
	value := new(DeleteService)
	value.AppointmentsByID = AppointmentsByID
	value.Appointment = Appointment
	return value
}

func (u *DeleteService) ReturnStatus() string {
	return u.status
}

func (d *DeleteService) Delete(ID *int) *storage.GetAppointment {
	output := *d.AppointmentsByID.Get([]int{*ID})
	d.Appointment.Delete(*ID)
	d.status = common.OK

	return &output[0]
}
