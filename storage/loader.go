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

func LoadIDs(rows pgx.Rows) (*[]int, error) {
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
