package handlers

import (
	"context"
	"log"
	"petprojectmed/dto"
	"petprojectmed/utils"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func GetAllAppointments(conn *pgx.Conn) *[]dto.AppointmentTable {
	query := `SELECT * FROM schedule
	order by id`

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()

	appointments, err := ScanLoadAppointmentEntries(rows)
	utils.CheckErr(err)

	return appointments
}

func GetAppointmentsByID(conn *pgx.Conn, appointmentID []int) *[]dto.AppointmentTable {
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
	log.Println(query)

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()
	appointments, err := ScanLoadAppointmentEntries(rows)
	utils.CheckErr(err)

	return appointments
}

func ScanLoadAppointmentEntries(rows pgx.Rows) (*[]dto.AppointmentTable, error) {
	var appointments []dto.AppointmentTable
	for rows.Next() {
		var appointment dto.AppointmentTable
		err := rows.Scan(&appointment.ID, &appointment.DoctorID, &appointment.DoctorInitials, &appointment.Specialization, &appointment.DateAppointment, &appointment.PatientID, &appointment.PatientInitials)
		if err != nil {
			return nil, err
		}
		appointments = append(appointments, appointment)
	}
	return &appointments, nil
}

func InsertAppointment(conn *pgx.Conn, appointment *dto.InsertAppointmentTable) {
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
	utils.CheckErr(err)
}

func DeleteAppointmentByID(conn *pgx.Conn, appointmentID int) {
	query := `
        DELETE FROM schedule WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": appointmentID,
	}

	_, err := conn.Exec(context.Background(), query, args)
	utils.CheckErr(err)
}

func GetAllDoctorIDsOfScheduleTable(conn *pgx.Conn) []int {
	query := `SELECT doctorid FROM schedule`

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()

	var IDs []int
	for rows.Next() {
		var value int
		err := rows.Scan(&value)
		utils.CheckErr(err)
		IDs = append(IDs, value)
	}

	return IDs
}
