package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
)

func CheckErr(err error) {
	if err != nil {
		pc, fileName, line, _ := runtime.Caller(1)
		log.Println(GenLocationError(runtime.FuncForPC(pc).Name(), fileName, strconv.Itoa(line)))
		log.Fatalln(err)
	}
}

func GenLocationError(extFunction string, fileName string, lineOfCallCheckErr string) error {
	outputString := extFunction + ", " + filepath.Base(fileName) + ", " + "line of triggering error handler: " + lineOfCallCheckErr
	outputErr := errors.New(outputString)
	return outputErr
}

func CheckParseTimeValue(timeValue string) (bool, time.Time) {
	parseTime, err := time.Parse("15", timeValue)
	if err == nil {
		return true, parseTime
	}
	parseTime, err = time.Parse("15:04", timeValue)
	if err == nil {
		return true, parseTime
	}
	parseTime, err = time.Parse("15:04:05", timeValue)
	if err == nil {
		return true, parseTime
	}
	return false, time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
}

func CheckParseDateValue(dateValue string) (bool, time.Time) {
	parseDate, err := time.Parse("2006-01-02", dateValue)
	if err == nil {
		return true, parseDate
	}
	parseDate, err = time.Parse("02-01-2006", dateValue)
	if err == nil {
		return true, parseDate
	}
	return false, time.Date(0, 0, 0, 0, 0, 0, 0, time.UTC)
}

func CheckExternalUsedOfPK(conn *pgx.Conn, tableName string, primaryKey string, keyValue any) bool {
	query := `SELECT table_name, column_name from information_schema.key_column_usage kcu 
			WHERE kcu.constraint_name in (
			SELECT rc.constraint_name from information_schema.referential_constraints rc 
			WHERE rc.unique_constraint_name  IN 
			(SELECT kcu.constraint_name
			FROM information_schema.key_column_usage kcu
			WHERE table_name = @tableName and column_name = @primaryKey))`

	args := pgx.NamedArgs{
		"tableName":  tableName,
		"primaryKey": primaryKey,
	}

	rows, err := conn.Query(context.Background(), query, args)
	CheckErr(err)
	defer rows.Close()

	var mapTable = map[string]string{}
	for rows.Next() {
		var table, column string
		err := rows.Scan(&table, &column)
		log.Println(table, column)
		CheckErr(err)
		mapTable[table] = column
	}

	log.Println(mapTable)

	if len(mapTable) == 0 {
		return false
	}

	log.Println(len(mapTable))

	for k, v := range mapTable {
		query := `SELECT * FROM %s WHERE %s = @inputValue`
		query = fmt.Sprintf(query, k, v)

		args := pgx.NamedArgs{
			"inputValue": keyValue,
		}

		rows, err := conn.Query(context.Background(), query, args)
		CheckErr(err)
		defer rows.Close()

		if rows.Next() {
			return true
		}
	}
	return false
}
