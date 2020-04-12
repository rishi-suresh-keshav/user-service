package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(host, port, username, password, dbname string) (*DB, error) {
	dbSource := fmt.Sprintf(
		"root:%s@tcp(%s:%s)/%s?charset=utf8",
		password,
		host,
		port,
		dbname,
	)
	db, err := sql.Open("mysql", dbSource)

	if err != nil {
		panic(err)
	}

	dbConn.SQL = db
	return dbConn, nil
}