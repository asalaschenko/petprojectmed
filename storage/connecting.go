package storage

import (
	"context"
	"fmt"
	"petprojectmed/common"

	"github.com/jackc/pgx/v5"
)

func GetConnectionDB() *pgx.Conn {
	username := "postgres"
	pass := "49236"
	host := "localhost"
	port := "5432"
	dbName := "clinicdb"

	conn, err := NewPostgres(username, pass, host, port, dbName)
	common.CheckErr(err)
	return conn
}

func NewPostgres(username, password, host, port, dbName string) (*pgx.Conn, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, dbName)
	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
