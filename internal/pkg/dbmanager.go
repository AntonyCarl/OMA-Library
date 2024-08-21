package database

import (
	"database/sql"
	"fmt"
)

//const dbConn = "host=127.0.0.1 port=5432 user=postgres password=0000 dbname=omalibdb sslmode=disable"

type ConnectionParametrs struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func dbConnection(params ConnectionParametrs) (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s password=%s",
			params.Host, params.Port, params.User, params.DBName, params.SSLMode, params.Password))
	if err != nil {
		return nil, err
	}
	return db, err
}
