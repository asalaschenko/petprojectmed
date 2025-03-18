package storage

import (
	"context"
	"petprojectmed/common"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

type AllAppointment struct {
	conn *pgx.Conn
}

func NewAllAppointment(conn *pgx.Conn) *AllAppointment {
	value := new(AllAppointment)
	value.conn = conn
	return value
}

func (a *AllAppointment) Get() *[]GetAppointment {
	query := `
		SELECT * FROM schedule
	`

	rows, err := a.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	appointments, err := LoadAppointmentEntries(rows)
	common.CheckErr(err)

	return appointments
}

type AppointmentsByID struct {
	conn *pgx.Conn
}

func NewAppointmentsByID(conn *pgx.Conn) *AppointmentsByID {
	value := new(AppointmentsByID)
	value.conn = conn
	return value
}

func (a *AppointmentsByID) Get(appointmentID []int) *[]GetAppointment {
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

	rows, err := a.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()
	appointments, err := LoadAppointmentEntries(rows)
	common.CheckErr(err)

	return appointments
}

type IDofAppointments struct {
	conn *pgx.Conn
}

func NewIDofAppointments(conn *pgx.Conn) *IDofAppointments {
	value := new(IDofAppointments)
	value.conn = conn
	return value
}

func (i *IDofAppointments) Get() *[]int {
	query := `
        SELECT id FROM schedule
    `

	rows, err := i.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	IDs, err := LoadInts(rows)
	common.CheckErr(err)

	return IDs
}

type ScheduleIDandDoctorID struct {
	conn *pgx.Conn
}

func NewScheduleIDandDoctorID(conn *pgx.Conn) *ScheduleIDandDoctorID {
	value := new(ScheduleIDandDoctorID)
	value.conn = conn
	return value
}

func (s *ScheduleIDandDoctorID) Get() *map[int]int {
	query := `
		SELECT id, doctorid FROM schedule
	`
	rows, err := s.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandIntField(rows)
	common.CheckErr(err)

	return m
}

type ScheduleIDandPatientID struct {
	conn *pgx.Conn
}

func NewScheduleIDandPatientID(conn *pgx.Conn) *ScheduleIDandPatientID {
	value := new(ScheduleIDandPatientID)
	value.conn = conn
	return value
}

func (s *ScheduleIDandPatientID) Get() *map[int]int {
	query := `
		SELECT id, patientid FROM schedule
	`
	rows, err := s.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandIntField(rows)
	common.CheckErr(err)

	return m
}

type ScheduleIDandDateAppointment struct {
	conn *pgx.Conn
}

func NewScheduleIDandDateAppointment(conn *pgx.Conn) *ScheduleIDandDateAppointment {
	value := new(ScheduleIDandDateAppointment)
	value.conn = conn
	return value
}

func (s *ScheduleIDandDateAppointment) Get() *map[int]time.Time {
	query := `
		SELECT id, dateappointment FROM schedule
	`
	rows, err := s.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadIDandDateField(rows)
	common.CheckErr(err)

	return m
}

type ScheduleDateAppointmentsByDoctorID struct {
	conn *pgx.Conn
}

func NewScheduleDateAppointmentsByDoctorID(conn *pgx.Conn) *ScheduleDateAppointmentsByDoctorID {
	value := new(ScheduleDateAppointmentsByDoctorID)
	value.conn = conn
	return value
}

func (s *ScheduleDateAppointmentsByDoctorID) Get(doctorid int) *[]time.Time {
	query := `
		SELECT dateappointment FROM schedule WHERE doctorid = @doctorid and dateappointment > CURRENT_DATE
	`
	rows, err := s.conn.Query(context.Background(), query)
	common.CheckErr(err)
	defer rows.Close()

	m, err := LoadDate(rows)
	common.CheckErr(err)

	return m
}

type AppointmentCreate struct {
	conn *pgx.Conn
}

func NewAppointmentCreate(conn *pgx.Conn) *AppointmentCreate {
	value := new(AppointmentCreate)
	value.conn = conn
	return value
}

func (a *AppointmentCreate) Insert(appointment *CreateAppointment) {
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

	_, err := a.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}

type AppointmentByID struct {
	conn *pgx.Conn
}

func NewAppointmentByID(conn *pgx.Conn) *AppointmentByID {
	value := new(AppointmentByID)
	value.conn = conn
	return value
}

func (a *AppointmentByID) Delete(appointmentID int) {
	query := `
        DELETE FROM schedule WHERE id = @id
    `
	args := pgx.NamedArgs{
		"id": appointmentID,
	}

	_, err := a.conn.Exec(context.Background(), query, args)
	common.CheckErr(err)
}
