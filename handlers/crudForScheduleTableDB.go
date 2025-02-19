package handlers

import (
	"context"
	"petprojectmed/dto"
	"petprojectmed/utils"

	"github.com/jackc/pgx/v5"
)

func GetAllAppointments(conn *pgx.Conn) *[]dto.DoctorTable {
	query := `SELECT * FROM schedule`

	rows, err := conn.Query(context.Background(), query)
	utils.CheckErr(err)
	defer rows.Close()

	doctors, err := ScanLoadDoctorEntries(rows)
	utils.CheckErr(err)

	return doctors
}

func GetAppointsmentByID
