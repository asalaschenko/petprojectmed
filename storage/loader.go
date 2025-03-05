package storage

import (
	"time"

	"github.com/jackc/pgx/v5"
)

func LoadDoctorEntries(rows pgx.Rows) (*[]Doctor, error) {
	var doctors []Doctor
	for rows.Next() {
		var doctor Doctor
		err := rows.Scan(&doctor.Name, &doctor.Family, &doctor.Specialization, &doctor.Cabinet, &doctor.DateOfBirth, &doctor.ID)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return &doctors, nil
}

func LoadPatientEntries(rows pgx.Rows) (*[]Patient, error) {
	var patients []Patient
	for rows.Next() {
		var patient Patient
		err := rows.Scan(&patient.Name, &patient.Family, &patient.DateOfBirth, &patient.Gender, &patient.PhoneNumber, &patient.ID)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return &patients, nil
}

func LoadAppointmentEntries(rows pgx.Rows) (*[]GetAppointment, error) {
	var appointments []GetAppointment
	for rows.Next() {
		var appointment GetAppointment
		err := rows.Scan(&appointment.ID, &appointment.DoctorID, &appointment.DoctorInitials, &appointment.Specialization, &appointment.DateAppointment, &appointment.PatientID, &appointment.PatientInitials)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	return &appointments, nil
}

func LoadIDandStringField(rows pgx.Rows) (*map[int]string, error) {
	m := make(map[int]string)
	for rows.Next() {
		var k int
		var v string
		err := rows.Scan(&k, &v)
		if err != nil {
			return nil, err
		}
		m[k] = v
	}
	return &m, nil
}

func LoadIDandDateField(rows pgx.Rows) (*map[int]time.Time, error) {
	m := make(map[int]time.Time)
	for rows.Next() {
		var k int
		var v time.Time
		err := rows.Scan(&k, &v)
		if err != nil {
			return nil, err
		}
		m[k] = v
	}
	return &m, nil
}

func LoadNonUniqueIDandDateField(rows pgx.Rows) (*map[int][]time.Time, error) {
	m := map[int][]time.Time{}
	for rows.Next() {
		var k int
		var v time.Time
		err := rows.Scan(&k, &v)
		if err != nil {
			return nil, err
		}
		m[k] = append(m[k], v)
	}
	return &m, nil
}

func LoadIDandIntField(rows pgx.Rows) (*map[int]int, error) {
	m := make(map[int]int)
	for rows.Next() {
		var k int
		var v int
		err := rows.Scan(&k, &v)
		if err != nil {
			return nil, err
		}
		m[k] = v
	}
	return &m, nil
}

func LoadInts(rows pgx.Rows) (*[]int, error) {
	var arrayInt []int
	for rows.Next() {
		var i int
		err := rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		arrayInt = append(arrayInt, i)
	}
	return &arrayInt, nil
}

func LoadStrings(rows pgx.Rows) (*[]string, error) {
	var arrayString []string
	for rows.Next() {
		var i string
		err := rows.Scan(&i)
		if err != nil {
			return nil, err
		}
		arrayString = append(arrayString, i)
	}
	return &arrayString, nil
}
