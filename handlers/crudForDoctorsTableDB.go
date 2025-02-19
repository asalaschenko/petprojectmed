package handlers

import (
	"context"
	"log"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func GetAllDoctors(conn *pgx.Conn) *[]dto.DoctorTable {
	query := `SELECT * FROM doctors`

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()

	doctors, err := ScanLoadDoctorEntries(rows)
	utils.CheckErr(err)

	return doctors
}

func GetDoctorsByID(conn *pgx.Conn, doctorID []int) *[]dto.DoctorTable {
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

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()
	doctors, err := ScanLoadDoctorEntries(rows)
	utils.CheckErr(err)

	return doctors
}

func GetDoctorsBySpecialization(conn *pgx.Conn, doctorSpecialization []string) *[]dto.DoctorTable {
	query := `SELECT * FROM doctors WHERE specialization in `
	str := "("
	for index, value := range doctorSpecialization {
		str += utils.WrapperSingleQuote(value)
		if index == len(doctorSpecialization)-1 {
			break
		}
		str += ","
	}
	str += ")"
	query += str

	log.Println(query)

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()

	doctors, err := ScanLoadDoctorEntries(rows)
	utils.CheckErr(err)
	return doctors
}

func InsertDoctorByID(conn *pgx.Conn, doctor *dto.DoctorTable) {
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

	_, err := conn.Exec(context.Background(), query, args)
	utils.CheckErr(err)
}

func ScanLoadDoctorEntries(rows pgx.Rows) (*[]dto.DoctorTable, error) {
	var doctors []dto.DoctorTable
	for rows.Next() {
		var doctor dto.DoctorTable
		err := rows.Scan(&doctor.Name, &doctor.Family, &doctor.Specialization, &doctor.Cabinet, &doctor.DateOfBirth, &doctor.ID)
		if err != nil {
			return nil, err
		}
		doctors = append(doctors, doctor)
	}
	return &doctors, nil
}

func DeleteDoctorByID(conn *pgx.Conn, doctorID int) {
	query := `
        DELETE FROM doctors WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": doctorID,
	}

	_, err := conn.Exec(context.Background(), query, args)
	utils.CheckErr(err)
}

func UpdateDoctorByID(conn *pgx.Conn, doctorID int, doctor *dto.DoctorTable) {
	query := `
        UPDATE doctors
        SET name = @name, family = @family, specialization = @specialization, cabinet = @cabinet, dateofbirth = @dateofbirth
        WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id":             doctor.ID,
		"name":           doctor.Name,
		"family":         doctor.Family,
		"specialization": doctor.Specialization,
		"cabinet":        doctor.Cabinet,
		"dateofbirth":    doctor.DateOfBirth,
	}

	_, err := conn.Exec(context.Background(), query, args)
	utils.CheckErr(err)
}
