package handlers

import (
	"context"
	"log"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func GetAllPatients(conn *pgx.Conn) *[]dto.PatientTable {
	query := `SELECT * FROM patients`

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()

	patients, err := ScanLoadPatientEntries(rows)
	utils.CheckErr(err)

	return patients
}

func GetPatientsByID(conn *pgx.Conn, patientID []int) *[]dto.PatientTable {
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
	log.Println(query)

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()
	patients, err := ScanLoadPatientEntries(rows)
	utils.CheckErr(err)

	return patients
}

func InsertPatientByID(conn *pgx.Conn, patient *dto.PatientTable) {
	query := `
        INSERT INTO patients (name, family, dateofbirth, gender, phonenumber) VALUES (@name, @family, @dateofbirth, @gender, @phoneNumber)
    `
	args := pgx.NamedArgs{
		"name":        patient.Name,
		"family":      patient.Family,
		"dateofbirth": patient.DateOfBirth,
		"gender":      patient.Gender,
		"phoneNumber": patient.PhoneNumber,
	}

	_, err := conn.Exec(context.Background(), query, args)
	utils.CheckErr(err)
}

func ScanLoadPatientEntries(rows pgx.Rows) (*[]dto.PatientTable, error) {
	var patients []dto.PatientTable
	for rows.Next() {
		var patient dto.PatientTable
		err := rows.Scan(&patient.Name, &patient.Family, &patient.DateOfBirth, &patient.Gender, &patient.PhoneNumber, &patient.ID)
		if err != nil {
			return nil, err
		}
		patients = append(patients, patient)
	}
	return &patients, nil
}

func DeletePatientByID(conn *pgx.Conn, patientID int) {
	query := `
        DELETE FROM patients WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": patientID,
	}

	_, err := conn.Exec(context.Background(), query, args)
	utils.CheckErr(err)
}

func UpdatePatientByID(conn *pgx.Conn, patientID int, patient *dto.PatientTable) {
	query := `
        UPDATE patients
        SET name = @name, family = @family, dateofbirth = @dateofbirth, gender = @gender, phoneNumber = @phoneNumber
        WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id":          patient.ID,
		"name":        patient.Name,
		"family":      patient.Family,
		"dateofbirth": patient.DateOfBirth,
		"gender":      patient.Gender,
		"phoneNumber": patient.PhoneNumber,
	}

	_, err := conn.Exec(context.Background(), query, args)
	utils.CheckErr(err)
}

func GetPatientsByPhoneNumber(conn *pgx.Conn, patientPhoneNumber []string) *[]dto.PatientTable {
	query := `SELECT * FROM patients WHERE phonenumber in `
	str := "("
	for index, value := range patientPhoneNumber {
		str += utils.WrapperSingleQuote(value)
		if index == len(patientPhoneNumber)-1 {
			break
		}
		str += ","
	}
	str += ")"
	str += "\n" + "order by id"
	query += str

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()

	patients, err := ScanLoadPatientEntries(rows)
	utils.CheckErr(err)
	return patients
}
