package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	HOST     = "db"
	DATABASE = "todoapi"
	USER     = "postgres"
	PASSWORD = "postgres"
)

func ConnectToDatabase() (*sql.DB, error) {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
