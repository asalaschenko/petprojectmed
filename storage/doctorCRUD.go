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

	IDs, err := LoadInts(rows)
	common.CheckErr(err)

	return IDs
}

func GetDoctorsIDandSpecializations(conn *pgx.Conn) *map[int]string {
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

func GetDoctorsIDandDateofBirth(conn *pgx.Conn) *map[int]time.Time {
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

func UpdateDoctorNameByID(conn *pgx.Conn, doctorID int, val string) {
	query := `
	UPDATE doctors
	SET name = @name
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":   doctorID,
		"name": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdateDoctorFamilyByID(conn *pgx.Conn, doctorID int, val string) {
	query := `
	UPDATE doctors
	SET family = @family
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":     doctorID,
		"family": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdateDoctorSpecializationByID(conn *pgx.Conn, doctorID int, val string) {
	query := `
	UPDATE doctors
	SET specialization = @specialization
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":             doctorID,
		"specialization": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdateDoctorCabinetByID(conn *pgx.Conn, doctorID int, val int) {
	query := `
	UPDATE doctors
	SET cabinet = @cabinet
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":      doctorID,
		"cabinet": val,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func UpdateDoctorDateOfBirthByID(conn *pgx.Conn, doctorID int, val time.Time) {
	query := `
	UPDATE doctors
	SET dateofbirth = @dateofbirth
	WHERE id = @id
`
	args := pgx.NamedArgs{
		"id":          doctorID,
		"dateofbirth": val,
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
