package schedule

import (
	"context"
	"log"
	"math"
	"petprojectmed/common"
	"petprojectmed/storage"
	"slices"
	"sort"
	"strings"
	"time"
)

func (val *QuerySheduleListFilter) GetList() (*[]storage.GetAppointment, string) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	switch val.List {
	case "all":
		return storage.GetAllAppointment(conn), common.OK
	case "filter":
		resultArray := [][]int{}
		flag := false

		if len(val.DoctorID) != 0 && val.DoctorID[0] != 0 {
			flag = true
			m := storage.GetScheduleIDandDoctorID(conn)

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
			m := storage.GetScheduleIDandPatientID(conn)

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
			m := storage.GetScheduleIDandDateAppointment(conn)

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
			return nil, common.FILTER_EMPTY
		}

		outputArray := common.FindIntersectionOfSetsValues(resultArray)

		if len(outputArray) == 0 {
			return nil, common.NOT_FOUND
		}

		return storage.GetAppointmentsByID(conn, outputArray), common.OK
	default:
		return nil, common.INVALID_REQUEST
	}
}

const doctorAppointmentPeriod = 90
const beginWork, endWork, breakWork = 9, 18, 12
const SATURDAY, SUNDAY = 6, 0

func (val *Appointment) Create() string {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())
	status := []string{}

	if !val.DoctorID.verify() {
		status = append(status, common.INVALID_DOCTOR_ID)
	}
	if !val.PatientID.verify() {
		status = append(status, common.INVALID_PATIENT_ID)
	}

	_, layout := common.CheckAndParseDateValue(val.Date)
	date := common.ReturnDateFormat(val.Date, layout)
	log.Println(date)
	if int(date.Weekday()) == SATURDAY || int(date.Weekday()) == SUNDAY {
		status = append(status, common.DAY_IS_OFF)
	}
	if !time.Now().Before(date) {
		status = append(status, common.EXPIRED_DATE)
	} else {
		diffFloat := math.Abs(float64((time.Since(date) / time.Hour) / 24))
		log.Println(diffFloat)
		if diffFloat > doctorAppointmentPeriod {
			status = append(status, common.TOO_LATE_DATE)
		}
	}

	_, layout = common.CheckAndParseTimeValue(val.Time)
	time1 := common.ReturnTimeFormat(val.Time, layout)
	trunc := time.Hour
	time1 = time1.Truncate(trunc)
	log.Println(time1)
	if time1.Hour() < beginWork || time1.Hour() == breakWork || time1.Hour() > endWork {
		status = append(status, common.NON_WORKING_HOUR)
	}

	dateTime := common.ReturnDateTimeFormat(date.Format(time.DateOnly), time1.Format(time.TimeOnly))
	m := *storage.GetScheduleDoctorIDandDateAppointment(conn)
	log.Println(m)
	if val1, ok := m[int(val.DoctorID)]; ok && slices.Contains(val1, dateTime) {
		status = append(status, common.TIME_IS_BUSY)
	}

	if len(status) != 0 {
		return strings.Join(status, ",")
	} else {
		appointment := storage.NewCreateAppointment(int(val.DoctorID), int(val.PatientID), dateTime)
		storage.InsertAppointment(conn, appointment)

		return common.OK
	}
}

func (ID *appointmentID) Delete() (string, *storage.GetAppointment) {
	conn := storage.GetConnectionDB()
	defer conn.Close(context.Background())

	output := *storage.GetAppointmentsByID(conn, []int{int(*ID)})
	storage.DeleteAppointmentByID(conn, int(*ID))

	return common.OK, &output[0]
}
