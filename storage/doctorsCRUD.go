package storage

import (
	"context"
	"petprojectmed/common"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetAllDoctors(conn *pgx.Conn) *[]Doctor {
	query := `SELECT * FROM doctors
	order by id`

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	doctors, err := LoadDoctorEntries(rows)
	common.CheckErr(err)

	return doctors
}

func GetDoctorsByID(conn *pgx.Conn, doctorID []int) *[]Doctor {
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
	common.CheckErr(err)
	defer rows.Close()
	doctors, err := LoadDoctorEntries(rows)
	common.CheckErr(err)

	return doctors
}

func GetIDofDoctors(conn *pgx.Conn) *[]int {
	query := `
        SELECT id FROM doctors
    `

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	IDs, err := LoadIDs(rows)
	common.CheckErr(err)

	return IDs
}

func GetIDandSpecializations(conn *pgx.Conn) *map[int]string {
	query := `
        SELECT id, specialization FROM doctors
    `

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandStringField(rows)
	common.CheckErr(err)

	return m
}

func GetIDandDateofBirth(conn *pgx.Conn) *map[int]time.Time {
	query := `
        SELECT id, dateofbirth FROM doctors
    `

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandDateField(rows)
	common.CheckErr(err)

	return m
}

func DeleteDoctorByID(conn *pgx.Conn, doctorID int) {
	query := `
        DELETE FROM doctors WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": doctorID,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdateDoctorByID(conn *pgx.Conn, doctorID int, doctor *Doctor) {
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
	common.CheckErr(err)
}

func InsertDoctor(conn *pgx.Conn, doctor *Doctor) {
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
	common.CheckErr(err)
}
