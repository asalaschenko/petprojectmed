package storage

import (
	"context"
	"petprojectmed/common"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type AllDoctors struct {
	conn *pgx.Conn
}

func NewAllDoctors(conn *pgx.Conn) *AllDoctors {
	value := new(AllDoctors)
	value.conn = conn
	return value
}

func (a *AllDoctors) Get() *[]Doctor {
	query := `SELECT * FROM doctors
	order by id`

	rows, err := a.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	doctors, err := LoadDoctorEntries(rows)
	common.CheckErr(err)

	return doctors
}

type DoctorsByID struct {
	conn *pgx.Conn
}

func NewDoctorsByID(conn *pgx.Conn) *DoctorsByID {
	value := new(DoctorsByID)
	value.conn = conn
	return value
}

func (d *DoctorsByID) Get(doctorID []int) *[]Doctor {
	query := `SELECT * FROM doctors WHERE id in `
	str := "("
	for index, value := range doctorID {
		str += strconv.Itoa(value)
		if index == len(doctorID)-1 {
			break
		}
		str += ","
	}
	str += ")"
	str += "\n" + "order by id"
	query += str

	rows, err := d.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()
	doctors, err := LoadDoctorEntries(rows)
	common.CheckErr(err)

	return doctors
}

type IDofDoctors struct {
	conn *pgx.Conn
}

func NewIDofDoctors(conn *pgx.Conn) *IDofDoctors {
	value := new(IDofDoctors)
	value.conn = conn
	return value
}

func (i *IDofDoctors) Get() *[]int {
	query := `
        SELECT id FROM doctors
    `
	rows, err := i.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	IDs, err := LoadInts(rows)
	common.CheckErr(err)

	return IDs
}

type DoctorsIDandSpecializations struct {
	conn *pgx.Conn
}

func NewDoctorsIDandSpecializations(conn *pgx.Conn) *DoctorsIDandSpecializations {
	value := new(DoctorsIDandSpecializations)
	value.conn = conn
	return value
}

func (d *DoctorsIDandSpecializations) Get() *map[int]string {
	query := `
        SELECT id, specialization FROM doctors
    `

	rows, err := d.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandStringField(rows)
	common.CheckErr(err)

	return m
}

type DoctorsIDandDateofBirth struct {
	conn *pgx.Conn
}

func NewDoctorsIDandDateOfBirth(conn *pgx.Conn) *DoctorsIDandDateofBirth {
	value := new(DoctorsIDandDateofBirth)
	value.conn = conn
	return value
}

func (d *DoctorsIDandDateofBirth) Get() *map[int]time.Time {
	query := `
        SELECT id, dateofbirth FROM doctors
    `

	rows, err := d.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandDateField(rows)
	common.CheckErr(err)

	return m
}

type DoctorNameByID struct {
	conn *pgx.Conn
}

func NewDoctorNameByID(conn *pgx.Conn) *DoctorNameByID {
	value := new(DoctorNameByID)
	value.conn = conn
	return value
}

func (d *DoctorNameByID) Update(doctorID int, val string) {
	query := `
	UPDATE doctors
	SET name = @name
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":   doctorID,
		"name": val,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type DoctorFamilyByID struct {
	conn *pgx.Conn
}

func NewDoctorFamilyByID(conn *pgx.Conn) *DoctorFamilyByID {
	value := new(DoctorFamilyByID)
	value.conn = conn
	return value
}

func (d *DoctorFamilyByID) Update(doctorID int, val string) {
	query := `
	UPDATE doctors
	SET family = @family
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":     doctorID,
		"family": val,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type DoctorSpecializationByID struct {
	conn *pgx.Conn
}

func NewDoctorSpecializationByID(conn *pgx.Conn) *DoctorSpecializationByID {
	value := new(DoctorSpecializationByID)
	value.conn = conn
	return value
}

func (d *DoctorSpecializationByID) Update(doctorID int, val string) {
	query := `
	UPDATE doctors
	SET specialization = @specialization
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":             doctorID,
		"specialization": val,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type DoctorCabinetByID struct {
	conn *pgx.Conn
}

func NewDoctorCabinetByID(conn *pgx.Conn) *DoctorCabinetByID {
	value := new(DoctorCabinetByID)
	value.conn = conn
	return value
}

func (d *DoctorCabinetByID) Update(doctorID int, val int) {
	query := `
	UPDATE doctors
	SET cabinet = @cabinet
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":      doctorID,
		"cabinet": val,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type DoctorDateOfBirthByID struct {
	conn *pgx.Conn
}

func NewDoctorDateOfBirthByID(conn *pgx.Conn) *DoctorDateOfBirthByID {
	value := new(DoctorDateOfBirthByID)
	value.conn = conn
	return value
}

func (d *DoctorDateOfBirthByID) Update(doctorID int, val time.Time) {
	query := `
	UPDATE doctors
	SET dateofbirth = @dateofbirth
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":          doctorID,
		"dateofbirth": val,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type DoctorCreate struct {
	conn *pgx.Conn
}

func NewDoctorCreate(conn *pgx.Conn) *DoctorCreate {
	value := new(DoctorCreate)
	value.conn = conn
	return value
}

func (d *DoctorCreate) Insert(doctor *Doctor) {
	query := `
        INSERT INTO doctors (name, family, specialization, cabinet, dateofbirth) VALUES (@name, @family, @specialization, @cabinet, @dateofbirth)
    `
	args := pgx.NamedArgs{
		"name":           doctor.Name,
		"family":         doctor.Family,
		"specialization": doctor.Specialization,
		"cabinet":        doctor.Cabinet,
		"dateofbirth":    doctor.DateOfBirth,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type DoctorByID struct {
	conn *pgx.Conn
}

func NewDoctorByID(conn *pgx.Conn) *DoctorByID {
	value := new(DoctorByID)
	value.conn = conn
	return value
}

func (d *DoctorByID) Delete(doctorID int) {
	query := `
        DELETE FROM doctors WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": doctorID,
	}

	_, err := d.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}
