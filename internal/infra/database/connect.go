package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	HOST      = "mysql"
	DATABASE  = "todoapi"
	USER      = "root"
	PASSWORD  = "pass"
	PORT      = "3306"
	DB_DRIVER = "mysql"
)

func ConnectToDatabase() (*sql.DB, error) {
	var connectionString string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", USER, PASSWORD, HOST, PORT, DATABASE)
	db, err := sql.Open(DB_DRIVER, connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
