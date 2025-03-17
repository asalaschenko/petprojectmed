package common

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

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

	if len(mapTable) == 0 {
		return false
	}

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
