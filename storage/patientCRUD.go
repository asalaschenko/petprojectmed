package storage

import (
	"context"
	"petprojectmed/common"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type AllPatients struct {
	conn *pgx.Conn
}

func NewAllPatients(conn *pgx.Conn) *AllPatients {
	value := new(AllPatients)
	value.conn = conn
	return value
}

func (a *AllPatients) Get() *[]Patient {
	query := `SELECT * FROM patients
	order by id`

	rows, err := a.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	patients, err := LoadPatientEntries(rows)
	common.CheckErr(err)

	return patients
}

type PatientsByID struct {
	conn *pgx.Conn
}

func NewPatientsByID(conn *pgx.Conn) *PatientsByID {
	value := new(PatientsByID)
	value.conn = conn
	return value
}

func (p *PatientsByID) Get(patientID []int) *[]Patient {
	query := `SELECT * FROM patients WHERE id in `
	str := "("
	for index, value := range patientID {
		str += strconv.Itoa(value)
		if index == len(patientID)-1 {
			break
		}
		str += ","
	}
	str += ")"
	str += "\n" + "order by id"
	query += str

	rows, err := p.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()
	patients, err := LoadPatientEntries(rows)
	common.CheckErr(err)

	return patients
}

type PhoneNumberOfPatients struct {
	conn *pgx.Conn
}

func NewPhoneNumberOfPatients(conn *pgx.Conn) *PhoneNumberOfPatients {
	value := new(PhoneNumberOfPatients)
	value.conn = conn
	return value
}

func (p *PhoneNumberOfPatients) Get() *[]string {
	query := `
	SELECT phonenumber FROM patients
`

	rows, err := p.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	PNs, err := LoadStrings(rows)
	common.CheckErr(err)

	return PNs
}

type IDofPatients struct {
	conn *pgx.Conn
}

func NewIDofPatients(conn *pgx.Conn) *IDofPatients {
	value := new(IDofPatients)
	value.conn = conn
	return value
}

func (i *IDofPatients) Get() *[]int {
	query := `
        SELECT id FROM patients
    `

	rows, err := i.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	IDs, err := LoadInts(rows)
	common.CheckErr(err)

	return IDs
}

type PatientIDandPhoneNumbers struct {
	conn *pgx.Conn
}

func NewPatientIDandPhoneNumbers(conn *pgx.Conn) *PatientIDandPhoneNumbers {
	value := new(PatientIDandPhoneNumbers)
	value.conn = conn
	return value
}

func (p *PatientIDandPhoneNumbers) Get() *map[int]string {
	query := `
        SELECT id, phonenumber FROM patients
    `

	rows, err := p.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandStringField(rows)
	common.CheckErr(err)

	return m
}

type PatientIDandDateOfBirth struct {
	conn *pgx.Conn
}

func NewPatientIDandDateOfBirth(conn *pgx.Conn) *PatientIDandDateOfBirth {
	value := new(PatientIDandDateOfBirth)
	value.conn = conn
	return value
}

func (p *PatientIDandDateOfBirth) Get() *map[int]time.Time {
	query := `
        SELECT id, dateofbirth FROM patients
    `

	rows, err := p.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandDateField(rows)
	common.CheckErr(err)

	return m
}

type PatientNameByID struct {
	conn *pgx.Conn
}

func NewPatientNameByID(conn *pgx.Conn) *PatientNameByID {
	value := new(PatientNameByID)
	value.conn = conn
	return value
}

func (p *PatientNameByID) Update(patientID int, val string) {
	query := `
	UPDATE patients
	SET name = @name
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":   patientID,
		"name": val,
	}

	_, err := p.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type PatientFamilyByID struct {
	conn *pgx.Conn
}

func NewPatientFamilyByID(conn *pgx.Conn) *PatientFamilyByID {
	value := new(PatientFamilyByID)
	value.conn = conn
	return value
}

func (p *PatientFamilyByID) Update(patientID int, val string) {
	query := `
	UPDATE patients
	SET family = @family
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":     patientID,
		"family": val,
	}

	_, err := p.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type PatientGenderByID struct {
	conn *pgx.Conn
}

func NewPatientGenderByID(conn *pgx.Conn) *PatientGenderByID {
	value := new(PatientGenderByID)
	value.conn = conn
	return value
}

func (p *PatientGenderByID) Update(patientID int, val string) {
	query := `
	UPDATE patients
	SET gender = @gender
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":     patientID,
		"gender": val,
	}

	_, err := p.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type PatientPhoneNumberByID struct {
	conn *pgx.Conn
}

func NewPatientPhoneNumberByID(conn *pgx.Conn) *PatientPhoneNumberByID {
	value := new(PatientPhoneNumberByID)
	value.conn = conn
	return value
}

func (p *PatientPhoneNumberByID) Update(patientID int, val string) {
	query := `
	UPDATE patients
	SET phonenumber = @phonenumber
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":          patientID,
		"phonenumber": val,
	}

	_, err := p.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type PatientDateOfBirthByID struct {
	conn *pgx.Conn
}

func NewPatientDateOfBirthByID(conn *pgx.Conn) *PatientDateOfBirthByID {
	value := new(PatientDateOfBirthByID)
	value.conn = conn
	return value
}

func (p *PatientDateOfBirthByID) Update(patientID int, val time.Time) {
	query := `
	UPDATE patients
	SET dateofbirth = @dateofbirth
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":          patientID,
		"dateofbirth": val,
	}

	_, err := p.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type PatientCreate struct {
	conn *pgx.Conn
}

func NewPatientCreate(conn *pgx.Conn) *PatientCreate {
	value := new(PatientCreate)
	value.conn = conn
	return value
}

func (p *PatientCreate) Insert(patient *Patient) {
	query := `
        INSERT INTO patients (name, family, dateofbirth, gender, phonenumber) VALUES (@name, @family, @dateofbirth, @gender, @phonenumber)
    `
	args := pgx.NamedArgs{
		"name":        patient.Name,
		"family":      patient.Family,
		"dateofbirth": patient.DateOfBirth,
		"gender":      patient.Gender,
		"phonenumber": patient.PhoneNumber,
	}

	_, err := p.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type PatientByID struct {
	conn *pgx.Conn
}

func NewPatientByID(conn *pgx.Conn) *PatientByID {
	value := new(PatientByID)
	value.conn = conn
	return value
}

func (p *PatientByID) Delete(patientID int) {
	query := `
        DELETE FROM patients WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": patientID,
	}

	_, err := p.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}
