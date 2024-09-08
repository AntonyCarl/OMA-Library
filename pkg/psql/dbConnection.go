package psql

import (
	"database/sql"
	"fmt"
	"log"
)

var DbConn *sql.DB

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "omalibdb"
	sslmode  = "disable"
)

func DbConnection() {
	var err error
	DbConn, err = sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			host, port, user, password, dbname, sslmode))
	if err != nil {
		log.Fatal(err)
	}
}
