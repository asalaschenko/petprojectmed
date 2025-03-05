package storage

import (
	"context"
	"petprojectmed/common"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetAllPatients(conn *pgx.Conn) *[]Patient {
	query := `SELECT * FROM patients
	order by id`

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	patients, err := LoadPatientEntries(rows)
	common.CheckErr(err)

	return patients
}

func GetPatientsByID(conn *pgx.Conn, patientID []int) *[]Patient {
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

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()
	patients, err := LoadPatientEntries(rows)
	common.CheckErr(err)

	return patients
}

func GetPhoneNumberOfPatients(conn *pgx.Conn) *[]string {
	query := `
	SELECT phonenumber FROM patients
`

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	PNs, err := LoadStrings(rows)
	common.CheckErr(err)

	return PNs
}

func GetIDofPatients(conn *pgx.Conn) *[]int {
	query := `
        SELECT id FROM patients
    `

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	IDs, err := LoadInts(rows)
	common.CheckErr(err)

	return IDs
}

func GetPatientIDandPhoneNumbers(conn *pgx.Conn) *map[int]string {
	query := `
        SELECT id, phonenumber FROM patients
    `

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandStringField(rows)
	common.CheckErr(err)

	return m
}

func GetPatientIDandDateOfBirth(conn *pgx.Conn) *map[int]time.Time {
	query := `
        SELECT id, dateofbirth FROM patients
    `

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandDateField(rows)
	common.CheckErr(err)

	return m
}

func UpdatePatientNameByID(conn *pgx.Conn, patientID int, val string) {
	query := `
	UPDATE patients
	SET name = @name
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":   patientID,
		"name": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdatePatientFamilyByID(conn *pgx.Conn, patientID int, val string) {
	query := `
	UPDATE patients
	SET family = @family
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":     patientID,
		"family": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdatePatientGenderByID(conn *pgx.Conn, patientID int, val string) {
	query := `
	UPDATE patients
	SET gender = @gender
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":     patientID,
		"gender": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdatePatientPhoneNumberByID(conn *pgx.Conn, patientID int, val string) {
	query := `
	UPDATE patients
	SET phonenumber = @phonenumber
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":          patientID,
		"phonenumber": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdatePatientDateOfBirthByID(conn *pgx.Conn, patientID int, val time.Time) {
	query := `
	UPDATE patients
	SET dateofbirth = @dateofbirth
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":          patientID,
		"dateofbirth": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func InsertPatient(conn *pgx.Conn, patient *Patient) {
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

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func DeletePatientByID(conn *pgx.Conn, patientID int) {
	query := `
        DELETE FROM patients WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": patientID,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}
