package psql

import (
	"database/sql"
	"fmt"
)

//const dbConn = "host=127.0.0.1 port=5432 user=postgres password=0000 dbname=omalibdb sslmode=disable"

// type connectionParametrs struct {
// 	Host     string
// 	Port     int
// 	User     string
// 	Password string
// 	DBName   string
// 	SSLMode  string
// }

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "omalibdb"
	sslmode  = "disable"
)

func DbConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode))
	if err != nil {
		return nil, err
	}
	return db, err
}
