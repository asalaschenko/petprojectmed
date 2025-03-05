package storage

import (
	"context"
	"petprojectmed/common"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

func GetAllAppointment(conn *pgx.Conn) *[]GetAppointment {
	query := `
		SELECT * FROM schedule
	`

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	appointments, err := LoadAppointmentEntries(rows)
	common.CheckErr(err)

	return appointments
}

func GetAppointmentsByID(conn *pgx.Conn, appointmentID []int) *[]GetAppointment {
	query := `SELECT * FROM schedule WHERE id in `
	str := "("
	for index, value := range appointmentID {
		str += strconv.Itoa(value)
		if index == len(appointmentID)-1 {
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
	appointments, err := LoadAppointmentEntries(rows)
	common.CheckErr(err)

	return appointments
}

func GetIDofAppointments(conn *pgx.Conn) *[]int {
	query := `
        SELECT id FROM schedule
    `

	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	IDs, err := LoadInts(rows)
	common.CheckErr(err)

	return IDs
}

func GetScheduleIDandDoctorID(conn *pgx.Conn) *map[int]int {
	query := `
		SELECT id, doctorid FROM schedule
	`
	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandIntField(rows)
	common.CheckErr(err)

	return m
}

func GetScheduleIDandPatientID(conn *pgx.Conn) *map[int]int {
	query := `
		SELECT id, patientid FROM schedule
	`
	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandIntField(rows)
	common.CheckErr(err)

	return m
}

func GetScheduleIDandDateAppointment(conn *pgx.Conn) *map[int]time.Time {
	query := `
		SELECT id, dateappointment FROM schedule
	`
	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandDateField(rows)
	common.CheckErr(err)

	return m
}

func GetScheduleDoctorIDandDateAppointment(conn *pgx.Conn) *map[int][]time.Time {
	query := `
		SELECT doctorid, dateappointment FROM schedule
	`
	rows, err := conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadNonUniqueIDandDateField(rows)
	common.CheckErr(err)

	return m
}

func InsertAppointment(conn *pgx.Conn, appointment *CreateAppointment) {
	query := `
        INSERT INTO schedule (doctorid, initialsdoctor, specialization, dateappointment, patientid, initialspatient) 
		VALUES 
		(
		@doctorid, 
		(SELECT d.name || ' ' || d.family AS initialsdoctor from doctors d where d.id=@doctorid), 
		(SELECT specialization FROM doctors d where d.id=@doctorid), 
		@dateappointment, 
		@patientid,
		(SELECT p.name || ' ' || p.family AS initialspatient from patients p where p.id=@patientid)
		)
    `
	args := pgx.NamedArgs{
		"doctorid":        appointment.DoctorID,
		"patientid":       appointment.PatientID,
		"dateappointment": appointment.DateAppointment,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

func DeleteAppointmentByID(conn *pgx.Conn, appointmentID int) {
	query := `
        DELETE FROM schedule WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": appointmentID,
	}

	_, err := conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}
