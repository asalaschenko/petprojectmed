package storage

import (
	"context"
	"fmt"
	"petprojectmed/common"

	"github.com/jackc/pgx/v5"
)

func GetConnectionDB() *pgx.Conn {
	conn, err := NewPostgres(USERNAME, PASS, HOST, PORT, DB_NAME)
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
